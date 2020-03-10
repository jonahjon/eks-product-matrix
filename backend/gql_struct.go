package main

import (
	"errors"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/kinds"
)

// _validate for ID support
func _validate(value string) error {
	if len(value) < 3 {
		return errors.New("The minimum length required is 3")
	}
	return nil
}

// ID support
var ID = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "ID",
	Description: "The `id` scalar type represents a ID Object.",
	Serialize: func(value interface{}) interface{} {
		return value
	},
	ParseValue: func(value interface{}) interface{} {
		var err error
		switch value.(type) {
		case string:
			err = _validate(value.(string))
		default:
			err = errors.New("Must be of type string.")
		}
		if err != nil {
			log.Fatal(err)
		}
		return value
	},
	ParseLiteral: func(valueAst ast.Value) interface{} {
		if valueAst.GetKind() == kinds.StringValue {
			err := _validate(valueAst.GetValue().(string))
			if err != nil {
				log.Fatal(err)
			}
			return valueAst
		} else {
			log.Fatal("Must be of type string.")
			return nil
		}
	},
})

//FeatureType is for the list of features
var FeatureType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Features",
	Fields: graphql.Fields{
		"feature": &graphql.Field{
			Type: graphql.String,
		},
	},
},
)

//ProductType
var ProductType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Product",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: ID,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"image": &graphql.Field{
			Type: graphql.String,
		},
		"saas_type": &graphql.Field{
			Type: graphql.String,
		},
		"price": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"features": &graphql.Field{
			Type: graphql.NewList(FeatureType),
		},
	},
})
