package main

import (
	"testing"
)

func TestServer(t *testing.T) {
	go startServer("8001")
}
