package dnas

import (
	"context"
	"fmt"
	"net/http"
)

func (s *MapService) GetHierarchy(ctx context.Context) (MapHierarchyResponse, error) {
	mhr := MapHierarchyResponse{}
	url := fmt.Sprintf("%s/map/hierarchy", s.client.BaseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return mhr, err
	}
	if err := s.client.makeRequest(ctx, req, &mhr); err != nil {
		return mhr, err
	}
	return mhr, nil
}

func (s *MapService) GetMapElement(ctx context.Context, id string) (MapElementResponse, error) {
	mer := MapElementResponse{}
	url := fmt.Sprintf("%s/map/elements/%s", s.client.BaseURL, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return mer, err
	}
	if err := s.client.makeRequest(ctx, req, &mer); err != nil {
		return mer, err
	}
	return mer, nil
}
