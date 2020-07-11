package vocabulary

import "github.com/benpate/form"

func Text(library form.Library) {
	RegisterTemplate(library, "text", `<input type="text" name="{{.Form.Path}}" value="{{.Value}}" />`)
}
