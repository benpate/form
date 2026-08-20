// Package form renders HTML forms from a JSON configuration, in the spirit of
// JSON-Forms (though not strictly compatible with it).
//
// A Form pairs a rosetta schema, which describes the shape of the data, with a
// tree of Elements, which describes the UI that edits it. Both halves can be
// written in Go or unmarshalled from JSON, so a form definition can live in a
// database alongside the records it edits:
//
//	f := form.New(
//		schema.New(schema.Object{
//			Properties: schema.ElementMap{
//				"name":  schema.String{Required: true},
//				"email": schema.String{Format: "email"},
//			},
//		}),
//		form.Element{
//			Type: "layout-vertical",
//			Children: []form.Element{
//				{Type: "text", Path: "name", Label: "Name"},
//				{Type: "text", Path: "email", Label: "Email"},
//			},
//		},
//	)
//
//	html, err := f.Editor(user, nil)
//
// Each Element names a Type, which is looked up in a package-level registry of
// Widgets. The widget subpackage ships the built-in set; call widget.UseAll once
// at startup to register them, or Use to add your own. Because the registry is a
// plain map, populate it during startup, before any goroutine renders a form.
//
// # The round trip
//
// Editor and Viewer render a form -- editable and read-only, respectively. An
// Element marked ReadOnly always renders read-only, even inside an Editor.
// SetURLValues completes the round trip, applying a posted url.Values back onto
// the original object through the schema, which filters out anything the schema
// does not allow. Validate checks a form definition against its schema and
// reports paths that do not exist; UnmarshalJSON calls it automatically.
//
// # Lookup codes
//
// Widgets that offer a fixed set of choices (select, radio, checkbox, and their
// relatives) draw them as LookupCodes -- value/label pairs that come from an
// "enum" option, from the schema's own enumeration, or from a LookupProvider
// named by the "provider" option. A provider whose group also implements
// WritableLookupGroup lets the user add a choice that was not on the list.
//
// # Escaping
//
// Widgets escape the data they render: field values and LookupCode labels are
// written with InnerText, so they may safely hold whatever an end-user typed.
//
// The form DEFINITION is trusted, and is NOT escaped. An Element's Description
// is written as raw markup on purpose, as is the Label of a layout element, and
// option values such as "style" and "script" are placed directly into the tag.
// Form definitions must therefore come from your own code or from an
// administrator -- never from an end-user.
package form
