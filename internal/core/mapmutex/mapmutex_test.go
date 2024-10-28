package mapmutex

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	l := Mapmutex{}

	start := time.Now()
	l.Lock(1)
	fmt.Println(time.Since(start))

	time.AfterFunc(time.Second*2, func() {
		l.UnLock(1)
	})
	l.Lock(1)
}
