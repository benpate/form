package vocabulary

import (
	"strings"

	"github.com/benpate/convert"
	"github.com/benpate/form"
	"github.com/benpate/path"
	"github.com/benpate/schema"
)

func Text(library form.Library) {

	library.Register("text", func(f form.Form, s schema.Schema, value interface{}, builder *strings.Builder) error {

		// Parse the path to the field value.
		p := path.New(f.Path)

		// If the schema is nil, then there's not much we can do here.
		if s != nil {
			s, _ = s.Path(p)
		}

		if s == nil {
			s = schema.Any{}
		}

		// Start building a new tag
		tag := TagBuilder("input", builder)

		// Always dd ID attribute (if values exist)
		tag.Attr("id", f.ID)

		// Try to find a value attribute
		if f.Path != "" {

			tag.Attr("name", f.Path)

			if value, err := p.Get(value); err == nil {

				if value, _ := convert.StringOk(value, ""); value != "" {
					tag.Attr("value", value)
				}
			}
		}

		// Add attributes that depend on what KIND of input we have.
		switch s := s.(type) {

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
