package route

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"url-shortener/model"
	"url-shortener/testingutils"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	mongoContainer := testingutils.SetupMongoTestContainer()
	retCode := m.Run()
	mongoContainer.Terminate(context.Background())
	os.Exit(retCode)
}

// TestUrlCreation
func TestUrlCreation(t *testing.T) {
	app := fiber.New()
	Configure(app)

	sURL := new(model.ShortURL)
	sURL.URL = "http://www.twitch.tv/soulxburn"

	bodyBytes, _ := json.Marshal(sURL)
	reader := bytes.NewReader(bodyBytes)

	req := httptest.NewRequest("POST", "/new", reader)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, resp.StatusCode, 200)
	byteBody := make([]byte, resp.ContentLength)
	resp.Body.Read(byteBody)

	var objmap map[string]string
	json.Unmarshal(byteBody, &objmap)

	shortUrl := objmap["shortUrl"]
	assert.NotEmpty(t, shortUrl)

	splits := strings.Split(shortUrl, "/")
	resp, err = app.Test(httptest.NewRequest("GET", "/"+splits[len(splits)-1], nil))
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 302, resp.StatusCode)
}

// TestResolveShortUrl
func TestResolveShortUrl(t *testing.T) {
	app := fiber.New()
	Configure(app)

	resp, err := app.Test(httptest.NewRequest("GET", "/abcd1234", nil))
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, resp.StatusCode, 404)
	assert.Equal(t, resp.Status, "404 Not Found")
}
