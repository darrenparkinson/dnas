// Package dnas provides a library for the Cisco DNA Spaces API
package dnas

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/google/go-querystring/query"
)

// Client represents connectivity to the DNA Spaces API
type Client struct {
	// BaseURL for DNA Spaces API.  Set to v1 and the relevant region using `dnas.New()`, or set directly.
	BaseURL string

	//HTTP Client to use for making requests, allowing the user to supply their own if required.
	HTTPClient *http.Client

	//API Key for DNA Spaces.  See [the documentation on how to generate one](https://developer.cisco.com/docs/dna-spaces/#!getting-started).
	APIKey string

	AccessPointsService  *AccessPointsService
	ActiveClientsService *ActiveClientsService
	HistoryService       *HistoryService
	NotificationsService *NotificationsService
	MapService           *MapService
}

// AccessPointsService represents the Access Points API group
type AccessPointsService struct {
	client *Client
}

// ActiveClientsService represents the Active Clients API group
type ActiveClientsService struct {
	client *Client
}

// HistoryService represents the Clients History API group
type HistoryService struct {
	client *Client
}

// MapService represents the Map API group
type MapService struct {
	client *Client
}

// NotificationsService represents the Notifications API group
type NotificationsService struct {
	client *Client
}

// NewClient is a helper function that returns an new dnas client given a region (io or eu) and API Key.
// Optionally you can provide your own http client or use nil to use the default.
func NewClient(apikey string, region string, client *http.Client) (*Client, error) {
	if apikey == "" {
		return nil, errors.New("apikey required")
	}
	if region == "" || (region != "io" && region != "eu") {
		return nil, errors.New("valid region required, either io or eu")
	}
	if client == nil {
		client = &http.Client{
			Timeout: 10 * time.Second,
		}
	}
	c := &Client{
		BaseURL:    fmt.Sprintf("https://dnaspaces.%s/api/location/v1", region),
		HTTPClient: client,
		APIKey:     apikey,
	}
	c.AccessPointsService = &AccessPointsService{client: c}
	c.ActiveClientsService = &ActiveClientsService{client: c}
	c.HistoryService = &HistoryService{client: c}
	c.MapService = &MapService{client: c}
	c.NotificationsService = &NotificationsService{client: c}

	return c, nil
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

// makeRequest provides a single function to add common items to the request.
func (c *Client) makeRequest(ctx context.Context, req *http.Request, v interface{}) error {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))
	req.Header.Set("Accept", "application/json")
	req.WithContext(ctx)
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {

		var ciscoErr error

		switch res.StatusCode {
		case 400:
			ciscoErr = ErrBadRequest
		case 401:
			ciscoErr = ErrUnauthorized
		case 403:
			ciscoErr = ErrForbidden
		case 500:
			ciscoErr = ErrInternalError
		default:
			ciscoErr = ErrUnknown
		}

		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			ciscoErr = fmt.Errorf("%w: %s", ciscoErr, errRes.Message)
		}

		return ciscoErr

	}

	if res.StatusCode == http.StatusCreated {
		return nil
	}

	if res.Header.Get("Content-Type") == "text/csv" {
		reader := csv.NewReader(res.Body)
		records, err := reader.ReadAll()
		if err != nil {
			return err
		}
		if p, ok := v.(*[][]string); ok {
			*p = records
		} else {
			return errors.New("invalid type assertion: v interface{} should be *[][]string for csv records")
		}
		return nil
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}

// addOptions adds the parameters in opts as URL query parameters to s. opts
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}
