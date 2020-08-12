package redis_rate_test

import (
	"fmt"

	"github.com/hossamElfar/redis_rate/v9"
	redis "gopkg.in/redis.v5"
)

func ExampleNewLimiter() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_ = rdb.FlushDb().Err()

	limiter := redis_rate.NewLimiter(rdb)
	res, err := limiter.Allow("project:123", redis_rate.PerSecond(10))
	if err != nil {
		panic(err)
	}
	fmt.Println("allowed", res.Allowed, "remaining", res.Remaining)
	// Output: allowed 1 remaining 9
}
