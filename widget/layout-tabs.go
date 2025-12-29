package widget

import (
	"strconv"
	"strings"

	"github.com/benpate/derp"
	"github.com/benpate/form"

	"github.com/benpate/html"
)

type LayoutTabs struct{}

func (widget LayoutTabs) View(_ *form.Form, _ *form.Element, _ form.LookupProvider, _ any, _ *html.Builder) error {
	return nil
}

func (widget LayoutTabs) Edit(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {

	if e.ID == "" {
		e.ID = "tabcontainer"
	}

	if len(e.Label) > 0 {
		b.Div().Class("layout-title").InnerText(e.Label).Close()
	}

	// Make a placeholder for labels
	labels := make([]string, 0)

	// If we have a configuration option for labels,
	// parse it into a slice
	if labelString, ok := e.Options.GetStringOK("labels"); ok && (labelString != "") {
		labels = strings.Split(labelString, ",")
	}

	b.Div().Class("tabs").Script("install TabContainer")
	b.Div().Role("tablist")

	selectedIndex := f.OptionInt("selected-tab")

	for index, child := range e.Children {

		indexString := strconv.Itoa(index)
		child.ID = e.ID + "-" + indexString // Set ID for tab + panel

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

		if index == selectedIndex {
			tab.Aria("selected", "true")
		}

		tab.InnerText(label).Close()
	}

	b.Close() // role=tablist

	for index, child := range e.Children {

		// Data overrides (just for this loop)
		child.ID = e.ID + "-" + strconv.Itoa(index) // Set ID for tab + panel
		child.Label = ""                            // Remove remaining labels labels

		panel := b.Div().
			Role("tabpanel").
			ID("panel-"+child.ID).
			Aria("labelledby", "tab-"+child.ID)

		if index > 0 {
			panel.Attr("hidden", "true")
		}

		panel.EndBracket()

		if err := child.Edit(f, provider, value, b.SubTree()); err != nil {
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

func (widget LayoutTabs) Encoding(e *form.Element) string {
	return collectEncoding(e.Children)
}
