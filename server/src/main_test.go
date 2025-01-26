package main

import (
	"testing"
)

func TestServer(t *testing.T) {
	t.Setenv("PORT", "8001")
	go main()
}
