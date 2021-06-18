package client

import (
	"context"
	"os"
	"testing"
	"url-shortener/testingutils"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	mongoContainer := testingutils.SetupMongoTestContainer()
	retCode := m.Run()
	mongoContainer.Terminate(context.Background())
	os.Exit(retCode)
}

// TestMongoClient
func TestGetMongoClient(t *testing.T) {
	mongoClient, err := GetMongoClient()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, mongoClient)
}
