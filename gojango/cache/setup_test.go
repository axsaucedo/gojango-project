package cache

import (
	"os"
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
)

const (
	redisHost = "localhost:6379"
)

var testRedisCache RedisCache

func TestMain(m *testing.M) {
	//s, err := miniredis.Run()
	//if err != nil {
	//    panic(err)
	//}
	//defer s.Close()

	pool := redis.Pool{
		MaxIdle:     50,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisHost)
		},
	}

	testRedisCache.Conn = &pool
	testRedisCache.Prefix = "test-gojango"

	defer testRedisCache.Conn.Close()

	os.Exit(m.Run())
}
