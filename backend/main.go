package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/graphql-go/handler"
)

var db mongoDB

func main() {
	fmt.Println("Golang+GraphQL+MongoDB Server v1.0.0")

	// MongoDB
	db = connectDB()
	defer db.closeDB()

	//Graphql
	schema := initSchema()
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// Serve
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./public")))
	mux.Handle("/graphql", handlers.LoggingHandler(os.Stdout, disableCors(h)))
	log.Println("Now server is running on port 8080")
	s := &http.Server{
		Addr:    config.serveUri,
		Handler: mux,
	}
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
