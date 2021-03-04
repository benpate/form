package html

import (
	"html"

	"github.com/benpate/convert"
)

// Element represents a element that is being written into the provided strings.Builder
type Element struct {
	builder    *Builder
	parent     *Element
	name       string
	container  bool
	endBracket bool
	closed     bool
}

// Start writes the initial tag name and opening bracket for this tag.
func (element *Element) Start() *Element {
	// Grow the buffer and write the new tag name
	element.builder.Grow(len(element.name) + 2)
	element.builder.WriteRune('<')
	element.builder.WriteString(element.name)

	return element
}

// Attr writes the attribute into the string builder.  It converts
// the value (second parameter) into a string, and then uses html.EscapeString
// to escape the attribute value.  Attribute names ARE NOT escaped.
func (element *Element) Attr(name string, value interface{}) *Element {

	// If the element already has an end bracket, then we can't add any more attributes.
	if element.endBracket {
		return element
	}

	// this *should* already be a string, but just in case
	if valueString, _ := convert.StringOk(value, ""); valueString != "" {

		// escape the value
		valueString = html.EscapeString(valueString)

		// length of: space + name + quote + escaped value + quote
		element.builder.Grow(len(name) + len(valueString) + 4)

		// write values to the builder
		element.builder.WriteRune(' ')
		element.builder.WriteString(name)
		element.builder.WriteString(`="`)
		element.builder.WriteString(valueString)
		element.builder.WriteRune('"')
	}

	return element
}

// EndBracket writes the final ">" of the beginning element to the strings.Builder
// It uses an internal variable to prevent duplicate calls
func (element *Element) EndBracket() *Element {

	// If we already have an end bracket, then skip
	if element.endBracket {
		return element
	}

	// If this element is not a container, then this closes it permanently
	if element.container == false {
		element.closed = true
	}

	element.endBracket = true
	element.builder.Grow(1)
	element.builder.WriteRune('>')
	return element
}

// InnerHTML does three things:
// 1) closes the beginning element (if needed)
// 2) appends innerHTML (if provided)
// 3) writes an ending element to the builder (ie. </element> )
func (element *Element) InnerHTML(innerHTML string) *Element {

	// If the element has already been closed, then we cannot add anything more.
	if element.closed {
		return element
	}

	// Only need to write additional content if innerHTML is not empty
	if innerHTML != "" {

		if element.endBracket == false {
			element.builder.Grow(len(innerHTML) + 1)
			element.builder.WriteRune('>')
			element.endBracket = true
		} else {
			element.builder.Grow(len(innerHTML))
		}

		// Write innerHTML (if present)
		element.builder.WriteString(innerHTML)
	}

	// Write the remaining end element
	return element.Close()
}

// Close writes the necessary closing tag for this element and marks it closed
func (element *Element) Close() *Element {

	// If it's already been closed, then nothing else is required.
	if element.closed {
		return element
	}

	// Mark this element as closed.
	element.closed = true
	element.builder.last = element.parent

	// If this is not a CONTAINER element, then ensure that we have an end bracket
	if element.container == false {
		return element.EndBracket()
	}

	// If we already have an end bracket, then we need to add a full closing tag
	if element.endBracket {
		element.builder.Grow(len(element.name) + 3)
		element.builder.WriteString("</")
		element.builder.WriteString(element.name)
		element.builder.WriteRune('>')
		return element
	}

	// Otherwise, we can use the shortened "/>" syntax
	element.endBracket = true
	element.builder.Grow(3)
	element.builder.WriteString(" />")
	return element
}
