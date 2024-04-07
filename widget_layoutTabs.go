package form

import (
	"strconv"
	"strings"

	"github.com/benpate/derp"

	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("layout-tabs", WidgetLayoutTabs{})
}

type WidgetLayoutTabs struct{}

func (widget WidgetLayoutTabs) View(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {
	return nil
}

func (widget WidgetLayoutTabs) Edit(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return WidgetLayoutTabs{}.View(element, s, lookupProvider, value, b)
	}

	if element.ID == "" {
		element.ID = "tabcontainer"
	}

	if len(element.Label) > 0 {
		b.Div().Class("layout-title").InnerText(element.Label).Close()
	}

	// Make a placeholder for labels
	labels := make([]string, 0)

	// If we have a configuration option for labels,
	// parse it into a slice
	if labelString, ok := element.Options.GetStringOK("labels"); ok && (labelString != "") {
		labels = strings.Split(labelString, ",")
	}

	b.Div().Class("tabs").Script("install TabContainer")
	b.Div().Role("tablist")

	for index, child := range element.Children {

		indexString := strconv.Itoa(index)
		child.ID = element.ID + "-" + indexString // Set ID for tab + panel

		var label string

		// Use the best label we have (configured, or generated)
		if index < len(labels) {
			label = labels[index]
		} else if child.Label != "" {
			label = child.Label
		} else {
			label = "Tab " + indexString
		}

		// Go!
		tab := b.Button().
			Type("button").
			Role("tab").
			ID("tab-"+child.ID).
			Class("tab-label").
			Aria("controls", "panel-"+child.ID).
			TabIndex("0")

		if script, ok := child.Options.GetStringOK("script"); ok {
			tab.Script(script)
		}

		if index == 0 {
			tab.Aria("selected", "true")
		}

		tab.InnerText(label).Close()
	}

	b.Close() // role=tablist

	for index, child := range element.Children {

		// Data overrides (just for this loop)
		child.ID = element.ID + "-" + strconv.Itoa(index) // Set ID for tab + panel
		child.Label = ""                                  // Remove remaining labels labels

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

func (widget WidgetLayoutTabs) ShowLabels() bool {
	return false
}

func (widget WidgetLayoutTabs) Encoding(element *Element) string {
	return collectEncoding(element.Children)
}
