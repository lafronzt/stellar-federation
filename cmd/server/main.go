package main

import "github.com/lafronzt/stellar-federation/internal"

// main is the entrypoint for the application.
func main() {

	// Create and starts federation server.
	internal.StartServer()

}
