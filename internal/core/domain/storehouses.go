package domain

import (
	"strconv"
	"strings"
)

type StorehouseItem struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type Storehouse struct {
	ID       int              `json:"id"`
	Name     string           `json:"name"`
	Location string           `json:"location"`
	Capacity int              `json:"capacity"`
	Items    []StorehouseItem `json:"items,omitempty"`
}

func (s *Storehouse) ParseLocation() (float64, float64, error) {
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
