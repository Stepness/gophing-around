package main

import (
	"api/handlers"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

var uri = "mongodb://localhost:27017"

func main() {
	fmt.Println("Best API started!")

	l := log.New(os.Stdout, "api", log.LstdFlags)
	opts := options.Client().ApplyURI(uri)
	ctx := context.Background()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	//ph := handlers.NewProducts(l)
	pch := handlers.NewPodcast(client)

	//mux.Handle("GET /", ph)
	mux.Handle("/podcast", pch)

	s := http.Server{
		Addr:     ":8080",
		Handler:  mux,
		ErrorLog: l,
	}

	_ = s.ListenAndServe()
}
