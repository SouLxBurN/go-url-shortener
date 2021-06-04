# URL Shortener | Golang

This is a learning project for exploring Go, and GoFiber.

## Prequisites
- MongoDB
	- You will need a instance of mongo, and you'll need to update the `CONNECTION_URL` in `client/mongo.go` to point to your instance.

## Running
The project isn't equipt with any build tooling as of yet. In order to build and run this project you will be using standard go commands.

`go build .`
`go run .`

## Try it out
Once you have the service running you can run the follow curl to create a short url.

```
curl -X POST -H "Content-Type: application/json" --data "{\"url\":\"https://monkeytype.com\"}" localhost:3000/new
```

The hash it returns can be used to hit `localhost:3000/{hash}` to be redirected to the saved url.
