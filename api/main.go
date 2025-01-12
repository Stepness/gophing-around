package main

import (
	"api/handlers"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:27017"

func main() {
	fmt.Println("Best API started!")

	l := log.New(os.Stdout, "api", log.LstdFlags)

	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), opts)

	quickstartDatabase := client.Database("quickstart")
	podcastsCollection := quickstartDatabase.Collection("podcasts")

	if err != nil {
		panic(err)
	}

	_, err = podcastsCollection.InsertOne(context.TODO(), bson.D{
		{Key: "title", Value: "The Polyglot Developer Podcast"},
		{Key: "author", Value: "Nic Raboy"},
	})

	mux := http.NewServeMux()
	//ph := handlers.NewProducts(l)
	pch := handlers.NewPodcast()

	//mux.Handle("GET /", ph)
	mux.Handle("/podcast", pch)

	s := http.Server{
		Addr:     ":8080",
		Handler:  mux,
		ErrorLog: l,
	}

	_ = s.ListenAndServe()
}
