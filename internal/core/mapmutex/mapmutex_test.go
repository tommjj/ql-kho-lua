package mapmutex

import (
	"testing"
)

func Test(t *testing.T) {
	l := Mapmutex{}

	l.Lock(1)

	l.Lock(1)
}
