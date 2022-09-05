package form

import (
	"github.com/benpate/rosetta/null"
	"github.com/benpate/rosetta/schema"
)

func getTestSchema() schema.Schema {

	return schema.Schema{
		ID:      "",
		Comment: "",
		Element: schema.Object{
			Properties: map[string]schema.Element{
				"username": schema.String{
					MinLength: null.NewInt(10),
					MaxLength: null.NewInt(100),
					Pattern:   "[a-z]+",
					Required:  true,
				},
				"name": schema.String{
					MaxLength: null.NewInt(50),
				},
				"email": schema.String{
					Format:    "email",
					MinLength: null.NewInt(10),
					MaxLength: null.NewInt(100),
					Required:  true,
				},
				"age": schema.Integer{
					Minimum:  null.NewInt64(10),
					Maximum:  null.NewInt64(100),
					Required: true,
				},
				"human": schema.Boolean{},
				"distance": schema.Number{
					Minimum:  null.NewFloat(10),
					Maximum:  null.NewFloat(100),
					Required: true,
				},
				"color": schema.String{
					Enum: []string{"Yellow", "Orange", "Red", "Violet", "Blue", "Green"},
				},
				"tags": schema.Array{
					Items: schema.String{
						Enum: []string{"pretty", "please", "my", "dear", "aunt", "sally"},
					},
				},
				"terms": schema.Boolean{},
				"ology": schema.Object{
					Properties: map[string]schema.Element{
						"biology":    schema.String{MaxLength: null.NewInt(1000)},
						"geology":    schema.String{MaxLength: null.NewInt(1000)},
						"psychology": schema.String{MaxLength: null.NewInt(1000)},
						"ontology":   schema.String{MaxLength: null.NewInt(1000)},
					},
				},
			},
		},
	}
}
