package dnas

import (
	"context"
	"fmt"
	"net/http"
)

// AccessPointStatus represents the status passed to get the access point count for a given status
type AccessPointStatus string

// Fields for AccessPointStatus
const (
	All      AccessPointStatus = "all"
	Active   AccessPointStatus = "active"
	Inactive AccessPointStatus = "inactive"
	Missing  AccessPointStatus = "missing"
)

// AccessPointsResponse provides a list of missing access points returned by ListAccessPoints
type AccessPointsResponse []struct {
	// ap mac
	ApMac string `json:"apMac,omitempty"`

	// The message count received.
	Count int64 `json:"count,omitempty"`

	// The message rate in 15 min.
	M15Rate float64 `json:"m15Rate,omitempty"`

	// The message rate in 1 min.
	M1Rate float64 `json:"m1Rate,omitempty"`

	// The message rate in 5 min.
	M5Rate float64 `json:"m5Rate,omitempty"`
}

// AccessPointsCountResponse provides the count for the active, inactive, missing or all the access points from GetCount()
type AccessPointsCountResponse struct {
	// Count of access points for a given status
	Count int64 `json:"count"`
}

// ListAccessPoints retrieves a list of missing access points.
// The only valid status is "missing" and is therefore not provided as an option.
func (s *AccessPointsService) ListAccessPoints(ctx context.Context) (AccessPointsResponse, error) {
	apr := AccessPointsResponse{}
	url := fmt.Sprintf("%s/accessPoints?status=missing", s.client.BaseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return apr, err
	}
	if err := s.client.makeRequest(ctx, req, &apr); err != nil {
		return apr, err
	}
	return apr, nil
}

// GetCount retrieves the count of the active, inactive, missing or all the access points.
// If no parameters are given, the count of all active access points is returned.
// Status may be missing, active, inactive or all
func (s *AccessPointsService) GetCount(ctx context.Context, status AccessPointStatus) (AccessPointsCountResponse, error) {
	if status == "" {
		status = "missing"
	}
	apcr := AccessPointsCountResponse{}
	url := fmt.Sprintf("%s/accessPoints/count?status=%s", s.client.BaseURL, status)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return apcr, err
	}
	if err := s.client.makeRequest(ctx, req, &apcr); err != nil {
		return apcr, err
	}
	return apcr, nil
}
