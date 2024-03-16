package main

import (
	"afauble/go/chessapp/internal/routes"
)

func main() {

	port := "8080"

	// Creates Url Paths
	routes.Gin_url_setup(port)
}
