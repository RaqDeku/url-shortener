package dbstore

import (
	"context"
	"url-shorter/server/internals/store"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UrlStore struct {
	collection mongo.Collection
}

func NewUrlStore() *UrlStore {
	return &UrlStore{
		collection: *GetDbCollection(),
	}
}

func (s *UrlStore) StoreUrl(longUrl string, shortUrl string) error {
	url := &store.Url{
		Id:       primitive.NewObjectID(),
		LongUrl:  longUrl,
		ShortUrl: shortUrl,
	}
	_, err := s.collection.InsertOne(context.TODO(), url)

	if err != nil {
		return err
	}

	return nil
}

func (s *UrlStore) GetUrl(shortUrl string) (*store.Url, error) {
	var url store.Url
	filter := bson.D{{Key: "shortUrl", Value: shortUrl}}
	err := s.collection.FindOne(context.TODO(), filter).Decode(&url)

	if err != nil {
		return nil, err
	}

	return &url, nil
}
