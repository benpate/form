package vocabulary

import "github.com/benpate/form"

func All(library *form.Library) {
	Checkbox(library)
	LayoutGroup(library)
	LayoutHorizontal(library)
	LayoutVertical(library)
	Multiselect(library)
	Option(library)
	Select(library)
	Tab(library)
	Text(library)
	Textarea(library)
	Toggle(library)
	WYSIWYG(library)
}
