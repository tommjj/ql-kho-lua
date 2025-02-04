package mapmutex_test

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tommjj/ql-kho-lua/internal/core/mapmutex"
)

func TestMapmutex_LockUnlock(t *testing.T) {
	mm := &mapmutex.Mapmutex{}
	key := "test_key"

	mm.Lock(key)
	locked := true

	go func() {
		mm.Lock(key)
		locked = false // Nếu khóa được, cập nhật biến
		mm.UnLock(key)
	}()

	time.Sleep(100 * time.Millisecond)
	assert.True(t, locked, "Goroutine không nên lấy được lock khi chưa unlock")

	mm.UnLock(key)
	time.Sleep(100 * time.Millisecond)
	assert.False(t, locked, "Goroutine phải lấy được lock sau khi unlock")
}

func TestMapmutex_ConcurrentAccess(t *testing.T) {
	mm := &mapmutex.Mapmutex{}
	key := "test_key"
	var wg sync.WaitGroup
	count := 0

	numGoroutines := 10
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			mm.Lock(key)
			count++ // Chỉ một goroutine có thể tăng giá trị tại một thời điểm
			mm.UnLock(key)
			wg.Done()
		}()
	}

	wg.Wait()
	assert.Equal(t, numGoroutines, count, "Giá trị count phải đúng với số goroutine")
}

func TestMapmutex_MultipleKeys(t *testing.T) {
	mm := &mapmutex.Mapmutex{}
	keys := []string{"key1", "key2", "key3"}
	var wg sync.WaitGroup
	results := make(map[string]int)
	var mu sync.Mutex

	numGoroutines := 10
	wg.Add(numGoroutines * len(keys))

	for _, key := range keys {
		for i := 0; i < numGoroutines; i++ {
			go func(k string) {
				mm.Lock(k)
				mu.Lock()
				results[k]++ // Đảm bảo từng khóa chỉ có một goroutine cập nhật giá trị
				mu.Unlock()
				mm.UnLock(k)
				wg.Done()
			}(key)
		}
	}

	wg.Wait()

	for _, key := range keys {
		assert.Equal(t, numGoroutines, results[key], "Giá trị count của từng khóa phải đúng với số goroutine")
	}
}
