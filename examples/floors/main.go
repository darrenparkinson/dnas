// Example: floors/main.go
//
// This example shows how to initialise the library and retrieve a list of floors.
//
// In order to use this library, you will  need two environment variables:
//   - DNAS_REGION - if you are in the USA, set to "io", elsewhere, set to "eu".
//   - DNAS_API_KEY - set this to your API Key from DNA Spaces: https://developer.cisco.com/docs/dna-spaces/#!getting-started/getting-started
//
// How to run:
//
//    go run ./examples/floors/main.go
//
package main

import (
	"context"
	"log"
	"os"

	"github.com/darrenparkinson/dnas"
)

func main() {
	c, err := dnas.NewClient(os.Getenv("DNAS_API_KEY"), os.Getenv("DNAS_REGION"), nil)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	floors, err := c.ActiveClientsService.ListFloors(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, floor := range floors.Results {
		log.Printf("FloorID: %s\tCount: %d\n", floor.FloorID, floor.Count)
	}
}
