package store

import "go.mongodb.org/mongo-driver/bson/primitive"

type Url struct {
	Id       primitive.ObjectID `bson:"_id"`
	LongUrl  string             `bson:"longUrl"`
	ShortUrl string             `bson:"shortUrl"`
}

type UrlStore interface {
	StoreUrl(longUrl string, shortUrl string) error
	GetUrl(shortUrl string) (*Url, error)
}
