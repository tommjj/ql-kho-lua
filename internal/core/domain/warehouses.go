package domain

import (
	"strconv"
	"strings"
)

type WarehouseItem struct {
	Type     Rice `json:"type"`
	Quantity int  `json:"quantity"`
}

type Warehouse struct {
	ID           int              `json:"id"`
	Name         string           `json:"name"`
	Location     string           `json:"location"`
	Capacity     int              `json:"capacity"`
	UsedCapacity *int             `json:"used_capacity,omitempty"`
	Image        string           `json:"image"`
	Items        *[]WarehouseItem `json:"items,omitempty"`
}

func (s *Warehouse) ParseLocation() (float64, float64, error) {
	before, after, found := strings.Cut(s.Location, ",")
	if !found {
		return 0, 0, ErrInvalidLocation
	}

	latitude, err := strconv.ParseFloat(strings.TrimSpace(before), 64)
	if err != nil {
		return 0, 0, ErrInvalidLocation
	}
	longitude, err := strconv.ParseFloat(strings.TrimSpace(after), 64)
	if err != nil {
		return 0, 0, ErrInvalidLocation
	}

	if latitude < -90 || latitude > 90 || longitude < -180 || longitude > 180 {
		return 0, 0, ErrInvalidLocation
	}

	return latitude, longitude, nil
}
