package widget

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("wysiwyg", WYSIWYG{})
}

type WYSIWYG struct{}

func (widget WYSIWYG) View(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {
	// find the path and schema to use
	valueString := element.GetString(value, s)
	b.WriteString(valueString) // TODO: LOW: apply schema formats?
	return nil
}

func (widget WYSIWYG) Edit(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return WYSIWYG{}.View(element, s, lookupProvider, value, b)
	}

	// find the path and schema to use
	valueString := element.GetString(value, s)

	// Start building a new tag
	b.Input("hidden", element.Path).
		Value(valueString).
		Close()

	b.Div().ID("content-editor")
	b.Div().Class("wysiwyg").Script("install wysiwyg(name:'" + element.Path + "') install hotkey")
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

func (widget WYSIWYG) Encoding(_ *Element) string {
	return ""
}
