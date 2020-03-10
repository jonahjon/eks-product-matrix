package main

import "fmt"

type configModel struct {
	mongoUri    string
	mongoDb     string
	tokenSecret string
	tokenExp    string
	serveUri    string
}

var config = configModel{
	mongoUri:    fmt.Sprintf("mongodb://%v:27017/eks", "localhost"), // mongodb://mongodb:27017/eks
	mongoDb:     "eks",                                              // DB name
	tokenSecret: "secret",                                           // Secret to use in Tokens
	tokenExp:    "1h",                                               // Expiration of Token
	serveUri:    ":8080",                                            // Serve
}
