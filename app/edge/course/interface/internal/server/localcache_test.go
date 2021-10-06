package server

import (
	"fmt"
	"github.com/goburrow/cache"
	"sync"
	"testing"
	"time"
)

func TestLocalCache(t *testing.T) {
	load := func(k cache.Key) (cache.Value, error) {
		time.Sleep(100 * time.Millisecond) // Slow task
		return fmt.Sprintf("%d", k), nil
	}
	// Create a loading cache
	c := cache.NewLoadingCache(load,
		cache.WithMaximumSize(100),                 // Limit number of entries in the cache.
		cache.WithExpireAfterAccess(1*time.Minute), // Expire entries after 1 minute since last accessed.
		cache.WithRefreshAfterWrite(2*time.Minute), // Expire entries after 2 minutes since last created.
	)

	//getTicker := time.Tick(100 * time.Millisecond)
	//reportTicker := time.Tick(5 * time.Second)
	//for {
	//	select {
	//	case <-getTicker:
	//		_, _ = c.Get(rand.Intn(200))
	//	case <-reportTicker:
	//		st := cache.Stats{}
	//		c.Stats(&st)
	//		t.Logf("%+v\n", st)
	//	}
	//}

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			v, _ := c.Get(i % 3)
			t.Logf("%d: %s\n", i % 3, v)
		}(i)
	}
	wg.Wait()
}
