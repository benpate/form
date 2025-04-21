package form

import (
	"net/url"
	"testing"

	"github.com/benpate/rosetta/mapof"
	"github.com/benpate/rosetta/schema"
	"github.com/stretchr/testify/require"
)

func TestFormSetURLValues(t *testing.T) {

	useTestWidget()

	form := New(
		schema.New(schema.Object{
			Properties: schema.ElementMap{
				"name":       schema.String{},
				"email":      schema.String{},
				"age":        schema.Integer{RequiredIf: "requireAge is true"},
				"requireAge": schema.Boolean{},
				"showEmail":  schema.Boolean{},
			},
		}),
		Element{
			Type: "test",
			Children: []Element{
				{Type: "test", Path: "name"},
				{Type: "test", Path: "age"},
				{Type: "test", Path: "email", Options: mapof.Any{"show-if": "showEmail is true"}},
				{Type: "test", Path: "requireAge"},
				{Type: "test", Path: "showEmail"},
			},
		},
	)

	{
		// First Test Email IS SET because showEmail is true
		data := url.Values{
			"name":       []string{"John Connor"},
			"email":      []string{"john@connor.mil"},
			"age":        []string{"42"},
			"requireAge": []string{"false"},
			"showEmail":  []string{"true"},
		}

		target := mapof.Any{}
		err := form.SetURLValues(&target, data, nil)
		require.Nil(t, err)
		require.Equal(t, "John Connor", target["name"])
		require.Equal(t, "john@connor.mil", target["email"])
	}

	{
		// Second Test: Email IS NOT SET because showEmail is false
		data := url.Values{
			"name":      []string{"John Connor"},
			"email":     []string{"john@connor.mil"},
			"showEmail": []string{"false"},
		}

		target := mapof.Any{}
		err := form.SetURLValues(&target, data, nil)
		require.Equal(t, "John Connor", target["name"])
		require.Nil(t, target["email"])
		require.Nil(t, err)
	}

	{
		// Third: Email IS NOT SET because showEmail is missing
		data := url.Values{
			"name":  []string{"John Connor"},
			"email": []string{"john@connor.mil"},
		}

		target := mapof.Any{}
		err := form.SetURLValues(&target, data, nil)
		require.Equal(t, "John Connor", target["name"])
		require.Nil(t, target["email"])
		require.Nil(t, err)
	}

	{
		// Fourth Test: Age IS NOT required, because requireAge is false
		data := url.Values{
			"name":       []string{"John Connor"},
			"requireAge": []string{"false"},
		}

		target := mapof.Any{}
		err := form.SetURLValues(&target, data, nil)
		require.Equal(t, "John Connor", target["name"])
		require.Equal(t, false, target["requireAge"])
		require.Nil(t, err)
	}

	{
		// Fifth Test: Age IS SET becasue it is present
		data := url.Values{
			"name":       []string{"John Connor"},
			"age":        []string{"42"},
			"requireAge": []string{"false"},
		}

		target := mapof.Any{}
		err := form.SetURLValues(&target, data, nil)
		require.Equal(t, "John Connor", target["name"])
		require.Equal(t, 42, target["age"])
		require.Equal(t, false, target["requireAge"])
		require.Nil(t, err)
	}

	{
		// Fifth Test: Form doesn't validate because age is conditionally required.
		data := url.Values{
			"name":       []string{"John Connor"},
			"requireAge": []string{"true"},
		}

		target := mapof.Any{}
		err := form.SetURLValues(&target, data, nil)
		require.Error(t, err)
	}

}
