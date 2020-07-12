package vocabulary

import (
	"html"
	"strings"

	"github.com/benpate/derp"
	"github.com/benpate/form"
	"github.com/benpate/schema"
)

func LayoutVertical(library form.Library) {

	library.Register("layout-vertical", func(form form.Form, schema schema.Schema, value interface{}, builder *strings.Builder) error {

		var result error

		builder.WriteString(`<div class="layout-vertical">`)

		if len(form.Label) > 0 {
			builder.WriteString(`<div class="label">`)
			builder.WriteString(html.EscapeString(form.Label))
			builder.WriteString(`</div>`)
			builder.WriteString(`<div class="elements>`)
		}

		for index, child := range form.Children {

			builder.WriteString(`<div class="element">`)

			TagBuilder("label", builder).Attr("for", child.ID).EndTag(child.Label)

			if err := child.Write(library, schema, value, builder); err != nil {
				result = derp.Wrap(err, "form.widget.LayoutVertical", "Error rendering child", index, form)
			}
			builder.WriteString(`</div>`)
		}

		builder.WriteString(`</div>`)
		builder.WriteString(`</div>`)

		return result
	})
}
