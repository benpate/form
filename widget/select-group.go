package widget

import (
	"encoding/json"
	"strings"

	"github.com/benpate/derp"
	"github.com/benpate/form"
	"github.com/benpate/form/groupie"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/schema"
	"github.com/benpate/rosetta/slice"
)

// SelectGroup renders two linked select boxes
type SelectGroup struct{}

func (widget SelectGroup) View(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	schemaElement := e.GetSchema(&f.Schema)
	valueString := e.GetString(value, &f.Schema)
	lookupCodes, _ := form.GetLookupCodes(e, schemaElement, provider)

	// Start building a new tag
	b.Div().Class("layout-value").EndBracket()
	for _, lookupCode := range lookupCodes {
		if lookupCode.Value == valueString {
			b.WriteString(lookupCode.Group)
			break
		}
	}

	b.Close()

	return nil
}

func (widget SelectGroup) Edit(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {

	const location = "form.widgets.SelectGroup.Edit"

	// find the path and schema to use
	schemaElement := e.GetSchema(&f.Schema)
	valueString := e.GetString(value, &f.Schema)

	if e, ok := schemaElement.(schema.Array); ok {
		schemaElement = e.Items
	}

	// Get all lookupCodes for this e...
	lookupCodes, _ := form.GetLookupCodes(e, schemaElement, provider)

	// Limit LookupCodes to ONLY the data we need.  No reason to send a HUGE file over JSON.
	lookupCodes = slice.Map(lookupCodes, func(lookupCode form.LookupCode) form.LookupCode {
		return form.LookupCode{
			Value: lookupCode.Value,
			Label: lookupCode.Label,
			Group: lookupCode.Group,
		}
	})

	// Clean up the element ID
	elementID := getElementID(e)

	selectBox := b.Container("select").
		ID(elementID).
		Name(e.Path).
		Aria("label", e.Label).
		Aria("description", e.Description).
		TabIndex("0")

	if err := widget.setChildWidget(f, e, lookupCodes, value, selectBox); err != nil {
		return derp.Wrap(err, location, "Unable to configure select-group widget.")
	}

	if isRequired(e, schemaElement) {
		selectBox.Attr("required", "true")
	}

	if focus, ok := e.Options.GetBoolOK("focus"); ok && focus {
		selectBox.Attr("autofocus", "true")
	}

	// Allow null options if not required
	if (schemaElement != nil) && (!schemaElement.IsRequired()) {
		b.Container("option").Value("").InnerText("").Close()
	}

	g := groupie.New()
	for _, lookupCode := range lookupCodes {
		if g.Header(lookupCode.Group) {

			opt := b.Container("option").Value(lookupCode.Group)
			if valueString == lookupCode.Group {
				opt.Attr("selected", "true")
			}
			opt.InnerText(lookupCode.Group).Close()
		}
	}

	b.CloseAll()
	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget SelectGroup) ShowLabels() bool {
	return true
}

func (widget SelectGroup) Encoding(_ *form.Element) string {
	return ""
}

/***********************************
 * Internal Helpers
 ***********************************/

func (widget SelectGroup) setChildWidget(f *form.Form, e *form.Element, lookupCodes []form.LookupCode, value any, selectBox *html.Element) error {

	const location = "form.widget.SelectGroup.setChildWidget"

	// Locate child selectbox
	children := e.Options.GetString("children")

	if children == "" {
		return derp.Internal(location, "Unable to link child widget because 'children' value is empty", e)
	}

	// Get the value of the child selectbox
	childValue, err := f.Schema.Get(value, children)

	if err != nil {
		return derp.Wrap(err, location, "Unable to retrieve child value from element", "element:", e, "children:", children, "value:", value)
	}

	// Marshal the LookupCodes into JSON
	lookupJSONbytes, err := json.Marshal(lookupCodes)

	if err != nil {
		return derp.Wrap(err, location, "Unable to marshal lookupCodes into JSON", e)
	}

	// Escape single quotes in the JSON string
	lookupJSON := string(lookupJSONbytes)
	lookupJSON = strings.ReplaceAll(lookupJSON, `\`, `\\`)
	lookupJSON = strings.ReplaceAll(lookupJSON, "'", "\\'")

	// Construct parameters for the hyperscript behavior
	hyperscript := "install SelectGroup(" +
		"children: '" + children + "', " +
		"options:'" + lookupJSON + "', " +
		"value:'" + convert.String(childValue) + "'" +
		")"

	// Add the behavior to the selectBox
	selectBox.Script(hyperscript)
	return nil
}
