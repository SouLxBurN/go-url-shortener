package service

import (
	"context"
	"fmt"
	"os"
	"testing"
	"url-shortener/model"
	"url-shortener/testingutils"
)

func TestMain(m *testing.M) {
	mongoContainer := testingutils.SetupMongoTestContainer()
	retCode := m.Run()
	mongoContainer.Terminate(context.Background())
	os.Exit(retCode)
}

func TestCreateUrlHash(t *testing.T) {
	hash := createUrlHash("hashme!")
	fmt.Println(hash)
}

func TestCreateShortUrl(t *testing.T) {
	sURL := model.ShortURL{
		Hash: "abc123",
		URL:  "https://www.twitch.tv/soulxburn",
	}

	err := CreateShortUrl(&sURL)
	if err != nil {
		t.Error(err)
	}

	fetchedURL, err := GetShortUrl(sURL.Hash)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(fetchedURL)

	// _, err = shortUrls.InsertOne(ctx, sURL)
	// if err != nil {
	// 	t.Error(err)
	// }

	// result := new(model.ShortURL)
	// filter := bson.D{{Key: "hash", Value: sURL.Hash}, {Key: "url", Value: sURL.URL}}

	// err = shortUrls.FindOne(ctx, filter).Decode(&result)
	// if err != nil {
	// 	t.Error(err)
	// }
}
