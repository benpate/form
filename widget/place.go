package widget

import (
	"net/url"
	"strings"

	"github.com/benpate/derp"
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/schema"
)

// Place is a form widget that presents an address lookup box. It returns
// a place name, latitude, and longitude.  If an endpoint URL is provided,
// then this widget will query the server for matching place names.
type Place struct{}

func (widget Place) View(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	valueString := widget.getString(e, &f.Schema, "name", value)

	// TODO: LOW: Apply formatting options?
	b.Div().Class("layout-value", e.Options.GetString("class")).InnerText(valueString).Close()
	return nil
}

func (widget Place) Edit(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {

	e.ID = strings.ReplaceAll(e.ID, ".", "_")

	searchID := e.ID + "_search"
	formattedID := e.ID + "_formatted"
	latitudeID := e.ID + "_latitude"
	longitudeID := e.ID + "_longitude"
	menuID := e.ID + "_menu"

	scripts := make([]string, 0)

	// Widget Wrapper + Positioning
	b.Div().
		ID(e.ID).
		Script("install PlaceSelect", "install Menu(input:#"+formattedID+")").
		Data("hx-indicator", "#"+e.ID)

	b.Div().Style("position:relative")

	// Complications
	b.Div().
		Style("position:absolute", "right:2px", "top:2px", "bottom:2px", "padding:var(--rhythm)", "background-color:var(--input-background)").
		EndBracket()

	{
		// Loading Icon
		b.Span().
			Class("PlaceSelectLoading htmx-request-show").
			Style("color:var(--gray50)")

		loadingIcon(b)
		b.Close()

		// PlaceSelect Icon
		b.Span().
			Class("PlaceSelectIcon").
			Style("display:none", "cursor:pointer", "color:var(--gray50)")

		locateIcon(b)
		b.Close()
	}

	b.Close()

	// Latitude (hidden)
	b.Input("hidden", e.Path+".latitude").
		ID(latitudeID).
		Class("PlaceSelectLatitude").
		Value(widget.getString(e, &f.Schema, "latitude", value)).
		Close()

	// Longitude (hidden)
	b.Input("hidden", e.Path+".longitude").
		ID(longitudeID).
		Class("PlaceSelectLongitude").
		Value(widget.getString(e, &f.Schema, "longitude", value)).
		Close()

	// Formatted (hidden)
	b.Input("hidden", e.Path+".formatted").
		ID(formattedID).
		Class("PlaceSelectFormatted").
		Value(widget.getString(e, &f.Schema, "formatted", value)).
		Close()

	// Search Box
	tag := b.Input("text", "q").
		ID(searchID).
		Class("PlaceSelectSearch").
		Value(widget.getString(e, &f.Schema, "formatted", value)).
		Aria("label", e.Label).
		Aria("description", e.Description).
		TabIndex("0").
		Attr("autocomplete", "off").
		Attr("data-1p-ignore", "true")

	if focus, ok := e.Options.GetBoolOK("focus"); ok && focus {
		tag.Attr("autofocus", "true")
	}

	if placeholder := e.Options.GetString("placeholder"); placeholder != "" {
		tag.Attr("placeholder", placeholder)
	}

	// Add attributes that depend on what KIND of input we have.
	switch schemaElement := e.GetSchema(&f.Schema); s := schemaElement.(type) {

	case schema.String:
		if s.Required || e.Options.GetBool("required") {
			tag.Attr("required", "true")
		}

		if s.RequiredIf != "" {
			scripts = append(scripts, "install requiredIf(condition:'"+s.RequiredIf+"')")
		} else if requiredIf := e.Options.GetString("required-if"); requiredIf != "" {
			scripts = append(scripts, "install requiredIf(condition:'"+requiredIf+"')")
		}

	default:
		tag.Type("text")
	}

	endpoint := e.Options.GetString("endpoint")

	if endpoint != "" {
		tag.Attr("list", menuID)
		tag.Data("hx-get", endpoint)
		tag.Data("hx-vals", `js:{longitude:localStorage.getItem('longitude'), latitude:localStorage.getItem('latitude')}`)
		tag.Data("hx-trigger", "keyup changed throttle:200ms")
		tag.Data("hx-target", "#"+menuID)
		tag.Data("hx-swap", "innerHTML")
		tag.Data("hx-push-url", "false")
	}

	if len(scripts) > 0 {
		tag.Script(strings.Join(scripts, " "))
	}

	b.Close()
	b.Close()

	if endpoint != "" {
		b.Div().
			ID(menuID).
			Role("menu").
			Style("display:none").
			Close()
	}

	b.CloseAll()
	return nil
}

/***********************************
 * Custom Setter
 ***********************************/

func (widget Place) SetURLValue(form *form.Form, element *form.Element, object any, values url.Values) error {

	const location = "form.widget.Place.SetURLValue"

	// Set Full Name
	formattedPath := element.Path + ".formatted"
	formatted := values.Get(formattedPath)

	if err := form.Schema.Set(object, formattedPath, formatted); err != nil {
		return derp.Wrap(err, location, "Unable to set formatted", formattedPath, formatted)
	}

	// Set Longitude
	longitudePath := element.Path + ".longitude"
	longitude := values.Get(longitudePath)

	if err := form.Schema.Set(object, longitudePath, longitude); err != nil {
		return derp.Wrap(err, location, "Unable to set longitude", longitudePath, longitude)
	}

	// Set Latitude
	latitudePath := element.Path + ".latitude"
	latitude := values.Get(latitudePath)

	if err := form.Schema.Set(object, latitudePath, latitude); err != nil {
		return derp.Wrap(err, location, "Unable to set latitude", latitudePath, latitude)
	}

	// Success!
	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget Place) ShowLabels() bool {
	return true
}

func (widget Place) Encoding(_ *form.Element) string {
	return ""
}

/***********************************
 * Helper Methods
 ***********************************/

func (widget Place) getString(element *form.Element, schema *schema.Schema, subPath string, value any) string {

	path := element.Path + "." + subPath

	if result, err := schema.Get(value, path); err == nil {
		return convert.String(result)
	}

	return ""
}
