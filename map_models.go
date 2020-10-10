package dnas

type MapHierarchyResponse struct {
	// Top level item for Map Items
	Map []MapItem `json:"map"`
}

type MapElementResponse struct {
	Success bool    `json:"success,omitempty"`
	Map     MapItem `json:"map"`
}

type MapItem struct {
	Address string `json:"address,omitempty"`

	// The date and time of the level being created.
	// Format: date-time
	CreatedOn string `json:"createdOn,omitempty"`

	// details
	Details MapItemDetails `json:"details,omitempty"`

	// Campus unique identifier.
	ID string `json:"id,omitempty"`

	// imported Id
	ImportedID string `json:"importedId,omitempty"`

	// The level in the hireachy.
	Level string `json:"level,omitempty"`

	// Campus name.
	Name string `json:"name,omitempty"`

	// relationship data
	RelationshipData MapItemRelationshipData `json:"relationshipData,omitempty"`

	SourceType string `json:"sourceType,omitempty"`
}

type MapItemDetails struct {
	CalibrationModelRef string `json:"calibrationModelRef,omitempty"`

	FloorNumber int64 `json:"floorNumber,omitempty"`

	GpsMarkers []string `json:"GpsMarkers,omitempty"`

	// The level's map height.
	Height float64 `json:"height,omitempty"`

	// image
	Image MapResImage `json:"image,omitempty"`

	InclusionExclusionRegion []map[string]interface{} `json:"inclusionExclusionRegion,omitempty"`

	Latitude float64 `json:"latitude,omitempty"`

	// The level's map length.
	Length float64 `json:"length,omitempty"`

	Longitude float64 `json:"longitude,omitempty"`

	Obstacles []string `json:"obstacles,omitempty"`

	OffsetX float64 `json:"offsetX,omitempty"`

	OffsetY float64 `json:"offsetY,omitempty"`

	// The level's map width.
	Width float64 `json:"width,omitempty"`
}

type MapInclusionExclusionRegionItem struct {

	// Indicate how to handle the vertices on map.
	Type string `json:"type,omitempty"`

	// The number of vertices.
	Vertices int64 `json:"vertices,omitempty"`
}

type MapItemCorner struct {
	// x
	X float64 `json:"x,omitempty"`
	// y
	Y float64 `json:"y,omitempty"`
	// z
	Z float64 `json:"z,omitempty"`
	// Unit
	Unit string `json:"unit,omitempty"`
}

type MapItemRelationshipData struct {

	// ancestor ids
	AncestorIds []string `json:"ancestorIds"`

	// ancestors
	Ancestors []string `json:"ancestors"`

	// children
	Children []MapItem `json:"children"`

	// parent
	Parent interface{} `json:"parent,omitempty"`
}

type MapResImage struct {
	Cksum string `json:"cksum,omitempty"`

	// color depth
	ColorDepth int64 `json:"colorDepth,omitempty"`

	// height
	Height int64 `json:"height,omitempty"`

	// Indicate if the source image is compressed(true) or not(false).
	ImageCompressed bool `json:"imageCompressed,omitempty"`

	// The image for the map.
	ImageName string `json:"imageName,omitempty"`

	// max resolution
	MaxResolution int64 `json:"maxResolution,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// source file
	SourceFile string `json:"sourceFile,omitempty"`

	// Indicate if the source image is valid(true) or not(false).
	ValidImageSupplied bool `json:"validImageSupplied,omitempty"`

	// width
	Width int64 `json:"width,omitempty"`

	// zoom level
	ZoomLevel int64 `json:"zoomLevel,omitempty"`
}
