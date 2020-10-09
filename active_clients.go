package dnas

import (
	"context"
	"fmt"
	"net/http"
)

// ClientParameters represent
type ClientParameters struct {
	// ApMacAddress The mac address of the Access Point (AP).  Available for associated clients only.
	ApMacAddress *string `url:"apMacAddress,omitempty"`

	// Associated true|false.  Whether or not a device has connected to a network.
	Associated *bool `url:"associated,omitempty"`

	// BuildingID Unique identifier for a building from the map import process
	BuildingID *string `url:"buildingID,omitempty"`

	// CampusID Unique identifier for a campus from the map import process
	CampusID *string `url:"campusID,omitempty"`

	// DeviceID The device unique identifier, for example the device macAddress.
	DeviceID *string `url:"deviceID,omitempty"`

	// DeviceType CLIENT, TAG, ROGUE_AP, ROGUE_CLIENT or INTERFERER
	DeviceType *string `url:"deviceType,omitempty"`

	// FloorID Unique identifier for a floor from the map import process
	FloorID *string `url:"floorID,omitempty"`

	// Format Indicate if using geojson, value is "geojson" if so.
	Format *string `url:"format,omitempty"`

	// IPAddress IP address of the connected device.  Available for associated clients only.
	IPAddress *string `url:"iPAddress,omitempty"`

	// Limit The maximum number of items that may be returned for a single request. For active client, the default value is 1000; For client location history, the default value is 2000.
	Limit *string `url:"limit,omitempty"`

	// Manufacturer Manufacturer of the device.
	Manufacturer *string `url:"manufacturer,omitempty"`

	// MapElementID Indicate the map element unique identifier.
	MapElementID *string `url:"mapElementID,omitempty"`

	// MapElementLevel Indicate the map element level, valid value is "campus", "building" and "floor".
	MapElementLevel *string `url:"mapElementLevel,omitempty"`

	// Page The page number requests for. Start from 1 and default value is 1.
	Page *string `url:"page,omitempty"`

	// RogueApClients When using deviceType=ROGUE_AP, this will return rogue APs that have connected clients.
	RogueApClients *bool `url:"rogueApClients,omitempty"`

	// Ssid Wifi service set identifier (SSID).  Available for associated clients only.
	Ssid *string `url:"ssid,omitempty"`

	// Username The user name of the connected user. Available for associated clients only.
	Username *string `url:"username,omitempty"`
}

// ClientCountResponse provides the count for the active devices from Count()
type ClientCountResponse struct {
	Results struct {
		Total int `json:"total"`
	} `json:"results"`
	QueryString struct {
		TenantID string `json:"tenantId"`
	} `json:"querystring"`
	Success bool `json:"success"`
}

// ClientFloorsResponse holds the results from Floors()
type ClientFloorsResponse struct {
	Results []struct {
		// The total number of associated devices on this floor
		Count int64 `json:"count,omitempty"`

		// The floors unique identifier
		FloorID string `json:"floorId,omitempty"`
	} `json:"results"`

	Success bool `json:"success,omitempty"`
}

// Clients returns active clients.  If no parameters are given, all active clients are returned with pagination. The default page number is 1, default number of items per page is 1000.
func (s *ActiveClientsService) Clients(ctx context.Context, opt *ClientParameters) error {
	return nil
}

// Count retrieves the active clients count. The API supports searching by a variety of parameters. If no parameters are given, the count of all active clients are returned.
func (s *ActiveClientsService) Count(ctx context.Context, opts *ClientParameters) (ClientCountResponse, error) {
	ccr := ClientCountResponse{}
	url := fmt.Sprintf("%s/clients/count", s.client.BaseURL)
	u, err := addOptions(url, opts)
	if err != nil {
		return ccr, err
	}
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return ccr, err
	}
	if err := s.client.makeRequest(ctx, req, &ccr); err != nil {
		return ccr, err
	}
	return ccr, nil
}

// Floors provides a list of all the floors unique identifiers which have associated clients.
func (s *ActiveClientsService) Floors(ctx context.Context) (ClientFloorsResponse, error) {
	cfr := ClientFloorsResponse{}
	url := fmt.Sprintf("%s/clients/floors", s.client.BaseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return cfr, err
	}
	if err := s.client.makeRequest(ctx, req, &cfr); err != nil {
		return cfr, err
	}
	return cfr, nil
}
