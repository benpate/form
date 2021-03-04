package vocabulary

import (
	"github.com/benpate/form"
	"github.com/benpate/form/html"
	"github.com/benpate/schema"
)

// Text registers a text <input> widget into the library
func Text(library form.Library) {

	library.Register("text", func(f form.Form, s *schema.Schema, v interface{}, b *html.Builder) error {

		// find the path and schema to use
		schemaObject, valueString := locateSchema(f.Path, s, v)

		// Start building a new tag
		tag := b.Input().
			ID(f.ID).
			Name(f.Path).
			Value(valueString)

		// Add attributes that depend on what KIND of input we have.
		switch s := schemaObject.(type) {

		case schema.Integer:
			tag.Attr("type", "number").Attr("step", "1")

			if s.Minimum.IsPresent() {
				tag.Attr("min", s.Minimum.String())
			}

			if s.Maximum.IsPresent() {
				tag.Attr("max", s.Maximum.String())
			}

			if s.Required {
				tag.Attr("required", "true")
			}

		case schema.Number:

			tag.Attr("type", "number")

			if s.Minimum.IsPresent() {
				tag.Attr("min", s.Minimum.String())
			}

			if s.Maximum.IsPresent() {
				tag.Attr("max", s.Maximum.String())
			}

			if s.Required {
				tag.Attr("required", true)
			}

		case schema.String:

			switch s.Format {
			case "email":
				tag.Attr("type", "email")
			case "tel":
				tag.Attr("type", "tel")
			case "url":
				tag.Attr("type", "url")
			default:
				tag.Attr("type", "text")
			}

			if s.MinLength.IsPresent() {
				tag.Attr("minlength", s.MinLength.Int())
			}

			if s.MaxLength.IsPresent() {
				tag.Attr("maxlength", s.MaxLength.Int())
			}

			if s.Pattern != "" {
				tag.Attr("pattern", s.Pattern)
			}

			if s.Required {
				tag.Attr("required", true)
			}

		default:
			tag.Attr("type", "text")
		}

		if f.CSSClass != "" {
			tag.Attr("class", f.CSSClass)
		}

		if f.Description != "" {
			tag.Attr("hint", f.Description)
		}

		tag.Close()
		return nil
	})
}

/*


   <input type="button">
   <input type="checkbox">
   <input type="color">
   <input type="date">
   <input type="datetime-local">
   <input type="email">
   <input type="file">
   <input type="hidden">
   <input type="image">
   <input type="month">
   <input type="number">
   <input type="password">
   <input type="radio">
   <input type="range">
   <input type="reset">
   <input type="search">
   <input type="submit">
   <input type="tel">
   <input type="text">
   <input type="time">
   <input type="url">
   <input type="week">
*/
