package handlers

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

type Podcast struct {
	Name   string `bson:"name,omitempty"`
	Rating int32  `bson:"rating,omitempty"`
}

const uri = "mongodb://localhost:27017"

func NewPodcast() *Podcast {
	return &Podcast{}
}

func (p *Podcast) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case "/podcast":
		switch r.Method {
		case "GET":
			result := getPodcast()
			json.NewEncoder(rw).Encode(result)
		case "POST":
			createPodcast()
		}
	}

}

func getPodcast() []Podcast {
	var slice []Podcast

	opts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	podcastsCollection := client.Database("quickstart").Collection("podcasts")

	cursor, err := podcastsCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	err = cursor.All(context.TODO(), &slice)

	if err != nil {
		panic(err)
	}

	return slice
}

func createPodcast() {
	opts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), opts)

	quickstartDatabase := client.Database("quickstart")
	podcastsCollection := quickstartDatabase.Collection("podcasts")

	if err != nil {
		panic(err)
	}

	podcast := Podcast{Name: "Best", Rating: 4}
	_, err = podcastsCollection.InsertOne(context.TODO(), podcast)
}
