package vocabulary

import (
	"html"
	"strings"

	"github.com/benpate/convert"
)

// Tag represents a tag that is being written into the provided strings.Builder
type Tag struct {
	name    string
	builder *strings.Builder
	closed  bool
}

// TagBuilder returns a Tag that can be written into the provided srings.Builder
func TagBuilder(name string, builder *strings.Builder) Tag {

	result := Tag{
		name:    name,
		builder: builder,
		closed:  false,
	}

	// grow length of "<" + name + ">"
	result.builder.Grow(len(name) + 2)
	result.builder.WriteRune('<')
	result.builder.WriteString(name)
	return result
}

// Attr writes the attribute into the string builder.  It converts
// the value (second parameter) into a string, and then uses html.EscapeString
// to escape the attribute value.  Attribute names ARE NOT escaped.
func (tag Tag) Attr(name string, value interface{}) Tag {

	// this *should* already be a string, but just in case
	if valueString, _ := convert.StringOk(value, ""); valueString != "" {

		// escape the value
		valueString = html.EscapeString(valueString)

		// length of: space + name + quote + escaped value + quote
		tag.builder.Grow(len(name) + len(valueString) + 4)

		// write values to the builder
		tag.builder.WriteRune(' ')
		tag.builder.WriteString(name)
		tag.builder.WriteString(`="`)
		tag.builder.WriteString(valueString)
		tag.builder.WriteRune('"')
	}

	return tag
}

// EndTag does three things:
// 1) closes the beginning tag (if needed)
// 2) appends innerHTML (if provided)
// 3) writes an ending tag to the builder (ie. </tag> )
func (tag Tag) EndTag(innerHTML string) {

	growSize := len(tag.name) + 3

	if tag.closed == false {
		tag.builder.Grow(growSize + 1)
		tag.builder.WriteRune('>')
		tag.closed = true
	} else {
		tag.builder.Grow(growSize)
	}

	if innerHTML != "" {
		tag.builder.WriteString(innerHTML)
	}

	tag.builder.WriteString("</")
	tag.builder.WriteString(tag.name)
	tag.builder.WriteRune('>')
}

// Close writes the final ">" of the beginning tag to the strings.Builder
// It uses an internal variable to prevent duplicate calls
func (tag Tag) Close() {

	if tag.closed == false {
		tag.closed = true
		tag.builder.WriteRune('>')
	}
}
