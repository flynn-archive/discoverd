package main

import (
	"fmt"
	"github.com/progrium/go-discover/discover"
)

func main() {
	server := discover.NewServer()
	fmt.Printf("Starting server on %s...\n", server.Address)
	discover.ListenAndServe(server)
}
