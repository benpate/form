package vocabulary

import (
	"html"
	"strings"

	"github.com/benpate/derp"
	"github.com/benpate/form"
	"github.com/benpate/schema"
)

func LayoutVertical(library form.Library) {

	library.Register("layout-table", func(form form.Form, schema schema.Schema, value interface{}, builder *strings.Builder) error {

		var result error

		builder.WriteString(`<div class="layout-vertical">`)

		builder.WriteString(`<div class="label">` + html.EscapeString(form.Label) + `</div>`)
		builder.WriteString(`<div class="elements>`)

		for index, child := range form.Children {
			builder.WriteString(`<div class="element">`)

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
