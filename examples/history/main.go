// Example: history/main.go
//
// This example shows how to initialise the library and retrieve a history of devices over the last two hours.
//
// In order to use this library, you will  need two environment variables:
//   - DNAS_REGION - if you are in the USA, set to "io", elsewhere, set to "eu".
//   - DNAS_API_KEY - set this to your API Key from DNA Spaces: https://developer.cisco.com/docs/dna-spaces/#!getting-started/getting-started
//
// How to run:
//
//    go run ./examples/history/main.go
//
package main

import (
	"context"
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/darrenparkinson/dnas"
)

func main() {
	c, err := dnas.NewClient(os.Getenv("DNAS_API_KEY"), os.Getenv("DNAS_REGION"), nil)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	fromTime := time.Now().Add(time.Hour*-2).UnixNano() / int64(time.Millisecond)
	toTime := time.Now().UnixNano() / int64(time.Millisecond)
	history, err := c.HistoryService.GetHistory(ctx,
		&dnas.HistoryParameters{
			StartTime: dnas.String(strconv.FormatInt(fromTime, 10)),
			EndTime:   dnas.String(strconv.FormatInt(toTime, 10)),
		})
	// Demonstrate use of errors.Is
	if errors.Is(err, dnas.ErrInternalError) {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Found %d results.\n", len(history.Results))
	log.Printf("First found: %+v\n", history.Results[0])
}
