package queries

import (
	"context"
	"fmt"
	"jonahjones777/product-compare/db"
	types "jonahjones777/product-compare/types"
	"os"
	"reflect"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
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

//GetProducts is the MongoDB call to retrieve items
var GetProducts = &graphql.Field{
	Type:        graphql.NewList(types.Product),
	Description: "Get all not todos",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		ctx := context.Background()
		productcollection := db.GetMongo(mongoCollection)
		product := productStruct{}
		todos, err := productcollection.Find(ctx, bson.D{})

		// Find() method raised an error
		if err != nil {
			fmt.Println("Finding all documents ERROR:", err)
			defer todos.Close(ctx)
		} else {
			for todos.Next(ctx) {
				// declare a result BSON object
				var result bson.M
				err := todos.Decode(&product)
				// If there is a cursor.Decode error
				if err != nil {
					fmt.Println("cursor.Next() error:", err)
					os.Exit(1)
					// If there are no cursor.Decode errors
				} else {
					fmt.Println("\nresult type:", reflect.TypeOf(result))
					fmt.Println("result:", result)
				}
			}
		}
		if err != nil {
			panic(err)
		}

		//Return a list of all elements
		var productList []productStruct
		productList = append(productList, product)
		return productList, nil
	},
}
