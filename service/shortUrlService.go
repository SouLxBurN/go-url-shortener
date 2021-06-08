package service

import (
	"context"
	"errors"
	"fmt"
	"hash/crc32"
	"log"
	"time"
	"url-shortener/client"
	"url-shortener/config"
	"url-shortener/model"

	"go.mongodb.org/mongo-driver/bson"
)

// CreateShortUrl
// Creates a short url in the mongo database
// and returns the object Id
func CreateShortUrl(sURL *model.ShortURL) (interface{}, error) {
	mClient, err := client.GetMongoClient()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB")
	}
	mongoConfig := config.GetMongoConfig()
	shortUrls := mClient.Database(mongoConfig.Database).Collection(mongoConfig.Collection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	urlHash := createUrlHash(sURL.URL)

	url, err := GetShortUrl(urlHash)
	if url != "" && err == nil {
		return nil, errors.New("Url Already Exists")
	}

	_, err = shortUrls.InsertOne(ctx,
		bson.D{
			{Key: "hash", Value: urlHash},
			{Key: "url", Value: sURL.URL},
		})
	if err != nil {
		return "", err
	}
	return urlHash, nil
}

// createUrlHash Hashes the given string URL and returns a hexidecimal
// representation of the hash
func createUrlHash(URL string) string {
	algo := crc32.NewIEEE()
	algo.Write([]byte(URL))
	return fmt.Sprintf("%x", algo.Sum32())
}

// GetShortUrl
// Return the original url associated with the hash
// for redirect purposes.
func GetShortUrl(hash string) (string, error) {
	mClient, err := client.GetMongoClient()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB")
	}

	mongoConfig := config.GetMongoConfig()
	shortUrls := mClient.Database(mongoConfig.Database).Collection(mongoConfig.Collection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result model.ShortURL
	err = shortUrls.FindOne(ctx, bson.D{{Key: "hash", Value: hash}}).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.URL, nil
}
