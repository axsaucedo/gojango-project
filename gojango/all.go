package gojango

import (
	badger "github.com/dgraph-io/badger/v3"
	"github.com/gomodule/redigo/redis"
)

type all struct {
	Conn *redis.Pool
	Bad  *badger.DB
}
