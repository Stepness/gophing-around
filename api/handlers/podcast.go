package handlers

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type Podcast struct {
	client *mongo.Client
}

type PodcastEntity struct {
	Name   string `bson:"name,omitempty"`
	Rating int32  `bson:"rating,omitempty"`
}

func NewPodcast(client *mongo.Client) *Podcast {
	return &Podcast{client: client}
}

func (p *Podcast) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case "/podcast":
		switch r.Method {
		case "GET":
			result, err := p.getPodcast(r.Context())
			if err != nil {
				json.NewEncoder(rw).Encode(err)
			} else {
				json.NewEncoder(rw).Encode(result)
			}
		case "POST":
			var body PodcastEntity
			json.NewDecoder(r.Body).Decode(&body)

			err := p.createPodcast(body, r.Context())

			if err != nil {
				json.NewEncoder(rw).Encode(err)
			} else {
				json.NewEncoder(rw).Encode("Data added")
			}
		}
	}

}

func (p *Podcast) getPodcast(ctx context.Context) ([]PodcastEntity, error) {
	var slice []PodcastEntity

	podcastsCollection := p.client.Database("quickstart").Collection("podcasts")

	cursor, err := podcastsCollection.Find(ctx, bson.D{})
	defer cursor.Close(ctx)

	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &slice)

	return slice, err
}

func (p *Podcast) createPodcast(entity PodcastEntity, ctx context.Context) error {
	podcastsCollection := p.client.Database("quickstart").Collection("podcasts")

	podcast := PodcastEntity{Name: entity.Name, Rating: entity.Rating}
	_, err := podcastsCollection.InsertOne(ctx, podcast)

	return err
}
