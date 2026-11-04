package widget

import (
	"net/url"
	"strings"
	"testing"

	"github.com/benpate/form"
	"github.com/benpate/rosetta/mapof"
	"github.com/benpate/rosetta/schema"
	"github.com/stretchr/testify/require"
)

// savedRecord stands in for a record the host has stored, as benpate/data's journal reports it.
type savedRecord struct {
	ID string
}

// IsNew reports that the record has been saved.
func (savedRecord) IsNew() bool {
	return false
}

// newRecord stands in for a record the host has not stored yet.
type newRecord struct {
	ID string
}

// IsNew reports that the record has never been saved.
func (newRecord) IsNew() bool {
	return true
}

// uploadMultipleOptions is the full set of options a template would write.
func uploadMultipleOptions() mapof.Template {
	return mapof.Template{
		"name":    "photos",
		"accept":  "image/*",
		"url":     "/{{.ID}}/photos",
		"staging": "/@me/upload-staging",
	}
}

// uploadMultipleForm returns a Form that draws an upload-multiple widget, which has no path.
func uploadMultipleForm(options mapof.Template) form.Form {

	UseAll()

	return form.New(
		schema.New(schema.Object{
			Properties: schema.ElementMap{
				"label": schema.String{},
			},
		}),
		form.Element{
			Type:        "upload-multiple",
			Label:       "Photos",
			Description: "Up to 64.",
			Options:     options,
		},
	)
}

// editUploadMultiple renders the editor for the given object.
func editUploadMultiple(t *testing.T, options mapof.Template, value any) string {
	t.Helper()

	f := uploadMultipleForm(options)
	result, err := f.Editor(value, nil)
	require.NoError(t, err)

	return result
}

func TestUploadMultiple_ZoneCarriesNameAcceptAndURL(t *testing.T) {

	result := editUploadMultiple(t, uploadMultipleOptions(), mapof.Any{"ID": "abc"})

	require.Contains(t, result, `<div id="upload-multiple-photos" class="upload-multiple" data-upload-multiple="true" data-name="photos" data-accept="image/*" data-url="/abc/photos">`)
}

func TestUploadMultiple_TileRowLoadsFromTheURL(t *testing.T) {

	result := editUploadMultiple(t, uploadMultipleOptions(), mapof.Any{"ID": "abc"})

	require.Contains(t, result, `<div id="upload-multiple-photos-tiles" class="upload-multiple-tiles" hx-get="/abc/photos" hx-trigger="load" hx-target="this" hx-swap="outerHTML" hx-push-url="false"></div>`)
}

func TestUploadMultiple_NoURLDrawsAPlainRow(t *testing.T) {

	options := uploadMultipleOptions()
	delete(options, "url")
	result := editUploadMultiple(t, options, mapof.Any{"ID": "abc"})

	require.Contains(t, result, `<div id="upload-multiple-photos" class="upload-multiple" data-upload-multiple="true" data-name="photos" data-accept="image/*">`)
	require.Contains(t, result, `<div id="upload-multiple-photos-tiles" class="upload-multiple-tiles"></div>`)
	require.NotContains(t, result, "hx-get")
	require.NotContains(t, result, "data-url")
}

func TestUploadMultiple_NewObjectStagesInsteadOfLoading(t *testing.T) {

	// A record that is not saved yet has nothing to load, and its route does not exist, so its
	// zone carries the staging route and its row appends each staged tile the host answers with
	result := editUploadMultiple(t, uploadMultipleOptions(), newRecord{ID: "abc"})

	require.Contains(t, result, `<div id="upload-multiple-photos" class="upload-multiple" data-upload-multiple="true" data-name="photos" data-accept="image/*" data-staging-url="/@me/upload-staging">`)
	require.Contains(t, result, `<div id="upload-multiple-photos-tiles" class="upload-multiple-tiles" hx-post="/@me/upload-staging" hx-trigger="upload-multiple-added" hx-params="photos" hx-target="this" hx-swap="beforeend" hx-push-url="false" hx-sync="closest .upload-multiple:queue all"></div>`)
	require.NotContains(t, result, "hx-get")
	require.NotContains(t, result, "data-url")

	// A saved record loads its files from its own route, and stages nothing
	result = editUploadMultiple(t, uploadMultipleOptions(), savedRecord{ID: "abc"})

	require.Contains(t, result, `data-url="/abc/photos"`)
	require.Contains(t, result, `hx-get="/abc/photos"`)
	require.NotContains(t, result, "staging")
}

func TestUploadMultiple_StagingIsRequiredForANewObject(t *testing.T) {

	options := uploadMultipleOptions()
	delete(options, "staging")

	// A new record with nowhere to stage a file fails the render, because the zone would be a dead end
	f := uploadMultipleForm(options)
	_, err := f.Editor(newRecord{ID: "abc"}, nil)
	require.Error(t, err)

	// A saved record never stages, so it renders without the option
	result, err := f.Editor(savedRecord{ID: "abc"}, nil)
	require.NoError(t, err)
	require.Contains(t, result, `data-url="/abc/photos"`)
}

func TestUploadMultiple_AcceptIsOptional(t *testing.T) {

	options := uploadMultipleOptions()
	delete(options, "accept")
	result := editUploadMultiple(t, options, mapof.Any{"ID": "abc"})

	require.NotContains(t, result, "data-accept")
	require.NotContains(t, result, `accept=`)
}

func TestUploadMultiple_NameIsRequired(t *testing.T) {

	options := uploadMultipleOptions()
	delete(options, "name")

	f := uploadMultipleForm(options)
	_, err := f.Editor(mapof.Any{"ID": "abc"}, nil)
	require.Error(t, err)
}

func TestUploadMultiple_FileInputIsHiddenSingleAndNeverRequired(t *testing.T) {

	result := editUploadMultiple(t, uploadMultipleOptions(), mapof.Any{"ID": "abc"})

	require.Contains(t, result, `<input type="file" name="photos" id="upload-multiple-photos-input" accept="image/*" hidden="hidden" aria-label="Photos" aria-description="Up to 64.">`)
	require.NotContains(t, result, "required")
	require.NotContains(t, result, ` multiple=`)
}

func TestUploadMultiple_PlusTileFollowsTheRowAndOpensThePicker(t *testing.T) {

	result := editUploadMultiple(t, uploadMultipleOptions(), mapof.Any{"ID": "abc"})

	require.Contains(t, result, `</div><button id="upload-multiple-photos-add" type="button" class="upload-multiple-add" title="Add a file" aria-label="Add a file"><i class="bi bi-plus-lg"></i></button><input type="file"`)
	require.Contains(t, result, `<div id="upload-multiple-photos-status" class="upload-multiple-status" role="status"></div>`)
}

func TestUploadMultiple_EveryButtonIsTypeButton(t *testing.T) {

	result := editUploadMultiple(t, uploadMultipleOptions(), mapof.Any{"ID": "abc"})

	// A bare <button> inside a <form> defaults to type="submit"
	require.Equal(t, 1, strings.Count(result, "<button"))
	require.Equal(t, 1, strings.Count(result, `type="button"`))
}

func TestUploadMultiple_ElementIDWinsOverTheName(t *testing.T) {

	f := uploadMultipleForm(uploadMultipleOptions())
	f.Element.ID = "photo.zone"
	result, err := f.Editor(mapof.Any{"ID": "abc"}, nil)
	require.NoError(t, err)

	require.Contains(t, result, `<div id="photo-zone" class="upload-multiple"`)
	require.Contains(t, result, `<div id="photo-zone-tiles" class="upload-multiple-tiles"`)
	require.Contains(t, result, `id="photo-zone-input"`)
	require.Contains(t, result, `id="photo-zone-add"`)
	require.Contains(t, result, `id="photo-zone-status"`)
}

func TestUploadMultiple_DottedNameMakesADashedID(t *testing.T) {

	options := uploadMultipleOptions()
	options["name"] = "data.photos"
	result := editUploadMultiple(t, options, mapof.Any{"ID": "abc"})

	require.Contains(t, result, `<div id="upload-multiple-data-photos" class="upload-multiple" data-upload-multiple="true" data-name="data.photos"`)
	require.Contains(t, result, `<input type="file" name="data.photos" id="upload-multiple-data-photos-input"`)
}

func TestUploadMultiple_URLAndNameAreEscaped(t *testing.T) {

	options := uploadMultipleOptions()
	options["name"] = `photos"onload="x`
	options["url"] = `/{{.ID}}/photos?a="b"&c=d`
	result := editUploadMultiple(t, options, mapof.Any{"ID": `a"b`})

	require.NotContains(t, result, `"onload="`)
	require.Contains(t, result, `data-name="photos&#34;onload=&#34;x"`)
	require.Contains(t, result, `data-url="/a&#34;b/photos?a=&#34;b&#34;&amp;c=d"`)
	require.Contains(t, result, `hx-get="/a&#34;b/photos?a=&#34;b&#34;&amp;c=d"`)
}

func TestUploadMultiple_BrokenURLTemplateFailsTheRender(t *testing.T) {

	options := uploadMultipleOptions()
	options["url"] = `/{{template "nosuch"}}`

	f := uploadMultipleForm(options)

	_, err := f.Editor(mapof.Any{"ID": "abc"}, nil)
	require.Error(t, err)

	_, err = f.Viewer(mapof.Any{"ID": "abc"}, nil)
	require.Error(t, err)

	// A broken staging template fails a new object's editor the same way
	options = uploadMultipleOptions()
	options["staging"] = `/{{template "nosuch"}}`

	f = uploadMultipleForm(options)
	_, err = f.Editor(newRecord{ID: "abc"}, nil)
	require.Error(t, err)
}

func TestUploadMultiple_StagingURLIsEscaped(t *testing.T) {

	options := uploadMultipleOptions()
	options["staging"] = `/@me/upload-staging?a="b"&c=d`
	result := editUploadMultiple(t, options, newRecord{ID: "abc"})

	require.Contains(t, result, `data-staging-url="/@me/upload-staging?a=&#34;b&#34;&amp;c=d"`)
	require.Contains(t, result, `hx-post="/@me/upload-staging?a=&#34;b&#34;&amp;c=d"`)
}

func TestUploadMultiple_ViewDrawsTheRowOnly(t *testing.T) {

	f := uploadMultipleForm(uploadMultipleOptions())
	result, err := f.Viewer(mapof.Any{"ID": "abc"}, nil)
	require.NoError(t, err)

	require.Equal(t, `<div class="layout-value"><div id="upload-multiple-photos-tiles" class="upload-multiple-tiles" hx-get="/abc/photos" hx-trigger="load" hx-target="this" hx-swap="outerHTML" hx-push-url="false"></div></div>`, result)

	// A new object has nothing to show
	result, err = f.Viewer(newRecord{ID: "abc"}, nil)
	require.NoError(t, err)
	require.Equal(t, `<div class="layout-value"><div id="upload-multiple-photos-tiles" class="upload-multiple-tiles"></div></div>`, result)
}

func TestUploadMultiple_EncodingIsMultipartAndPromotes(t *testing.T) {

	require.Equal(t, "multipart/form-data", UploadMultiple{}.Encoding(&form.Element{}))

	// A layout holding the widget promotes the encoding to the whole form
	f := uploadMultipleForm(uploadMultipleOptions())
	f.Element = form.Element{Type: "layout-vertical", Children: []form.Element{f.Element}}
	require.Equal(t, "multipart/form-data", f.Encoding())
}

func TestUploadMultiple_NothingIsBoundToTheObject(t *testing.T) {

	// With no path, the widget is not in the form's write path at all
	f := uploadMultipleForm(uploadMultipleOptions())
	object := mapof.Any{}
	values := url.Values{
		"photos":       {"ignored"},
		"photos.order": {"aaa", "bbb"},
	}

	require.NoError(t, f.SetURLValues(&object, values, nil))
	require.Empty(t, object)
	require.NoError(t, f.Validate())
}

func TestUploadMultiple_WireFormatConstant(t *testing.T) {

	// Emissary's steps import this, so a change here changes what they read
	require.Equal(t, ".order", form.UploadOrderSuffix)
}

func TestUploadMultiple_Metadata(t *testing.T) {

	require.True(t, UploadMultiple{}.ShowLabels())
	require.Equal(t, "BOTTOM", UploadMultiple{}.ShowDescriptions())
}

func TestIsNew(t *testing.T) {

	require.True(t, isNew(newRecord{}))
	require.False(t, isNew(savedRecord{}))
	require.False(t, isNew(mapof.Any{"ID": "abc"}))
	require.False(t, isNew(nil))
	require.False(t, isNew("abc"))
}
