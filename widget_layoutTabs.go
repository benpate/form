package form

import (
	"strings"

	"github.com/benpate/derp"
	"github.com/segmentio/ksuid"

	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("layout-tabs", WidgetLayoutTabs{})
}

type WidgetLayoutTabs struct{}

func (WidgetLayoutTabs) View(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {
	return nil
}

func (WidgetLayoutTabs) Edit(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return WidgetLayoutTabs{}.View(element, s, lookupProvider, value, b)
	}

	if len(element.Label) > 0 {
		b.Div().Class("layout-title").InnerHTML(element.Label).Close()
	}

	// Make a placeholder for labels
	labels := make([]string, 0)

	// If we have a configuration option for labels,
	// parse it into a slice
	if labelString, ok := element.Options.GetString("labels"); ok && (labelString != "") {
		labels = strings.Split(labelString, ",")
	}

	b.Div().Class("tabs").Script("install TabContainer")
	b.Div().Role("tablist")

	for index := range element.Children {

		child := &element.Children[index]

		// Default ID for this child element
		if child.ID == "" {
			child.ID = ksuid.New().String() // TODO: LOW: Remove KSUID dependency from tabs (this is the only place)
		}

		var label string

		// Use the best label we have (configured, or generated)
		if index < len(labels) {
			label = labels[index]
		} else if child.Label != "" {
			label = child.Label
			child.Label = ""
		} else {
			label = "Tab " + convert.String(index)
		}

		// Go!
		tab := b.Button().
			Type("button").
			Role("tab").
			ID("tab-"+child.ID).
			Class("tab-label").
			Aria("controls", "panel-"+child.ID).
			TabIndex("0")

		if index == 0 {
			tab.Aria("selected", "true")
		}

		tab.InnerHTML(label).Close()
	}

	b.Close() // role=tablist

	for index, child := range element.Children {

		panel := b.Div().
			Role("tabpanel").
			ID("panel-"+child.ID).
			Aria("labelledby", "tab-"+child.ID)

		if index > 0 {
			panel.Attr("hidden", "true")
		}

		panel.EndBracket()

		if err := child.Edit(s, lookupProvider, value, b.SubTree()); err != nil {
			return derp.Wrap(err, "form.HTMLLayoutTabs", "Error writing child", child)
		}

		panel.Close() // role=tabpanel
	}

	b.Close() // role=tabs

	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (WidgetLayoutTabs) ShowLabels() bool {
	return false
}
