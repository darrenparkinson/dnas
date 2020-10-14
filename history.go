package dnas

import (
	"context"
	"fmt"
	"net/http"
)

// HistoryParameters represent the options for GetHistory()
type HistoryParameters struct {
	// ApMacAddress The mac address of the Access Point (AP).  Available for associated clients only.
	ApMacAddress *string `url:"apMacAddress,omitempty"`

	// BuildingID Unique identifier for a building from the map import process
	BuildingID *string `url:"buildingId,omitempty"`

	// CampusID Unique identifier for a campus from the map import process
	CampusID *string `url:"campusId,omitempty"`

	// DeviceID The device unique identifier, for example the device macAddress.
	DeviceID *string `url:"deviceId,omitempty"`

	// The end time(in epoc time format based on milisecond) of the range of the time zone the request is initiated.
	EndTime *string `url:"endTime,omitempty"`

	// FloorID Unique identifier for a floor from the map import process
	FloorID *string `url:"floorId,omitempty"`

	// Format Indicate if using geojson, value is "geojson" if so.
	Format *string `url:"format,omitempty"`

	// Ssid Wifi service set identifier (SSID).  Available for associated clients only.
	Ssid *string `url:"ssid,omitempty"`

	// The start time(in epoc time format based on milisecond) of the range of the time zone the request is initiated.
	StartTime *string `url:"startTime,omitempty"`

	// The time zone the request is initiated, default is time zone 0(UTC time).
	TimeZone *int64 `url:"timeZone,omitempty"`

	// Username The user name of the connected user. Available for associated clients only.
	Username *string `url:"username,omitempty"`
}

// HistoryCountParameters represent the options for GetCount()
type HistoryCountParameters struct {
	// BuildingID Unique identifier for a building from the map import process
	BuildingID *string `url:"buildingId,omitempty"`

	// CampusID Unique identifier for a campus from the map import process
	CampusID *string `url:"campusId,omitempty"`

	// The end time(in epoc time format based on milisecond) of the range of the time zone the request is initiated.
	EndTime *string `url:"endTime,omitempty"`

	// FloorID Unique identifier for a floor from the map import process
	FloorID *string `url:"floorId,omitempty"`

	// The start time(in epoc time format based on milisecond) of the range of the time zone the request is initiated.
	StartTime *string `url:"startTime,omitempty"`

	// The time zone the request is initiated, default is time zone 0(UTC time).
	TimeZone *int64 `url:"timeZone,omitempty"`
}

// HistoryClientsParameters represent the options for GetClientsHistory()
type HistoryClientsParameters struct {
	// ApMacAddress The mac address of the Access Point (AP).  Available for associated clients only.
	ApMacAddress *string `url:"apMacAddress,omitempty"`

	// BuildingID Unique identifier for a building from the map import process
	BuildingID *string `url:"buildingId,omitempty"`

	// CampusID Unique identifier for a campus from the map import process
	CampusID *string `url:"campusId,omitempty"`

	// The end time(in epoc time format based on milisecond) of the range of the time zone the request is initiated.
	EndTime *string `url:"endTime,omitempty"`

	// FloorID Unique identifier for a floor from the map import process
	FloorID *string `url:"floorId,omitempty"`

	// The radius, it should go with x and y.
	Radius *float64 `url:"radius,omitempty"`

	// Ssid Wifi service set identifier (SSID).  Available for associated clients only.
	Ssid *string `url:"ssid,omitempty"`

	// The start time(in epoc time format based on milisecond) of the range of the time zone the request is initiated.
	StartTime *string `url:"startTime,omitempty"`

	// The time zone the request is initiated, default is time zone 0(UTC time).
	TimeZone *int64 `url:"timeZone,omitempty"`

	// The x coordinate of the radius center.
	X *float64 `url:"x,omitempty"`

	// The y coordinate of the radius center.
	Y *float64 `url:"y,omitempty"`
}

// HistoryClientsDeviceParameters represent the options for GetClient()
type HistoryClientsDeviceParameters struct {
	// ApMacAddress The mac address of the Access Point (AP).  Available for associated clients only.
	ApMacAddress *string `url:"apMacAddress,omitempty"`

	// BuildingID Unique identifier for a building from the map import process
	BuildingID *string `url:"buildingId,omitempty"`

	// CampusID Unique identifier for a campus from the map import process
	CampusID *string `url:"campusId,omitempty"`

	// The end time(in epoc time format based on milisecond) of the range of the time zone the request is initiated.
	EndTime *string `url:"endTime,omitempty"`

	// FloorID Unique identifier for a floor from the map import process
	FloorID *string `url:"floorId,omitempty"`

	// Format Indicate if using geojson, value is "geojson" if so.
	Format *string `url:"format,omitempty"`

	// The radius, it should go with x and y.
	Radius *float64 `url:"radius,omitempty"`

	// Ssid Wifi service set identifier (SSID).  Available for associated clients only.
	Ssid *string `url:"ssid,omitempty"`

	// The start time(in epoc time format based on milisecond) of the range of the time zone the request is initiated.
	StartTime *string `url:"startTime,omitempty"`

	// The time zone the request is initiated, default is time zone 0(UTC time).
	TimeZone *int64 `url:"timeZone,omitempty"`

	// Username The user name of the connected user. Available for associated clients only.
	Username *string `url:"username,omitempty"`

	// The x coordinate of the radius center.
	X *float64 `url:"x,omitempty"`

	// The y coordinate of the radius center.
	Y *float64 `url:"y,omitempty"`
}

// HistoryCountResponse provides the count for the active, inactive, missing or all the access points from GetCount()
type HistoryCountResponse struct {
	// Count of clients given filter
	Count string `json:"count"`
}

// HistoryClientsResponse contains the response from ListClients which consists of an array of Mac Addresses
type HistoryClientsResponse []struct {
	MacAddress string `json:"macAddress,omitempty"`
}

// HistoryClientsDeviceResponse  contains the response from GetClient()
type HistoryClientsDeviceResponse []struct {
	FloorID         string    `json:"floorId"`
	SourceTimestamp int64     `json:"sourceTimestamp"`
	Coordinates     []float64 `json:"coordinates"`
	Associated      bool      `json:"associated"`
	AssociatedApmac string    `json:"associatedApmac"`
}

// HistoryResponse represents the CSV response from GetHistory()
type HistoryResponse struct {
	Results []HistoryItem `json:"results"`
}

// HistoryItem represents a single item in the HistoryResponse from the CSV output of GetHistory()
// Note that all results are returned as string since they're coming from CSV.  This provides flexibility for conversion.
type HistoryItem struct {
	TenantID                string `json:"tenantid"`
	MacAddress              string `json:"macaddress"`
	DeviceType              string `json:"devicetype"`
	CampusID                string `json:"campusid"`
	BuildingID              string `json:"buildingid"`
	FloorID                 string `json:"floorid"`
	FloorHierarchy          string `json:"floorhierarchy"`
	CoordinateX             string `json:"coordinatex"`
	CoordinateY             string `json:"coordinatey"`
	SourceTimestamp         string `json:"sourcetimestamp"`
	MaxDetectedApMac        string `json:"maxdetectedapmac"`
	MaxDetectedBand         string `json:"maxdetectedband"`
	DetectingControllers    string `json:"detectingcontrollers"`
	FirstActiveAt           string `json:"firstactiveat"`
	LocatedSinceActiveCount string `json:"locatedsinceactivecount"`
	ChangedOn               string `json:"changedon"`
	Manufacturer            string `json:"manufacturer"`
	Associated              string `json:"associated"`
	MaxDetectedRssi         string `json:"maxdetectedrssi"`
	Ssid                    string `json:"ssid"`
	Username                string `json:"username"`
	AssociatedApMac         string `json:"associatedapmac"`
	AssociatedApRssi        string `json:"associatedaprssi"`
	MaxDetectedSlot         string `json:"maxdetectedslot"`
	IPAddress               string `json:"ipaddress"`
	StaticDevice            string `json:"staticdevic"`
	RecordType              string `json:"recordtype"`
	ComputeType             string `json:"computetype"`
	Source                  string `json:"source"`
	MacHashed               string `json:"machashed"`
}

// GetHistory retrieves a small amount of clients history to csv format.
// If startTime and endTime is not given, the time period is last 24 hours.
// If records amount is more than 50K, the user receives error response and indicates the time range needs to be reduced.
func (s *HistoryService) GetHistory(ctx context.Context, opts *HistoryParameters) (HistoryResponse, error) {
	var hr HistoryResponse
	url := fmt.Sprintf("%s/history", s.client.BaseURL)
	u, err := addOptions(url, opts)
	if err != nil {
		return hr, err
	}
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return hr, err
	}
	var records [][]string
	if err := s.client.makeRequest(ctx, req, &records); err != nil {
		return hr, err
	}
	for _, record := range records[1:] {
		hr.Results = append(hr.Results, HistoryItem{
			TenantID:                record[0],
			MacAddress:              record[1],
			DeviceType:              record[2],
			CampusID:                record[3],
			BuildingID:              record[4],
			FloorID:                 record[5],
			FloorHierarchy:          record[6],
			CoordinateX:             record[7],
			CoordinateY:             record[8],
			SourceTimestamp:         record[9],
			MaxDetectedApMac:        record[10],
			MaxDetectedBand:         record[11],
			DetectingControllers:    record[12],
			FirstActiveAt:           record[13],
			LocatedSinceActiveCount: record[14],
			ChangedOn:               record[15],
			Manufacturer:            record[16],
			Associated:              record[17],
			MaxDetectedRssi:         record[18],
			Ssid:                    record[19],
			Username:                record[20],
			AssociatedApMac:         record[21],
			AssociatedApRssi:        record[22],
			MaxDetectedSlot:         record[23],
			IPAddress:               record[24],
			StaticDevice:            record[25],
			RecordType:              record[26],
			ComputeType:             record[27],
			Source:                  record[28],
			MacHashed:               record[29],
		})
	}
	return hr, nil
}

// GetCount retrieves the clients history records amount in given time range.
// If startTime and endTime is not being given, the time range is last 24 hours.
func (s *HistoryService) GetCount(ctx context.Context, opts *HistoryCountParameters) (HistoryCountResponse, error) {
	hcr := HistoryCountResponse{}
	url := fmt.Sprintf("%s/history/records/count", s.client.BaseURL)
	u, err := addOptions(url, opts)
	if err != nil {
		return hcr, err
	}
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return hcr, err
	}
	if err := s.client.makeRequest(ctx, req, &hcr); err != nil {
		return hcr, err
	}
	return hcr, nil
}

// ListClients retrieves the clients mac address list by using filters.
// If startTime and endTime are not given, all the clients' mac addresses in the last 1 day are being returned.
func (s *HistoryService) ListClients(ctx context.Context, opts *HistoryClientsParameters) (HistoryClientsResponse, error) {
	hcr := HistoryClientsResponse{}
	url := fmt.Sprintf("%s/history/clients", s.client.BaseURL)
	u, err := addOptions(url, opts)
	if err != nil {
		return hcr, err
	}
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return hcr, err
	}
	if err := s.client.makeRequest(ctx, req, &hcr); err != nil {
		return hcr, err
	}
	return hcr, nil
}

// GetClient retrieves the given client history details by using filters.
// Pagination is provided. The startTime and endTime time peroid is at most 1 day, if not being given, then the last 1 day's history of the client is returned.
// Default page is 1, 20k items per page (Note - 20k is requested by UI, pending to adjust to smaller page size based on test result).
func (s *HistoryService) GetClient(ctx context.Context, deviceID string, opts *HistoryClientsParameters) (HistoryClientsDeviceResponse, error) {
	hcdr := HistoryClientsDeviceResponse{}
	url := fmt.Sprintf("%s/history/clients/%s", s.client.BaseURL, deviceID)
	u, err := addOptions(url, opts)
	if err != nil {
		return hcdr, err
	}
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return hcdr, err
	}
	if err := s.client.makeRequest(ctx, req, &hcdr); err != nil {
		return hcdr, err
	}
	return hcdr, nil
}
