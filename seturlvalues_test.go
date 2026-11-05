package form

import (
	"net/url"
	"strconv"
	"sync"
	"testing"

	"github.com/benpate/derp"
	"github.com/benpate/rosetta/mapof"
	"github.com/benpate/rosetta/schema"
	"github.com/benpate/rosetta/sliceof"
	"github.com/stretchr/testify/require"
)

func TestSetURLValues_Basic(t *testing.T) {

	useTestWidget()

	s := schema.New(schema.Object{
		Properties: schema.ElementMap{
			"name": schema.String{},
		},
	})

	form := New(s, Element{Type: "test", Path: "name"})

	object := mapof.Any{}
	values := url.Values{"name": []string{"Sarah"}}

	require.NoError(t, form.SetURLValues(&object, values, nil))
	require.Equal(t, "Sarah", object.GetString("name"))
}

func TestSetURLValues_SkipsReadOnly(t *testing.T) {

	useTestWidget()

	s := schema.New(schema.Object{
		Properties: schema.ElementMap{
			"name": schema.String{},
		},
	})

	form := New(s, Element{Type: "test", Path: "name", ReadOnly: true})

	object := mapof.Any{"name": "original"}
	values := url.Values{"name": []string{"changed"}}

	require.NoError(t, form.SetURLValues(&object, values, nil))
	// Read-only fields are never updated
	require.Equal(t, "original", object.GetString("name"))
}

func TestSetURLValues_ShowIf(t *testing.T) {

	useTestWidget()

	s := schema.New(schema.Object{
		Properties: schema.ElementMap{
			"type":    schema.String{},
			"comment": schema.String{},
		},
	})

	// "comment" is only visible (and writable) when type == "detailed"
	form := New(s, Element{
		Children: []Element{
			{Type: "test", Path: "type"},
			{Type: "test", Path: "comment", Options: mapof.Template{"show-if": "type is detailed"}},
		},
	})

	// When the condition is NOT met, the dependent field is not written
	{
		object := mapof.Any{"type": "simple"}
		values := url.Values{"type": []string{"simple"}, "comment": []string{"hello"}}
		require.NoError(t, form.SetURLValues(&object, values, nil))
		require.Equal(t, "", object.GetString("comment"))
	}

	// When the condition IS met, the dependent field is written
	{
		object := mapof.Any{"type": "detailed"}
		values := url.Values{"type": []string{"detailed"}, "comment": []string{"hello"}}
		require.NoError(t, form.SetURLValues(&object, values, nil))
		require.Equal(t, "hello", object.GetString("comment"))
	}
}

// --- Array-typed paths ---
// Multi-value widgets (multiselect, check-button-group) post plain []string via url.Values,
// but rosetta validates Array schemas through its ArrayGetterSetter interface, which the
// builtin slice does not implement.  These tests pin the *sliceof.String bridge in
// schemaSafeValue; without it, every array save fails schema validation silently.

type arrayTestObject struct {
	Channels sliceof.String
}

func (o *arrayTestObject) GetPointer(name string) (any, bool) {
	if name == "channels" {
		return &o.Channels, true
	}
	return nil, false
}

func arrayTestSchema() schema.Schema {
	return schema.New(schema.Object{
		Properties: schema.ElementMap{
			"channels": schema.Array{Items: schema.String{Enum: []string{"ALPHA", "BETA", "GAMMA"}}},
		},
	})
}

func TestSetURLValues_ArrayPath(t *testing.T) {

	useTestWidget()

	form := New(arrayTestSchema(), Element{Type: "test", Path: "channels"})

	object := arrayTestObject{}
	values := url.Values{"channels": []string{"ALPHA", "GAMMA"}}

	require.NoError(t, form.SetURLValues(&object, values, nil))
	require.Equal(t, sliceof.String{"ALPHA", "GAMMA"}, object.Channels)
}

func TestSetURLValues_ArrayPath_ClearsWhenAbsent(t *testing.T) {

	useTestWidget()

	form := New(arrayTestSchema(), Element{Type: "test", Path: "channels"})

	// No posted values for the path (nothing checked) clears the array.
	object := arrayTestObject{Channels: sliceof.String{"ALPHA"}}
	require.NoError(t, form.SetURLValues(&object, url.Values{}, nil))
	require.Equal(t, sliceof.String{}, object.Channels)
}

func TestSetURLValues_ArrayPath_RejectsInvalidEnum(t *testing.T) {

	useTestWidget()

	form := New(arrayTestSchema(), Element{Type: "test", Path: "channels"})

	// Values outside the Enum fail schema validation; SetURLValues logs and
	// continues (schema-filtered data never reaches the object).
	object := arrayTestObject{Channels: sliceof.String{"ALPHA"}}
	values := url.Values{"channels": []string{"BOGUS"}}

	require.NoError(t, form.SetURLValues(&object, values, nil))
	require.Equal(t, sliceof.String{"ALPHA"}, object.Channels)
}

// --- replaceNewLookup with a writable provider ---

type mockProvider struct {
	group LookupGroup
}

func (p mockProvider) Group(string) LookupGroup {
	return p.group
}

type mockWritableGroup struct {
	added []string
}

func (g *mockWritableGroup) Get() []LookupCode { return nil }

func (g *mockWritableGroup) Add(name string) (string, error) {
	g.added = append(g.added, name)
	return "generated-id", nil
}

func TestReplaceNewLookup_Writable(t *testing.T) {

	group := &mockWritableGroup{}
	provider := mockProvider{group: group}

	element := Element{Options: mapof.Template{"provider": "things"}}

	value, updated, err := element.replaceNewLookup(provider, NewItemIdentifier+"My New Thing")

	require.NoError(t, err)
	require.True(t, updated)
	require.Equal(t, "generated-id", value)
	require.Equal(t, []string{"My New Thing"}, group.added)
}

func TestReplaceNewLookup_NilProvider(t *testing.T) {
	element := Element{}
	value, updated, err := element.replaceNewLookup(nil, NewItemIdentifier+"x")
	require.NoError(t, err)
	require.False(t, updated)
	require.Equal(t, NewItemIdentifier+"x", value)
}

func TestReplaceNewLookup_NotNewValue(t *testing.T) {
	element := Element{Options: mapof.Template{"provider": "things"}}
	provider := mockProvider{group: &mockWritableGroup{}}

	// A value without the new-item prefix is passed through unchanged
	value, updated, err := element.replaceNewLookup(provider, "existing-value")
	require.NoError(t, err)
	require.False(t, updated)
	require.Equal(t, "existing-value", value)
}

func TestReplaceNewLookup_NoProviderOption(t *testing.T) {
	element := Element{Options: mapof.Template{}}
	provider := mockProvider{group: &mockWritableGroup{}}

	value, updated, err := element.replaceNewLookup(provider, NewItemIdentifier+"x")
	require.NoError(t, err)
	require.False(t, updated)
	require.Equal(t, NewItemIdentifier+"x", value)
}

func TestReplaceNewLookup_ReadOnlyGroup(t *testing.T) {

	// A non-writable group cannot accept new values
	element := Element{Options: mapof.Template{"provider": "things"}}
	provider := mockProvider{group: NewReadOnlyLookupGroup()}

	value, updated, err := element.replaceNewLookup(provider, NewItemIdentifier+"x")
	require.NoError(t, err)
	require.False(t, updated)
	require.Equal(t, NewItemIdentifier+"x", value)
}

// TestReplaceNewLookup_BlankName refuses a blank name, whether or not the group could store it.
func TestReplaceNewLookup_BlankName(t *testing.T) {

	element := Element{Path: "circle", Options: mapof.Template{"provider": "things"}}
	writable := &mockWritableGroup{}

	test := func(provider LookupProvider, posted string) {
		t.Helper()
		value, updated, err := element.replaceNewLookup(provider, posted)
		require.True(t, derp.IsValidationError(err), "provider %T, posted %q: %v", provider, posted, err)
		require.False(t, updated, "provider %T, posted %q", provider, posted)
		require.Equal(t, posted, value, "provider %T, posted %q", provider, posted)
	}

	for _, posted := range []string{NewItemIdentifier, NewItemIdentifier + "   ", NewItemIdentifier + "\t\n"} {
		test(nil, posted)
		test(mockProvider{group: NewReadOnlyLookupGroup()}, posted)
		test(mockProvider{group: writable}, posted)
	}

	require.Empty(t, writable.added, "a blank name must never reach Add")
}

// TestReplaceNewLookup_KeepsNameAsTyped passes a name with surrounding spaces to Add unchanged.
func TestReplaceNewLookup_KeepsNameAsTyped(t *testing.T) {

	group := &mockWritableGroup{}
	element := Element{Options: mapof.Template{"provider": "things"}}

	_, updated, err := element.replaceNewLookup(mockProvider{group: group}, NewItemIdentifier+" Book Club ")
	require.NoError(t, err)
	require.True(t, updated)
	require.Equal(t, []string{" Book Club "}, group.added)
}

// TestSetURLValues_BlankNewValue fails the save on a blank "add another", before it reaches Add.
func TestSetURLValues_BlankNewValue(t *testing.T) {

	useTestWidget()

	s := schema.New(schema.Object{Properties: schema.ElementMap{
		"name":   schema.String{},
		"circle": schema.String{},
	}})

	form := New(s, Element{Children: []Element{
		{Type: "test", Path: "name"},
		{Type: "test", Path: "circle", Options: mapof.Template{"provider": "things"}},
	}})

	group := &mockWritableGroup{}
	object := mapof.Any{"name": "Sarah", "circle": "stored"}
	values := url.Values{"name": {"Sam"}, "circle": {NewItemIdentifier + " "}}

	err := form.SetURLValues(&object, values, mockProvider{group: group})
	require.True(t, derp.IsValidationError(err), "a blank name must fail the save as invalid input: %v", err)
	require.Empty(t, group.added)
	require.Equal(t, "stored", object.GetString("circle"))
}

// --- Elements the form does not show ---

// TestSetURLValues_ShowIfOption writes a field only when its Form option is TRUE.
func TestSetURLValues_ShowIfOption(t *testing.T) {

	useTestWidget()

	s := schema.New(schema.Object{Properties: schema.ElementMap{
		"name":  schema.String{},
		"admin": schema.String{},
	}})

	element := Element{Children: []Element{
		{Type: "test", Path: "name"},
		{Type: "test", Path: "admin", Options: mapof.Template{"show-if-option": "is-admin"}},
	}}

	test := func(expected string, options ...string) {
		t.Helper()
		form := New(s, element, options...)
		object := mapof.Any{"name": "Sarah", "admin": "stored"}
		values := url.Values{"name": {"Sam"}, "admin": {"posted"}}
		require.NoError(t, form.SetURLValues(&object, values, nil))
		require.Equal(t, "Sam", object.GetString("name"), "options %v", options)
		require.Equal(t, expected, object.GetString("admin"), "options %v", options)
	}

	// A missing option counts as FALSE, so a caller that drops its options hides the field
	test("stored", "is-admin:false")
	test("stored")
	test("posted", "is-admin:true")
}

// TestSetURLValues_HideIfOption skips a field when its Form option is TRUE.
func TestSetURLValues_HideIfOption(t *testing.T) {

	useTestWidget()

	s := schema.New(schema.Object{Properties: schema.ElementMap{
		"guest": schema.String{},
	}})

	element := Element{Children: []Element{
		{Type: "test", Path: "guest", Options: mapof.Template{"hide-if-option": "is-member"}},
	}}

	test := func(expected string, options ...string) {
		t.Helper()
		form := New(s, element, options...)
		object := mapof.Any{"guest": "stored"}
		require.NoError(t, form.SetURLValues(&object, url.Values{"guest": {"posted"}}, nil))
		require.Equal(t, expected, object.GetString("guest"), "options %v", options)
	}

	test("stored", "is-member:true")
	test("posted", "is-member:false")
	test("posted")
}

// TestSetURLValues_OptionRemovesSubtree skips every field inside a container that an option removes.
func TestSetURLValues_OptionRemovesSubtree(t *testing.T) {

	useTestWidget()

	s := schema.New(schema.Object{Properties: schema.ElementMap{
		"admin": schema.String{},
		"deep":  schema.String{},
	}})

	element := Element{Children: []Element{
		{Options: mapof.Template{"show-if-option": "is-admin"}, Children: []Element{
			{Type: "test", Path: "admin"},
			{Children: []Element{{Type: "test", Path: "deep"}}},
		}},
	}}

	test := func(expected string, options ...string) {
		t.Helper()
		form := New(s, element, options...)
		object := mapof.Any{"admin": "stored", "deep": "stored"}
		values := url.Values{"admin": {"posted"}, "deep": {"posted"}}
		require.NoError(t, form.SetURLValues(&object, values, nil))
		require.Equal(t, expected, object.GetString("admin"), "options %v", options)
		require.Equal(t, expected, object.GetString("deep"), "options %v", options)
	}

	test("stored", "is-admin:false")
	test("posted", "is-admin:true")
}

// TestSetURLValues_ReadOnlyContainer skips every field inside a read-only container.
func TestSetURLValues_ReadOnlyContainer(t *testing.T) {

	useTestWidget()

	s := schema.New(schema.Object{Properties: schema.ElementMap{
		"name":   schema.String{},
		"locked": schema.String{},
	}})

	form := New(s, Element{Children: []Element{
		{Type: "test", Path: "name"},
		{ReadOnly: true, Children: []Element{{Type: "test", Path: "locked"}}},
	}})

	object := mapof.Any{"name": "Sarah", "locked": "stored"}
	values := url.Values{"name": {"Sam"}, "locked": {"posted"}}
	require.NoError(t, form.SetURLValues(&object, values, nil))
	require.Equal(t, "Sam", object.GetString("name"))
	require.Equal(t, "stored", object.GetString("locked"))
}

// TestSetURLValues_ContainerShowIf skips every field inside a container whose "show-if" fails,
// judged against the values posted alongside them.
func TestSetURLValues_ContainerShowIf(t *testing.T) {

	useTestWidget()

	s := schema.New(schema.Object{Properties: schema.ElementMap{
		"role":   schema.String{},
		"secret": schema.String{},
		"note":   schema.String{},
	}})

	form := New(s, Element{Children: []Element{
		{Options: mapof.Template{"show-if": "role is admin"}, Children: []Element{
			{Type: "test", Path: "secret"},
			{Children: []Element{{Type: "test", Path: "note"}}},
		}},
		{Type: "test", Path: "role"},
	}})

	test := func(role string, expected string) {
		t.Helper()
		object := mapof.Any{"role": "user", "secret": "stored", "note": "stored"}
		values := url.Values{"role": {role}, "secret": {"posted"}, "note": {"posted"}}
		require.NoError(t, form.SetURLValues(&object, values, nil))
		require.Equal(t, role, object.GetString("role"))
		require.Equal(t, expected, object.GetString("secret"), "role %s", role)
		require.Equal(t, expected, object.GetString("note"), "role %s", role)
	}

	// "role" comes after the container, and is still written first
	test("user", "stored")
	test("admin", "posted")
}

// TestSetURLValues_NestedShowIf writes a field only when its own "show-if" and its container's both pass.
func TestSetURLValues_NestedShowIf(t *testing.T) {

	useTestWidget()

	s := schema.New(schema.Object{Properties: schema.ElementMap{
		"role":  schema.String{},
		"level": schema.String{},
		"note":  schema.String{},
	}})

	form := New(s, Element{Children: []Element{
		{Type: "test", Path: "role"},
		{Type: "test", Path: "level"},
		{Options: mapof.Template{"show-if": "role is admin"}, Children: []Element{
			{Type: "test", Path: "note", Options: mapof.Template{"show-if": "level is high"}},
		}},
	}})

	test := func(role string, level string, expected string) {
		t.Helper()
		object := mapof.Any{"note": "stored"}
		values := url.Values{"role": {role}, "level": {level}, "note": {"posted"}}
		require.NoError(t, form.SetURLValues(&object, values, nil))
		require.Equal(t, expected, object.GetString("note"), "role %s, level %s", role, level)
	}

	test("user", "high", "stored")
	test("admin", "low", "stored")
	test("admin", "high", "posted")
}

// TestSetURLValues_SharedForm saves through one form in sequence and then from many goroutines,
// and requires that each object receives its own values while the form itself stays unchanged.
func TestSetURLValues_SharedForm(t *testing.T) {

	useTestWidget()

	newForm := func() Form {
		s := schema.New(schema.Object{Properties: schema.ElementMap{
			"role":   schema.String{},
			"name":   schema.String{},
			"admin":  schema.String{},
			"secret": schema.String{},
		}})

		return New(s, Element{Children: []Element{
			{Type: "test", Path: "role"},
			{Type: "test", Path: "name"},
			{Type: "test", Path: "admin", Options: mapof.Template{"show-if-option": "is-admin"}},
			{Options: mapof.Template{"show-if": "role is admin"}, Children: []Element{
				{Type: "test", Path: "secret"},
			}},
		}}, "is-admin:true")
	}

	form := newForm()

	// save returns a description of what went wrong, because require cannot run off the test goroutine
	save := func(index int) string {
		name := "user" + strconv.Itoa(index)
		role := "user"

		if index%2 == 0 {
			role = "admin"
		}

		object := mapof.Any{}
		values := url.Values{"role": {role}, "name": {name}, "admin": {name}, "secret": {name}}

		if err := form.SetURLValues(&object, values, nil); err != nil {
			return name + ": " + err.Error()
		}

		if object.GetString("name") != name {
			return name + ": wrong name " + object.GetString("name")
		}

		if object.GetString("admin") != name {
			return name + ": wrong admin " + object.GetString("admin")
		}

		if (object.GetString("secret") == name) != (role == "admin") {
			return name + ": secret written for role " + role
		}

		return ""
	}

	require.Empty(t, save(0))
	require.Empty(t, save(1))

	failures := make(chan string, 20)

	var wait sync.WaitGroup
	for index := range 20 {
		wait.Go(func() {
			if failure := save(index); failure != "" {
				failures <- failure
			}
		})
	}

	wait.Wait()
	close(failures)

	for failure := range failures {
		t.Error(failure)
	}

	require.Equal(t, newForm(), form, "saving must never write to the shared form")
}

// TestSetURLValues_URLValueSetter writes through a widget's own setter, and carries on past one that fails.
func TestSetURLValues_URLValueSetter(t *testing.T) {

	Use("test-setter", testSetterWidget{})

	s := schema.New(schema.Object{Properties: schema.ElementMap{
		"first":  schema.String{},
		"second": schema.String{},
	}})

	form := New(s, Element{Children: []Element{
		{Type: "test-setter", Path: "first"},
		{Type: "test-setter", Path: "second"},
	}})

	object := mapof.Any{"first": "stored", "second": "stored"}
	values := url.Values{"first": {"fail"}, "second": {"posted"}}
	require.NoError(t, form.SetURLValues(&object, values, nil), "a setter's error is logged, not returned")
	require.Equal(t, "stored", object.GetString("first"))
	require.Equal(t, "set by widget: posted", object.GetString("second"))
}

// TestSetURLValues_UnknownWidget fails the save when a field names a widget that is not registered.
func TestSetURLValues_UnknownWidget(t *testing.T) {

	s := schema.New(schema.Object{Properties: schema.ElementMap{
		"name": schema.String{},
	}})

	form := New(s, Element{Children: []Element{
		{Type: "no-such-widget", Path: "name"},
	}})

	object := mapof.Any{}
	require.Error(t, form.SetURLValues(&object, url.Values{"name": {"posted"}}, nil))
	require.Empty(t, object)
}

// testSetterWidget is a test widget that writes its value through its own URLValueSetter.
type testSetterWidget struct {
	testWidget
}

// SetURLValue writes the posted value with a prefix, and refuses the value "fail".
func (w testSetterWidget) SetURLValue(form *Form, element *Element, object any, values url.Values) error {

	value := values.Get(element.Path)

	if value == "fail" {
		return derp.Validation("Refused by the test widget", element.Path)
	}

	return form.Schema.Set(object, element.Path, "set by widget: "+value)
}
