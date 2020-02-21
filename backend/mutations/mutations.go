package mutations

import (
	fields "jonahjones777/product-compare/mutations/fields"

	"github.com/graphql-go/graphql"
)

//RootMutation is the function to get add to the fields from base queries
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createNotTodo": fields.CreateProducts,
	},
})
