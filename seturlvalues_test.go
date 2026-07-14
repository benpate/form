package form

import (
	"net/url"
	"testing"

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
			{Type: "test", Path: "comment", Options: mapof.Any{"show-if": "type is detailed"}},
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

	element := Element{Options: mapof.Any{"provider": "things"}}

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
	element := Element{Options: mapof.Any{"provider": "things"}}
	provider := mockProvider{group: &mockWritableGroup{}}

	// A value without the new-item prefix is passed through unchanged
	value, updated, err := element.replaceNewLookup(provider, "existing-value")
	require.NoError(t, err)
	require.False(t, updated)
	require.Equal(t, "existing-value", value)
}

func TestReplaceNewLookup_NoProviderOption(t *testing.T) {
	element := Element{Options: mapof.Any{}}
	provider := mockProvider{group: &mockWritableGroup{}}

	value, updated, err := element.replaceNewLookup(provider, NewItemIdentifier+"x")
	require.NoError(t, err)
	require.False(t, updated)
	require.Equal(t, NewItemIdentifier+"x", value)
}

func TestReplaceNewLookup_ReadOnlyGroup(t *testing.T) {

	// A non-writable group cannot accept new values
	element := Element{Options: mapof.Any{"provider": "things"}}
	provider := mockProvider{group: NewReadOnlyLookupGroup()}

	value, updated, err := element.replaceNewLookup(provider, NewItemIdentifier+"x")
	require.NoError(t, err)
	require.False(t, updated)
	require.Equal(t, NewItemIdentifier+"x", value)
}
