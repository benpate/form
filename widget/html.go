package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
)

type HTML struct{}

func (widget HTML) View(_ *form.Form, e *form.Element, _ form.LookupProvider, _ any, b *html.Builder) error {
	b.Div().InnerHTML(e.Description).Close()
	return nil
}

func (widget HTML) Edit(_ *form.Form, e *form.Element, _ form.LookupProvider, _ any, b *html.Builder) error {
	b.Div().InnerHTML(e.Description).Close()
	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

// ShowLabels is a part of the Widget interface.
// It returns TRUE if this widget requires labels to be displayed around it.
// For HTML widgets, labels are not shown, so this always returns FALSE.
func (widget HTML) ShowLabels() bool {
	return false
}

// Encoding is a part of the Widget interface.
// It returns the encoding type for this widget.
// For HTML widgets, there is no special encoding,
// so this always returns an empty string.
func (widget HTML) Encoding(_ *form.Element) string {
	return ""
}
