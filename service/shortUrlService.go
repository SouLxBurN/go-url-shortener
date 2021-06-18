package service

import (
	"context"
	"fmt"
	"hash/crc32"
	"log"
	"time"
	"url-shortener/client"
	"url-shortener/config"
	"url-shortener/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateShortUrl
// Creates a short url in the mongo database
// and returns the object Id
func CreateShortUrl(sURL *model.ShortURL) error {
	mClient, err := client.GetMongoClient()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB")
	}
	mongoConfig := config.GetMongoConfig()
	shortUrls := mClient.Database(mongoConfig.Database).Collection(mongoConfig.Collection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sURL.Hash = createUrlHash(sURL.URL)

	// Check if hash already exists
	shortUrl, err := GetShortUrl(sURL.Hash)
	if shortUrl != nil && err == nil {
		return nil
	}

	_, err = shortUrls.InsertOne(ctx, sURL)
	if err != nil {
		return err
	}
	return nil
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
func GetShortUrl(hash string) (*model.ShortURL, error) {
	mClient, err := client.GetMongoClient()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB")
	}

	mongoConfig := config.GetMongoConfig()
	shortUrls := mClient.Database(mongoConfig.Database).Collection(mongoConfig.Collection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := new(model.ShortURL)
	filter := bson.D{{Key: "hash", Value: hash}}
	err = shortUrls.FindOne(ctx, filter).Decode(result)
	if err != nil {
		// If not found, return no result and no error
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}
