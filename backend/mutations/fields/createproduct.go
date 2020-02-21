package mutations

import (
	"context"
	"jonahjones777/product-compare/db"
	types "jonahjones777/product-compare/types"

	"github.com/graphql-go/graphql"
)

type productStruct struct {
	ID          string `json:"id"`
	NAME        string `json:"name"`
	IMAGE       string `json:"image"`
	SAASTYPE    string `json:"saas_type"`
	PRICE       string `json:"price"`
	DESCRIPTION string `json:"description"`
}

const mongoCollection string = "eks"

//CreateProducts is the MongoDB call to add a new item
var CreateProducts = &graphql.Field{
	Type:        types.Product,
	Description: "Create a not Todo",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"image": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"saas_type": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"price": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id, _ := params.Args["id"].(string)
		name, _ := params.Args["name"].(string)
		image, _ := params.Args["image"].(string)
		saastype, _ := params.Args["saas_type"].(string)
		price, _ := params.Args["price"].(string)
		description, _ := params.Args["description"].(string)
		ctx := context.Background()
		productcollection := db.GetMongo(mongoCollection)
		_, err := productcollection.InsertOne(ctx, map[string]string{
			"id":          id,
			"name":        name,
			"image":       image,
			"saas_type":   saastype,
			"price":       price,
			"description": description,
		})
		if err != nil {
			panic(err)
		}
		return productStruct{id, name, image, saastype, price, description}, nil
	},
}
