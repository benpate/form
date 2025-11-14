package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
)

type WYSIWYG struct{}

func (widget WYSIWYG) View(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {
	// find the path and schema to use
	valueString := e.GetString(value, &f.Schema)
	b.WriteString(valueString) // TODO: LOW: apply schema formats?
	return nil
}

func (widget WYSIWYG) Edit(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	valueString := e.GetString(value, &f.Schema)

	// Start building a new tag
	b.Input("hidden", e.Path).
		Value(valueString).
		Close()

	b.Div().ID("content-editor")
	b.Div().Class("wysiwyg").Script("install wysiwyg(name:'" + e.Path + "') install hotkey")
	b.Div().Class("wysiwyg-toolbar").Attr("hidden", "true")
	{
		b.Span().Class("wysiwyg-toolbar-group").EndBracket()
		b.Button().Data("command", "formatBlock").Data("command-value", "h1").InnerText("H1").Close()
		b.Button().Data("command", "formatBlock").Data("command-value", "h2").InnerText("H2").Close()
		b.Button().Data("command", "formatBlock").Data("command-value", "h3").InnerText("H3").Close()
		b.Button().Data("command", "formatBlock").Data("command-value", "p").InnerText("P").Close()
		b.Close()
	}
	{
		b.Span().Class("wysiwyg-toolbar-group").EndBracket()
		b.Button().Data("command", "bold").Aria("keyshortcuts", "Ctrl+B").InnerText("B").Close()
		b.Button().Data("command", "italic").Aria("keyshortcuts", "Ctrl+I").InnerText("I").Close()
		b.Button().Data("command", "underline").Aria("keyshortcuts", "Ctrl+U").InnerText("U").Close()
		b.Close()
	}
	{
		b.Span().Class("wysiwyg-toolbar-group").EndBracket()
		b.Button().Data("command", "createLink").Aria("keyshortcuts", "Ctrl+K")
		b.Container("i").Class("ti", "ti-link").Close()
		b.Close()
		b.Button().Data("command", "unlink").Aria("keyshortcuts", "Ctrl+Shift+K")
		b.Container("i").Class("ti", "ti-unlink").Close()
		b.Close()
		b.Close()
	}
	{
		b.Span().Class("wysiwyg-toolbar-group").Attr("hidden", "true").EndBracket()
		b.Button().Data("command", "cut").Aria("keyshortcuts", "Ctrl+X").InnerText("Cut").Close()
		b.Button().Data("command", "copy").Aria("keyshortcuts", "Ctrl+C").InnerText("Copy").Close()
		b.Button().Data("command", "paste").Aria("keyshortcuts", "Ctrl+V").InnerText("Paste").Close()
		b.Button().Data("command", "undo").Aria("keyshortcuts", "Ctrl+Z").InnerText("Undo").Close()
		b.Button().Data("command", "redo").Aria("keyshortcuts", "Ctrl+Shift+Z").InnerText("Redo").Close()
		b.Close()
	}
	b.Close()

	b.Div().Class("wysiwyg-editor").InnerText(valueString)
	b.CloseAll()

	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget WYSIWYG) ShowLabels() bool {
	return true
}

func (widget WYSIWYG) Encoding(_ *form.Element) string {
	return ""
}
