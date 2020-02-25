package main

type configModel struct {
	mongoUri    string
	mongoDb     string
	tokenSecret string
	tokenExp    string
	serveUri    string
}

var config = configModel{
	mongoUri:    "mongodb://mongodb:27017/eks", // Mongo Uri example: mongodb://admin:123456@localhost:37812/react-recipes
	mongoDb:     "eks",                         // DB name
	tokenSecret: "secret",                      // Secret to use in Tokens
	tokenExp:    "1h",                          // Expiration of Token
	serveUri:    ":8080",                       // Serve
}
