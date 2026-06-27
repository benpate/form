# form/widget

This package supplies the concrete widgets for the [`form`](../README.md) package — the things that actually render an `Element` into HTML. Each widget is a tiny stateless type (a `struct{}`) that satisfies `form.Widget`: it knows how to draw the read-only `View` and the editable `Edit` for one kind of form control, plus a little metadata (`ShowLabels`, `ShowDescriptions`, `Encoding`). The parent `form` package owns the data model, schema binding, and the registry; this package only knows how to draw.

## The registry contract

Widgets are looked up by name from a package-level map in `form` (`form.Use(name, widget)` registers, `Element.Widget()` reads). Nothing here works until that map is populated.

- **Call `UseAll()` once at startup before drawing any form.** It registers all ~31 built-in widgets under their string keys (`"text"`, `"select"`, `"layout-tabs"`, …). An `Element.Type` that has no registered widget makes `Element.Widget()` return an error, not a zero widget.
- **The registry is init-time only.** It has no mutex, so finish all registration before the first concurrent render. This is the standard Go registration pattern (cf. `image.RegisterFormat`), not an oversight — don't "fix" it with a lock.
- **`UseAll` is the source of truth for the type→widget mapping.** When you add a widget, register it there; the name string in that file is the `type` value form definitions reference.

## What matters here

- **`View` is read-only, `Edit` is interactive — and they can diverge.** Most input widgets render a plain `layout-value` div in `View` and the real `<input>` in `Edit`. Some widgets (`Heading`, `HTML`, `Label`) render the same thing both ways; `LayoutTabs.View` deliberately renders *nothing* because tabs are an editing-only affordance. Keep that split intentional when adding a widget.
- **Three widgets emit raw, unescaped HTML on purpose: `HTML`, `Heading`, `Label` (via `InnerHTML(e.Description)`/`InnerHTML(e.Label)`), and `WYSIWYG.View` (via `WriteString`).** That content comes from the *form definition* or stored rich text, which is authoring-trusted input — not end-user form submissions. End-user *values* always go through `InnerText`/`Value`, which escape. If you add a widget, follow the same rule: never `InnerHTML` an end-user-supplied value.
- **`Encoding()` is how `multipart/form-data` propagates.** Only `Upload` returns a non-empty encoding; the layout/container widgets return `collectEncoding(children)` so a single upload anywhere in the tree promotes the whole form's enctype. A leaf widget that needs a special enctype must say so here or the form won't submit its files.
- **Layout widgets recurse through children; leaf widgets bind to a `Path`.** `Container`, `LayoutGroup`, `LayoutHorizontal`, `LayoutVertical`, and `LayoutTabs` draw `e.Children` (often via `b.SubTree()`); they have no `Path` of their own. Leaf widgets read/write the value at `e.Path` through `f.Schema`.
- **`Select`/`SelectGroup` support writable lookup providers via the `::NEWVALUE::` sentinel.** The hyperscript in `Select.Edit` and the constant `form.NewItemIdentifier` must stay in sync; the `form` package strips that prefix in `Element.replaceNewLookup` when saving. Changing the literal in one place silently breaks "add new option."
- **`groupie` drives `<optgroup>`/header boundaries.** `Multiselect` and `SelectGroup` use [`form/groupie`](../groupie/README.md) to emit a header only when the group changes; it assumes the lookup codes are already sorted by group.

## Shared helpers

`utils.go` holds the small internals every widget reuses: `iif` (ternary), `isRequired` (schema-or-option required check), and `getElementID` (derives a DOM id from `ID`/`Type`/`Path`, dot-to-dash). `utils_icons.go` holds icon helpers. Prefer these over re-deriving the same logic in a new widget.

## Known smell

`Label.View`/`Label.Edit` and `HTML.View`/`HTML.Edit` are byte-identical (SonarLint S4144). It's harmless duplication kept for now because the two methods are part of a stable interface and may legitimately diverge later; collapsing them isn't worth coupling the two interface methods. Don't be alarmed by the warning.
