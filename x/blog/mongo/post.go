package mongo

import (
	"blog/x/blog/types"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

const mogo_uri string = "mongodb://127.0.0.1:27017"

type MongoPost struct {
	id			uint64 `bson:"creator,omitempty"`
	Creator     string `bson:"creator,omitempty"`
	Title       string `bson:"title,omitempty"`
	Body        string `bson:"body,omitempty"`
}

func AddPostMongo(post types.Post) error {
	client, err := GetMongoClient(mogo_uri)
	if err != nil {
		log.Fatal(err)
		return err
	}

	collection := client.Database("blog").Collection("posts")
	collection.InsertOne(context.Background(), bson.M{
		"id":			post.Id,
		"creator":     	post.Creator,
		"title":       	post.Title,
		"body":        	post.Body,
	})
	return nil
}