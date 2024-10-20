package mapmutex

import (
	"testing"
	"time"
)

func Test(t *testing.T) {
	l := Mapmutex{}

	l.Lock(1)
	time.AfterFunc(time.Second*2, func() {
		l.UnLock(1)
	})
	l.Lock(1)
}
