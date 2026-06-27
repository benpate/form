# form/groupie

A one-type helper for detecting group boundaries while iterating an already-sorted list. It exists so the [`widget`](../widget/README.md) select widgets can emit a group header (an `<optgroup>` or a header row) exactly once per group instead of on every row. See the [`form`](../README.md) package for the bigger picture.

## What matters here

- **It only detects *changes*, not *groups*.** `Header(value)` returns `true` the first time it sees a value and whenever the value differs from the previous call. It does **not** sort or deduplicate. If the input isn't already sorted by the grouping key, the same group will trigger multiple headers — sort first.
- **State is per-iteration; make a fresh `Groupie` each loop.** Each `Groupie` remembers exactly one previous value. Reusing one across two independent renders would carry the last value over and suppress the first header of the second loop. The widgets call `groupie.New()` at the top of every draw for this reason.
- **Comparison is `==` on `any`.** Values must be comparable (strings, in practice). Passing an uncomparable type (slice, map) would panic — callers only ever pass the string `Group` field.
