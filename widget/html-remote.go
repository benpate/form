package widget

import (
	"bytes"
	"text/template"

	"github.com/benpate/derp"
	"github.com/benpate/form"
	"github.com/benpate/html"
)

type HTMLRemote struct{}

func (widget HTMLRemote) View(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {
	return widget.Edit(f, e, nil, value, b)
}

func (widget HTMLRemote) Edit(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {

	const location = "widget.HTMLRemote.Edit"

	// Collect and parse the Remote URL template
	remoteURL := e.Options.GetString("url")
	remoteTemplate, err := template.New("").Parse(remoteURL)

	if err != nil {
		return derp.Wrap(err, location, "Unable to parse remote URL template", remoteURL)
	}

	// Replace values in the template
	buffer := bytes.Buffer{}
	if err := remoteTemplate.Execute(&buffer, value); err != nil {
		return derp.Wrap(err, location, "Unable to ececute remote URL template", remoteURL)
	}

	b.Div().
		Attr("hx-get", buffer.String()).
		Attr("hx-swap", "innerHTML").
		Attr("hx-trigger", "intersect once").
		Attr("hx-target", "this").
		Attr("hx-push-url", "false").
		Close()

	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget HTMLRemote) ShowLabels() bool {
	return false
}

func (widget HTMLRemote) Encoding(_ *form.Element) string {
	return ""
}
