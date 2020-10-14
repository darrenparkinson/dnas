// Example: count/main.go
//
// This example shows how to initialise the library and retrieve a count of devices.
//
// In order to use this library, you will  need two environment variables:
//   - DNAS_REGION - if you are in the USA, set to "io", elsewhere, set to "eu".
//   - DNAS_API_KEY - set this to your API Key from DNA Spaces: https://developer.cisco.com/docs/dna-spaces/#!getting-started/getting-started
//
// How to run:
//
//    go run ./examples/count/main.go
//
package main

import (
	"context"
	"fmt"
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
	count, err := c.ActiveClientsService.GetCount(ctx, &dnas.ClientParameters{Associated: dnas.Bool(true), DeviceType: dnas.String("CLIENT")})
	fmt.Println("Count of Associated Clients:", count.Results.Total)
}
