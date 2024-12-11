package rdb_test

import (
	"sync"
	"testing"

	"github.com/redis/go-redis/v9"

	"redis-pub-sub/rdb"
)


func TestGetRedisClient(t *testing.T) {
    const goroutines = 100
    var wg sync.WaitGroup
    wg.Add(goroutines)

    clients := make([]*redis.Client, goroutines)
    for i := 0; i < goroutines; i++ {
        go func(index int) {
            defer wg.Done()
            clients[index] = rdb.GetRedisClient()
        }(i)
    }

    wg.Wait()

    for i := 1; i < goroutines; i++ {
       if clients[0] != clients[i] {
		t.Errorf("Client instances at index 0 and %d are not the same", i)
	   }
    }
}