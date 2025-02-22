package utils

import (
	"errors"
	"fmt"
)

// ToLocationString is a helper to conv location to string
func LocationToString(location []float64) (string, error) {
	if len(location) != 2 {
		return "", errors.New("location invalid")
	}

	latitude := location[0]

	longitude := location[1]

	if latitude < -90 || latitude > 90 || longitude < -180 || longitude > 180 {
		return "", errors.New("location invalid")

	}

	return fmt.Sprintf("%v, %v", latitude, longitude), nil
}
