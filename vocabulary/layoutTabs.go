package vocabulary

import (
	"strings"

	"github.com/benpate/convert"
	"github.com/benpate/derp"
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/schema"
)

func Tab(library *form.Library) {

	library.Register("layout-tabs", func(f form.Form, s *schema.Schema, v interface{}, b *html.Builder) error {

		if len(f.Label) > 0 {
			b.Div().Class("layout-title").InnerHTML(f.Label).Close()
		}

		// Make a placeholder for labels
		labels := make([]string, 0)

		// If we have a configuration option for labels,
		// parse it into a slice
		if labelString, ok := f.Options["labels"].(string); ok {
			labels = strings.Split(labelString, ",")
		}

		b.Div().Class("tabs").Script("install TabContainer")
		b.Div().Role("tablist")

		for index := range f.Children {

			var label string

			indexString := convert.String(index)

			// Use the best label we have (configured, or generated)
			if index < len(labels) {
				label = labels[index]
			} else if f.Children[index].Label != "" {
				label = f.Children[index].Label
				f.Children[index].Label = ""
			} else {
				label = "Tab " + indexString
			}

			// Go!
			tab := b.Button().
				Type("button").
				Role("tab").
				ID("tab-"+indexString).
				Class("tab-label").
				Aria("controls", "panel-"+indexString).
				TabIndex("0")

			if index == 0 {
				tab.Aria("selected", "true")
			}

			tab.InnerHTML(label).Close()
		}

		b.Close() // role=tablist

		for index, child := range f.Children {
			indexString := convert.String(index)

			panel := b.Div().
				Role("tabpanel").
				ID("panel-"+indexString).
				Aria("labelledby", "tab-"+indexString)

			if index > 0 {
				panel.Attr("hidden", "true")
			}

			panel.EndBracket()

			if err := child.Write(library, s, v, b.SubTree()); err != nil {
				return derp.Wrap(err, "form.vocabulary.Tab", "Error writing child", child)
			}

			panel.Close() // role=tabpanel
		}

		b.Close() // role=tabs

		return nil
	})
}
