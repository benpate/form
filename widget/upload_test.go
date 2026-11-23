package widget

import (
	"strings"
	"testing"

	"github.com/benpate/form"
	"github.com/benpate/rosetta/mapof"
	"github.com/benpate/rosetta/schema"
	"github.com/stretchr/testify/require"
)

// uploadSchema returns a schema with the string paths the upload tests draw
func uploadSchema() schema.Schema {
	return schema.Schema{
		Element: schema.Object{
			Properties: map[string]schema.Element{
				"iconUrl":  schema.String{},
				"audioUrl": schema.String{},
				"filename": schema.String{},
				"id":       schema.String{},
			},
		},
	}
}

// uploadEditor draws one upload element for an object, and fails the test if the render fails
func uploadEditor(t *testing.T, element form.Element, value mapof.Any) string {
	t.Helper()
	UseAll()

	result, err := form.Editor(uploadSchema(), element, value, nil)
	require.NoError(t, err)
	return result
}

// TestUpload_Empty requires that a field with no file draws the zone, a visible plus, the file input, and the status line
func TestUpload_Empty(t *testing.T) {

	element := form.Element{Type: "upload", Path: "iconUrl", Label: "Icon", Options: mapof.Template{"accept": "image/*"}}
	result := uploadEditor(t, element, mapof.Any{})

	require.Contains(t, result, `<div class="upload-single" data-upload-single="true">`)
	require.Contains(t, result, `<button type="button" class="upload-multiple-add" title="Choose a file" aria-label="Choose a file">`)
	require.Contains(t, result, `<input type="file" name="iconUrl" id="iconUrl.upload" accept="image/*" aria-label="Icon">`)
	require.Contains(t, result, `<div class="upload-multiple-status" role="status"></div>`)

	// No saved file means no tile, no hidden value, and no X
	require.NotContains(t, result, "upload-multiple-tile")
	require.NotContains(t, result, `type="hidden"`)
	require.NotContains(t, result, "upload-multiple-remove")
	require.NotContains(t, result, "hidden=")
}

// TestUpload_SavedImage requires that a saved image draws as a preview tile that keeps its URL, with the plus hidden
func TestUpload_SavedImage(t *testing.T) {

	element := form.Element{Type: "upload", Path: "iconUrl", Options: mapof.Template{"accept": "image/*"}}
	result := uploadEditor(t, element, mapof.Any{"iconUrl": "/123/attachments/abc.webp?width=300"})

	require.Contains(t, result, `<div class="upload-multiple-tile hover-trigger" data-upload-single-saved="true" tabindex="0" aria-label="Replace abc.webp">`)
	require.Contains(t, result, `<img src="/123/attachments/abc.webp?width=300" class="upload-multiple-preview" alt="abc.webp" draggable="false">`)
	require.Contains(t, result, `<input type="hidden" name="iconUrl" value="/123/attachments/abc.webp?width=300">`)
	require.Contains(t, result, `class="upload-multiple-add" title="Choose a file" hidden="hidden"`)

	// Without a "delete" option there is no X
	require.NotContains(t, result, "upload-multiple-remove")
	require.NotContains(t, result, "hx-post")
}

// TestUpload_SavedOtherFile requires that anything but an image draws as an icon and a filename
func TestUpload_SavedOtherFile(t *testing.T) {

	for _, accept := range []string{"audio/*", "", "application/pdf"} {
		element := form.Element{Type: "upload", Path: "audioUrl", Options: mapof.Template{"accept": accept}}
		result := uploadEditor(t, element, mapof.Any{"audioUrl": "/123/attachments/episode.mp3"})

		require.Contains(t, result, `<i class="bi bi-file-earmark upload-multiple-icon"></i>`, "accept %q", accept)
		require.Contains(t, result, `<span class="upload-multiple-filename">episode.mp3</span>`, "accept %q", accept)
		require.NotContains(t, result, "<img", "accept %q", accept)
		require.NotContains(t, result, "<audio", "accept %q", accept)
	}
}

// TestUpload_FilenameOption requires that the "filename" option's path names the file, and that an empty or unknown path falls back to the URL
func TestUpload_FilenameOption(t *testing.T) {

	element := form.Element{Type: "upload", Path: "audioUrl", Options: mapof.Template{"accept": "audio/*", "filename": "filename"}}

	result := uploadEditor(t, element, mapof.Any{"audioUrl": "/123/attachments/abc", "filename": "My Episode.mp3"})
	require.Contains(t, result, `<span class="upload-multiple-filename">My Episode.mp3</span>`)
	require.Contains(t, result, `aria-label="Replace My Episode.mp3"`)

	result = uploadEditor(t, element, mapof.Any{"audioUrl": "/123/attachments/abc"})
	require.Contains(t, result, `<span class="upload-multiple-filename">abc</span>`)

	element.Options["filename"] = "no.such.path"
	result = uploadEditor(t, element, mapof.Any{"audioUrl": "/123/attachments/abc"})
	require.Contains(t, result, `<span class="upload-multiple-filename">abc</span>`)
}

// TestUpload_Delete requires that the "delete" option draws the X, with its URL rendered against the object
func TestUpload_Delete(t *testing.T) {

	parsed, err := form.Parse(map[string]any{"type": "upload", "path": "iconUrl", "options": map[string]any{"accept": "image/*", "delete": "/{{.id}}/delete-icon"}})
	require.NoError(t, err)

	result := uploadEditor(t, parsed, mapof.Any{"iconUrl": "/123/attachments/abc.webp", "id": "123"})
	require.Contains(t, result, `<button type="button" class="upload-multiple-remove hover-show" hx-post="/123/delete-icon" hx-confirm="Delete this file?" hx-swap="none" hx-push-url="false" aria-label="Delete abc.webp">`)

	// With no saved file there is nothing to delete
	result = uploadEditor(t, parsed, mapof.Any{"id": "123"})
	require.NotContains(t, result, "upload-multiple-remove")
}

// TestUpload_BrokenDeleteTemplate requires that a delete URL that cannot render fails the whole render
func TestUpload_BrokenDeleteTemplate(t *testing.T) {
	UseAll()

	parsed, err := form.Parse(map[string]any{"type": "upload", "path": "iconUrl", "options": map[string]any{"delete": "/{{.id.missing.deeper}}/delete"}})
	require.NoError(t, err)

	_, err = form.Editor(uploadSchema(), parsed, mapof.Any{"iconUrl": "/a.webp", "id": "123"}, nil)
	require.Error(t, err)
}

// TestUpload_Escaping requires that a hostile URL and filename never reach the page as markup
func TestUpload_Escaping(t *testing.T) {

	hostile := `"><script>alert(1)</script>`

	for _, accept := range []string{"image/*", "audio/*"} {
		element := form.Element{Type: "upload", Path: "iconUrl", Options: mapof.Template{"accept": accept, "filename": "filename", "delete": "/delete"}}
		result := uploadEditor(t, element, mapof.Any{"iconUrl": "/files/" + hostile, "filename": hostile})

		require.NotContains(t, result, "<script>", "accept %q", accept)
		require.NotContains(t, result, `"">`, "accept %q", accept)
	}
}

// TestUpload_Filename covers every shape of URL the fallback name is cut from
func TestUpload_Filename(t *testing.T) {

	schema := uploadSchema()
	element := form.Element{Type: "upload", Path: "iconUrl"}

	check := func(url string, expected string) {
		t.Helper()
		require.Equal(t, expected, Upload{}.filename(&element, &schema, mapof.Any{}, url), "url %q", url)
	}

	check("/123/attachments/abc.webp", "abc.webp")
	check("/123/attachments/abc.webp?width=300&height=2", "abc.webp")
	check("/123/attachments/abc.webp#top", "abc.webp")
	check("https://example.com/a/b/c.png?x=/y/z", "c.png")
	check("plain.png", "plain.png")
	check("/trailing/slash/", "slash")
	check("/", "file")
	check("?only=query", "file")
	check("", "file")
	check(strings.Repeat("a/", 10000)+"end", "end")
}

// TestUpload_IDAndOptions requires that the element's own ID, "multiple", and "class" reach the markup
func TestUpload_IDAndOptions(t *testing.T) {

	element := form.Element{Type: "upload", ID: "my-upload", Path: "iconUrl", Description: "Square works best", Options: mapof.Template{"multiple": "true", "class": "wide"}}
	result := uploadEditor(t, element, mapof.Any{})

	require.Contains(t, result, `<div class="upload-single wide" data-upload-single="true">`)
	require.Contains(t, result, `<input type="file" name="iconUrl" id="my-upload" multiple="multiple" aria-description="Square works best">`)
}

// TestUpload_View requires that the read-only view prints the URL as text, or N/A without one
func TestUpload_View(t *testing.T) {
	UseAll()

	element := form.Element{Type: "upload", Path: "iconUrl"}

	result, err := form.Viewer(uploadSchema(), element, mapof.Any{"iconUrl": "/a/<b>.webp"}, nil)
	require.NoError(t, err)
	require.Contains(t, result, `/a/&lt;b&gt;.webp`)

	result, err = form.Viewer(uploadSchema(), element, mapof.Any{}, nil)
	require.NoError(t, err)
	require.Contains(t, result, "N/A")
}

// TestUpload_Metadata pins the widget's labels, description position, and multipart encoding
func TestUpload_Metadata(t *testing.T) {
	require.True(t, Upload{}.ShowLabels())
	require.Equal(t, "BOTTOM", Upload{}.ShowDescriptions())
	require.Equal(t, "multipart/form-data", Upload{}.Encoding(nil))
}

// FuzzUpload_Filename requires that no URL panics the fallback name, and that the name is never empty or holds a slash
func FuzzUpload_Filename(f *testing.F) {

	for _, seed := range []string{"", "/", "?", "#", "/a/b.png?x=1#y", "https://example.com/", "a//b//", "\xff\xfe/\x00"} {
		f.Add(seed)
	}

	schema := uploadSchema()
	element := form.Element{Type: "upload", Path: "iconUrl"}

	f.Fuzz(func(t *testing.T, url string) {
		result := Upload{}.filename(&element, &schema, mapof.Any{}, url)
		require.NotEmpty(t, result)
		require.NotContains(t, result, "/")
	})
}
