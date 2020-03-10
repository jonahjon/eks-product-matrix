package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Queries

//getAllProducts
func (db mongoDB) getAllProducts() (interface{}, error) {
	var results []ProductModel
	var err error

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := db.eks.Find(ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var elem ProductModel
		err = cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem)
		fmt.Printf("%v\n", elem.Features)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(ctx)
	return results, nil
}

//getProduct
func (db mongoDB) getProduct(name string) (interface{}, error) {
	var results ProductModel
	var err error
	q := bson.M{"name": name}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = db.eks.FindOne(ctx, q).Decode(&results)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%v\n", results)
	return results, nil
}

// addProduct
func (db mongoDB) addProduct(name string, image string, saastype string, price string, description string) (interface{}, error) {
	var err error
	var results ProductModel
	results.ID = primitive.NewObjectID()
	results.Name = name
	results.Image = image
	results.SaasType = saastype
	results.Price = price
	results.Description = description
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	_, err = db.eks.InsertOne(ctx, results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// deleteProduct
func (db mongoDB) deleteProduct(name string) (interface{}, error) {
	var err error
	var results ProductModel
	q := bson.M{"name": name}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = db.eks.FindOneAndDelete(ctx, q).Decode(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// updateProduct
func (db mongoDB) updateProduct(name string, image string, saastype string, price string, description string) (interface{}, error) {
	var err error
	var results ProductModel
	q := bson.M{"name": name}
	q2 := bson.M{"$set": bson.M{
		"name":        name,
		"image":       image,
		"saas_type":   saastype,
		"price":       price,
		"description": description,
	}}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = db.eks.FindOneAndUpdate(ctx, q, q2).Decode(&results)
	if err != nil {
		return nil, err
	}
	results.Name = name
	results.Image = image
	results.SaasType = saastype
	results.Price = price
	results.Description = description
	return results, nil
}
