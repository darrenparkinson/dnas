// Example: clients/main.go
//
// This example shows how to initialise the library and retrieve a list of devices.
//
// In order to use this library, you will  need two environment variables:
//   - DNAS_REGION - if you are in the USA, set to "io", elsewhere, set to "eu".
//   - DNAS_API_KEY - set this to your API Key from DNA Spaces: https://developer.cisco.com/docs/dna-spaces/#!getting-started/getting-started
//
// How to run:
//
//    go run ./examples/clients/main.go
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
	ac, err := c.ActiveClientsService.ListClients(ctx, &dnas.ClientParameters{Associated: dnas.Bool(true), DeviceType: dnas.String("CLIENT"), Limit: dnas.String("10"), Page: dnas.String("1")})
	if err != nil {
		log.Fatal(err)
	}
	for _, client := range ac.Results {
		log.Printf("Hierarchy: %s\tIPAddress: %s\n", client.Hierarchy, client.IPAddress)
	}
}
