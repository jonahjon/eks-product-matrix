package main

import (
	"log"

	"github.com/graphql-go/graphql"
)

// Init the schema of GraphQL
func initSchema() graphql.Schema {
	graphqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"getAllProducts": &graphql.Field{
					Type:    graphql.NewList(ProductType),
					Args:    graphql.FieldConfigArgument{},
					Resolve: getAllProducts,
				},
				"getProduct": &graphql.Field{
					Type: ProductType,
					Args: graphql.FieldConfigArgument{
						"name": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						}},
					Resolve: getProduct,
				},
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"addProduct": &graphql.Field{
					Type: ProductType,
					Args: graphql.FieldConfigArgument{
						"name": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"image": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"saas_type": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"price": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"description": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"features": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.NewList(FeatureType)),
						}},
					Resolve: addProduct,
				},
				"deleteProduct": &graphql.Field{
					Type: ProductType,
					Args: graphql.FieldConfigArgument{
						"name": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						}},
					Resolve: deleteProduct,
				},
				"updateProduct": &graphql.Field{
					Type: ProductType,
					Args: graphql.FieldConfigArgument{
						"name": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"image": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"saas_type": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"price": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"description": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"features": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.NewList(FeatureType)),
						},
					},
					Resolve: updateProduct,
				},
			},
		}),
		Types: []graphql.Type{ID},
	})
	if err != nil {
		log.Fatal(err)
	}
	return graphqlSchema

}
