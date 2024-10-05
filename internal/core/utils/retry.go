package utils

import (
	"time"
)

// Retry is a helper func to retry a function fc a specified number of times if it encounters an error.
func Retry[T any](fc func() (T, error), duration time.Duration, times int) (T, error) {
	var result T
	var err error

	result, err = fc()
	if err == nil {
		return result, nil
	}

	for range times {
		result, err = fc()
		if err == nil {
			return result, err
		}
		time.Sleep(duration)
	}
	return result, err
}
