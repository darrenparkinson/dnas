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

	// Limit The maximum number of items that may be returned for a single request. For active client, the default value is 1000; For client location history, the default value is 2000. Sorry it's a string, but that's what we got!
	Limit *string `url:"limit,omitempty"`

	// Manufacturer Manufacturer of the device.
	Manufacturer *string `url:"manufacturer,omitempty"`

	// MapElementID Indicate the map element unique identifier.
	MapElementID *string `url:"mapElementID,omitempty"`

	// MapElementLevel Indicate the map element level, valid value is "campus", "building" and "floor".
	MapElementLevel *string `url:"mapElementLevel,omitempty"`

	// Page The page number requests for. Start from 1 and default value is 1. Sorry it's a string, but that's what we got!
	Page *string `url:"page,omitempty"`

	// RogueApClients When using deviceType=ROGUE_AP, this will return rogue APs that have connected clients.
	RogueApClients *bool `url:"rogueApClients,omitempty"`

	// Ssid Wifi service set identifier (SSID).  Available for associated clients only.
	Ssid *string `url:"ssid,omitempty"`

	// Username The user name of the connected user. Available for associated clients only.
	Username *string `url:"username,omitempty"`
}

// LocationDeviceQuery represents the QueryString values used.
// It's the same as ClientParameters, but the types are different, typically string for everything.
// Empty strings are returned for values that are missing.
type LocationDeviceQuery struct {
	ApMacAddress    string `url:"apMacAddress,omitempty"`
	Associated      string `url:"associated,omitempty"`
	BuildingID      string `url:"buildingID,omitempty"`
	CampusID        string `url:"campusID,omitempty"`
	DeviceID        string `url:"deviceID,omitempty"`
	DeviceType      string `url:"deviceType,omitempty"`
	FloorID         string `url:"floorID,omitempty"`
	Format          string `url:"format,omitempty"`
	IPAddress       string `url:"iPAddress,omitempty"`
	Limit           string `url:"limit,omitempty"`
	Manufacturer    string `url:"manufacturer,omitempty"`
	MapElementID    string `url:"mapElementID,omitempty"`
	MapElementLevel string `url:"mapElementLevel,omitempty"`
	Page            string `url:"page,omitempty"`
	RogueApClients  string `url:"rogueApClients,omitempty"`
	Ssid            string `url:"ssid,omitempty"`
	Username        string `url:"username,omitempty"`
}

// LocationDeviceResults provides device location data
type LocationDeviceResults struct {
	// True to inidcate there is next page, false otherwise
	MorePage bool `json:"morePage,omitempty"`

	Querystring LocationDeviceQuery `json:"querystring,omitempty"`

	// The list of device location data, include associated AP devices location data as the first item and each device's mac address and coordinates as following items in the list.
	Results []LocationDevice `json:"results"`

	// True in a successful response.
	Success bool `json:"success,omitempty"`
}

// LocationDevice represents device location data for a single device
type LocationDevice struct {
	// The list of APs.
	ApList []struct {
		ApMacAddress string   `json:"apMacAddress,omitempty"`
		Bands        []string `json:"bands,omitempty"`
		Rssi         int64    `json:"rssi,omitempty"`
	} `json:"apList"`

	// AP Mac Address
	ApMacAddress string `json:"apMacAddress,omitempty"`

	// true|false.  Whether or not a device has connected to a network.
	Associated bool `json:"associated,omitempty"`

	// Band
	Band string `json:"band,omitempty"`

	// unique identifier for a building from the map import process
	BuildingID string `json:"buildingId,omitempty"`

	// unique identifier for a campus from the map import process
	CampusID string `json:"campusId,omitempty"`

	// The UTC time the device state changed.
	// ChangedOn string `json:"changedOn,omitempty"`
	ChangedOn int64 `json:"changedOn,omitempty"`

	// The compute type, possible values are RSSI and AOA.
	ComputeType string `json:"computeType,omitempty"`

	// confidence factor
	ConfidenceFactor int64 `json:"confidenceFactor,omitempty"`

	// The controller IP.
	// Format: ipv4
	Controller string `json:"controller,omitempty"`

	// x and y coordinates of a device.
	Coordinates []float64 `json:"coordinates"`

	// CLIENT, TAG, ROGUE_AP, or ROGUE_CLIENT
	DeviceType string `json:"deviceType,omitempty"`

	// The first time of the device location being detected.
	FirstLocatedAt string `json:"firstLocatedAt,omitempty"`

	// unique identifier for a floor from the map import process\
	FloorID string `json:"floorId,omitempty"`

	// The Geo coordinates of a device.
	GeoCoordinates []float64 `json:"geoCoordinates"`

	// Site hierarchy
	Hierarchy string `json:"hierarchy,omitempty"`

	// IP Address
	IPAddress string `json:"ipAddress,omitempty"`

	// Is MAC Hashed
	IsMacHashed bool `json:"isMacHashed,omitempty"`

	// The last time of the location being detected.
	// Format: date-time
	LastLocationAt string `json:"lastLocationAt,omitempty"`

	// macaddress of the device.
	MacAddress string `json:"macAddress,omitempty"`

	// Manufacturer of the device.
	Manufacturer string `json:"manufacturer,omitempty"`

	// max detected rssi
	MaxDetectedRssi struct {
		ApMacAddress string `json:"apMacAddress,omitempty"`
		Band         string `json:"band,omitempty"`
		Slot         int64  `json:"slot,omitempty"`
		Rssi         int64  `json:"rssi,omitempty"`
		AntennaIndex int64  `json:"antennaIndex,omitempty"`
		LastHeard    int64  `json:"lastHeard,omitempty"`
	} `json:"maxDetectedRssi,omitempty"`

	// The number of detecting APs.
	NumDetectingAps int64 `json:"numDetectingAps,omitempty"`

	// The raw x and y coordinates of a device.
	RawCoordinates []float64 `json:"rawCoordinates"`

	// Source
	Source string `json:"source,omitempty"`

	// SSID
	SSID string `json:"ssid,omitempty"`

	// the tenant unique identifier
	TenantID string `json:"tenantId,omitempty"`

	// Username
	Username string `json:"userName,omitempty"`
}

// ClientCountResponse provides the count for the active devices from Count()
type ClientCountResponse struct {
	Results struct {
		Total int64 `json:"total"`
	} `json:"results"`
	Querystring LocationDeviceQuery `json:"querystring,omitempty"`
	Success     bool                `json:"success"`
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
func (s *ActiveClientsService) ListClients(ctx context.Context, opts *ClientParameters) (LocationDeviceResults, error) {
	ldr := LocationDeviceResults{}
	url := fmt.Sprintf("%s/clients", s.client.BaseURL)
	u, err := addOptions(url, opts)
	if err != nil {
		return ldr, err
	}
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return ldr, err
	}
	if err := s.client.makeRequest(ctx, req, &ldr); err != nil {
		return ldr, err
	}
	return ldr, nil
}

// Count retrieves the active clients count. The API supports searching by a variety of parameters. If no parameters are given, the count of all active clients are returned.
func (s *ActiveClientsService) GetCount(ctx context.Context, opts *ClientParameters) (ClientCountResponse, error) {
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
func (s *ActiveClientsService) ListFloors(ctx context.Context) (ClientFloorsResponse, error) {
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
