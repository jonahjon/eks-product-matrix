package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ProductModel struct
type ProductModel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name        string             `bson:"name" json:"name,omitempty"`
	Image       string             `bson:"image,omitempty" json:"image,omitempty"`
	SaasType    string             `bson:"saas_type,omitempty" json:"saas_type,omitempty"`
	Price       string             `bson:"price,omitempty" json:"price,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
}
