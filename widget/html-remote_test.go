package widget

import (
	"testing"

	"github.com/benpate/form"
	"github.com/benpate/rosetta/mapof"
	"github.com/benpate/rosetta/schema"
	"github.com/stretchr/testify/require"
)

// remoteRecord stands in for the object a form edits, with the fields a URL template names.
type remoteRecord struct {
	ID       string
	ParentID string
}

// TestHTMLRemote_LoadsTheURLAtOnce draws one empty div that fills itself from the templated URL
// as soon as it is on the page, with the element's ID so the layout's label names it
func TestHTMLRemote_LoadsTheURLAtOnce(t *testing.T) {

	f := htmlRemoteForm(form.Element{Type: "html-remote", Label: "Photos", Description: "Up to 16.", Options: mapof.Template{"url": "/{{.ID}}/photos"}})
	html, err := f.Editor(remoteRecord{ID: "abc"}, nil)
	require.NoError(t, err)

	// The layout gives an element without an ID one of its own, from its path and type
	require.Contains(t, html, `<label for="_html-remote" aria-hidden="true">Photos</label><div id="_html-remote" hx-get="/abc/photos" hx-swap="innerHTML" hx-trigger="load" hx-target="this" hx-push-url="false"></div><div aria-hidden="true" class="text-sm gray40">Up to 16.</div>`)
	require.NotContains(t, html, "intersect")
}

// TestHTMLRemote_URLIsATemplateOfTheObject renders the URL against the object, so a create
// form can name its parent's route and an edit form the object's own
func TestHTMLRemote_URLIsATemplateOfTheObject(t *testing.T) {

	f := htmlRemoteForm(form.Element{Type: "html-remote", ID: "zone", Options: mapof.Template{"url": "/{{.ParentID}}/add?action=photos"}})
	html, err := f.Editor(remoteRecord{ID: "abc", ParentID: "parent"}, nil)
	require.NoError(t, err)

	require.Contains(t, html, `<div id="zone" hx-get="/parent/add?action=photos" hx-swap="innerHTML"`)
	require.NotContains(t, html, "_html-remote")
}

// TestHTMLRemote_BrokenURLFailsTheRender fails the render rather than drawing a dead element
func TestHTMLRemote_BrokenURLFailsTheRender(t *testing.T) {

	f := htmlRemoteForm(form.Element{Type: "html-remote", Options: mapof.Template{"url": "/{{.Missing.Field}}/photos"}})
	_, err := f.Editor(remoteRecord{ID: "abc"}, nil)
	require.Error(t, err)
}

// TestHTMLRemote_URLIsEscaped writes the URL as an attribute, so a hostile value cannot leave it
func TestHTMLRemote_URLIsEscaped(t *testing.T) {

	f := htmlRemoteForm(form.Element{Type: "html-remote", Options: mapof.Template{"url": "/{{.ID}}/photos"}})
	html, err := f.Editor(remoteRecord{ID: `abc"><script>alert(1)</script>`}, nil)
	require.NoError(t, err)

	require.NotContains(t, html, "<script>")
	require.Contains(t, html, `hx-get="/abc&#34;&gt;&lt;script&gt;alert(1)&lt;/script&gt;/photos"`)
}

// TestHTMLRemote_ViewIsTheSameElement draws the same element in a read-only form, so a view
// that wants read-only content asks for it in the URL
func TestHTMLRemote_ViewIsTheSameElement(t *testing.T) {

	f := htmlRemoteForm(form.Element{Type: "html-remote", Options: mapof.Template{"url": "/{{.ID}}/photos?readonly=true"}})
	html, err := f.Viewer(remoteRecord{ID: "abc"}, nil)
	require.NoError(t, err)

	require.Contains(t, html, `<div id="_html-remote" hx-get="/abc/photos?readonly=true" hx-swap="innerHTML" hx-trigger="load" hx-target="this" hx-push-url="false"></div>`)
}

// TestHTMLRemote_Metadata pins the widget's part in the form around it
func TestHTMLRemote_Metadata(t *testing.T) {

	widget := HTMLRemote{}

	require.True(t, widget.ShowLabels())
	require.Equal(t, "BOTTOM", widget.ShowDescriptions())
	require.Equal(t, "", widget.Encoding(&form.Element{}))
}

// htmlRemoteForm returns a Form that draws an html-remote element, which has no path, beside a
// text field, so that a test can see the layout give it a label, an ID, and a description.
func htmlRemoteForm(element form.Element) form.Form {

	UseAll()

	return form.New(
		schema.New(schema.Object{
			Properties: schema.ElementMap{
				"label": schema.String{},
			},
		}),
		form.Element{
			Type: "layout-vertical",
			Children: []form.Element{
				{Type: "text", Path: "label"},
				element,
			},
		},
	)
}
