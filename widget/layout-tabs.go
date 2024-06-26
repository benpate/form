package widget

import (
	"strconv"
	"strings"

	"github.com/benpate/derp"
	"github.com/benpate/form"

	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

type LayoutTabs struct{}

func (widget LayoutTabs) View(element *form.Element, s *schema.Schema, provider form.LookupProvider, value any, b *html.Builder) error {
	return nil
}

func (widget LayoutTabs) Edit(element *form.Element, s *schema.Schema, provider form.LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return LayoutTabs{}.View(element, s, provider, value, b)
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

		if err := child.Edit(s, provider, value, b.SubTree()); err != nil {
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

func (widget LayoutTabs) ShowLabels() bool {
	return false
}

func (widget LayoutTabs) Encoding(element *form.Element) string {
	return collectEncoding(element.Children)
}
