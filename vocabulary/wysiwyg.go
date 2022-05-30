package vocabulary

import (
	"github.com/benpate/convert"
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/schema"
)

func WYSIWYG(library *form.Library) {

	library.Register("wysiwyg", func(f form.Form, s *schema.Schema, v interface{}, b *html.Builder) error {

		// find the path and schema to use
		value, _, _ := s.Get(v, f.Path)

		valueString := convert.String(value)

		// Start building a new tag
		b.Input("hidden", f.Path).
			Value(valueString).
			Close()

		b.Div().Class("wysiwyg").Script("install wysiwyg(name:'" + f.Path + "') install hotkey")
		b.Div().Class("wysiwyg-toolbar").Attr("hidden", "true")
		{
			b.Span().Class("wysiwyg-toolbar-group").EndBracket()
			b.Button().Data("command", "formatBlock").Data("command-value", "h1").InnerHTML("H1").Close()
			b.Button().Data("command", "formatBlock").Data("command-value", "h2").InnerHTML("H2").Close()
			b.Button().Data("command", "formatBlock").Data("command-value", "h3").InnerHTML("H3").Close()
			b.Button().Data("command", "formatBlock").Data("command-value", "p").InnerHTML("P").Close()
			b.Close()
		}
		{
			b.Span().Class("wysiwyg-toolbar-group").EndBracket()
			b.Button().Data("command", "bold").Aria("keyshortcuts", "Ctrl+B").InnerHTML("B").Close()
			b.Button().Data("command", "italic").Aria("keyshortcuts", "Ctrl+I").InnerHTML("I").Close()
			b.Button().Data("command", "underline").Aria("keyshortcuts", "Ctrl+U").InnerHTML("U").Close()
			b.Close()
		}
		{
			b.Span().Class("wysiwyg-toolbar-group").EndBracket()
			b.Button().Data("command", "createLink").Aria("keyshortcuts", "Ctrl+K")
			b.Container("i").Class("fa-solid", "fa-link").Close()
			b.Close()
			b.Button().Data("command", "unlink").Aria("keyshortcuts", "Ctrl+Shift+K")
			b.Container("i").Class("fa-solid", "fa-unlink").Close()
			b.Close()
			b.Close()
		}
		{
			b.Span().Class("wysiwyg-toolbar-group").Attr("hidden", "true").EndBracket()
			b.Button().Data("command", "cut").Aria("keyshortcuts", "Ctrl+X").InnerHTML("Cut").Close()
			b.Button().Data("command", "copy").Aria("keyshortcuts", "Ctrl+C").InnerHTML("Copy").Close()
			b.Button().Data("command", "paste").Aria("keyshortcuts", "Ctrl+V").InnerHTML("Paste").Close()
			b.Button().Data("command", "undo").Aria("keyshortcuts", "Ctrl+Z").InnerHTML("Undo").Close()
			b.Button().Data("command", "redo").Aria("keyshortcuts", "Ctrl+Shift+Z").InnerHTML("Redo").Close()
			b.Close()
		}
		b.Close()

		b.Div().Class("wysiwyg-editor").InnerHTML(valueString)
		b.CloseAll()

		return nil
	})
}
