package testingutils

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"url-shortener/config"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// setupMongoTestContainer Creates a mongo test container
// for testing integrations with mongo.
func SetupMongoTestContainer() testcontainers.Container {
	_, filename, _, _ := runtime.Caller(0)
	fmt.Println("filename: " + filename)
	// The ".." may change depending on you folder structure
	dir := path.Join(path.Dir(filename), "..")
	fmt.Println("directory: " + dir)
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "mongo",
		ExposedPorts: []string{"27017"},
		WaitingFor:   wait.ForListeningPort("27017"),
	}

	mCont, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatal(err)
	}

	ip, err := mCont.Host(ctx)
	if err != nil {
		log.Fatal(err)
	}

	port, err := mCont.MappedPort(ctx, "27017")
	if err != nil {
		log.Fatal(err)
	}

	config.GetMongoConfig().ConnectionString = fmt.Sprintf("mongodb://%s:%s", ip, port.Port())

	return mCont
}
