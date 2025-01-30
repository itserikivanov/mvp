package main

import "os"

// @title Swagger MVP API
// @version 1.0
// @description Go server for MVP app
// @BasePath /api
func main() {
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8000"
	}

	startServer(httpPort)
}
