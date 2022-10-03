package main

import (
	router "main.go/internal/adapters/api"
	"main.go/internal/adapters/stream"
)

func main() {

	go stream.Consume()
	router.HandleRequest()

}
