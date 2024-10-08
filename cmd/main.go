package main

import (
	"go-hypefast/cmd/api"

	_ "github.com/lib/pq"
)

func main() {
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		panic(err)
	}
}
