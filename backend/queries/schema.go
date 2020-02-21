package queries

import (
	fields "jonahjones777/product-compare/queries/fields"

	"github.com/graphql-go/graphql"
)

//RootQuery is the function to get all the fields from base queries
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"products": fields.GetProducts,
	},
})
