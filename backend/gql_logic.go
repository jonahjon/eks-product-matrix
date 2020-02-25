package main

import (
	"github.com/graphql-go/graphql"
)

// Queries

// getAllRecipes
func getAllProducts(_ graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}
	results, err = db.getAllProducts()
	if err != nil {
		return nil, err
	}
	return results, nil
}

// getRecipe
func getProduct(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}
	name := p.Args["name"].(string)
	results, err = db.getProduct(name)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// addProduct
func addProduct(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}
	name := p.Args["name"].(string)
	image := p.Args["image"].(string)
	saastype := p.Args["saas_type"].(string)
	price := p.Args["price"].(string)
	description := p.Args["description"].(string)
	results, err = db.addProduct(name, image, saastype, price, description)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// deleteProduct
func deleteProduct(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}
	name := p.Args["name"].(string)
	results, err = db.deleteProduct(name)
	if err != nil {
		return nil, err
	}
	return results, nil
}

//updateProduct
func updateProduct(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}
	name := p.Args["name"].(string)
	image := p.Args["image"].(string)
	saastype := p.Args["saas_type"].(string)
	price := p.Args["price"].(string)
	description := p.Args["description"].(string)
	results, err = db.updateProduct(name, image, saastype, price, description)
	if err != nil {
		return nil, err
	}
	return results, nil
}
