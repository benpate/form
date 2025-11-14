package widget

import (
	"strings"

	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
	"github.com/benpate/rosetta/slice"
)

// Select renders a select box widget
type Select struct{}

func (widget Select) View(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	schemaElement := e.GetSchema(&f.Schema)
	valueString := e.GetString(value, &f.Schema)
	lookupCodes, _ := form.GetLookupCodes(e, schemaElement, provider)

	// Start building a new tag
	b.Div().Class("layout-value").EndBracket()
	for _, lookupCode := range lookupCodes {
		if lookupCode.Value == valueString {
			b.WriteString(lookupCode.Label)
			break
		}
	}

	b.Close()

	return nil
}

func (widget Select) Edit(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	schemaElement := e.GetSchema(&f.Schema)
	valueSlice := e.GetSliceOfString(value, &f.Schema)

	if e, ok := schemaElement.(schema.Array); ok {
		schemaElement = e.Items
	}

	// Get all lookupCodes for this e...
	lookupCodes, isWritable := form.GetLookupCodes(e, schemaElement, provider)

	elementID := e.ID

	if elementID == "" {
		elementID = "select-" + strings.ReplaceAll(e.Path, ".", "-")
	}

	selectBox := b.Container("select").
		ID(elementID).
		Name(e.Path).
		Aria("label", e.Label).
		Aria("description", e.Description).
		TabIndex("0")

	if isRequired(e, schemaElement) {
		selectBox.Attr("required", "true")
	}

	if focus, ok := e.Options.GetBoolOK("focus"); ok && focus {
		selectBox.Attr("autofocus", "true")
	}

	// Support for writable lookup providers
	if isWritable {
		selectBox.Script(`on change if my value is '::NEWVALUE::' then 
			set newName to window.prompt("Add New Value")
			if newName is not null then 
				make an Option from newName, ("::NEWVALUE::" + newName), true, true called newOption 
				call me.add(newOption)
			end`)
	}

	// Allow null options if not required
	if (schemaElement != nil) && (!schemaElement.IsRequired()) {
		b.Container("option").Value("").InnerText("").Close()
	}

	group := ""
	// Display all lookup codes
	for index, lookupCode := range lookupCodes {

		// Begin <optgroup> tags (if needed)
		if lookupCode.Group != group {
			group = lookupCode.Group
			b.Container("optgroup").Label(group).EndBracket()
		}

		// Render the <option> tag
		opt := b.Container("option").Value(lookupCode.Value)
		if slice.Contains(valueSlice, lookupCode.Value) {
			opt.Attr("selected", "true")
		}
		opt.InnerText(lookupCode.Label).Close()

		// Close the <optgroup> tag (if needed)
		if group != widget.groupValue(lookupCodes, index+1) {
			b.Close()
		}
	}

	// Support for writable lookup providers
	if isWritable {
		b.Container("option").
			Class("add-new").
			Value("::NEWVALUE::").
			InnerHTML("+ Add Another...").
			Close()
	}

	b.CloseAll()
	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget Select) ShowLabels() bool {
	return true
}

func (widget Select) Encoding(_ *form.Element) string {
	return ""
}

func (widget Select) groupValue(lookupCodes []form.LookupCode, index int) string {
	if index >= len(lookupCodes) {
		return ""
	}

	return lookupCodes[index].Group
}
