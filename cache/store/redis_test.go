package store_test

import (
	"testing"

	"github.com/garyburd/redigo/redis"
	"github.com/thbourlove/design-patterns/cache/store"
)

type MockRedisConn struct {
	redis.Conn
}

func (*MockRedisConn) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	switch commandName {
	case "GET":
		return "bar", nil
	case "SET":
		return nil, nil
	default:
		return nil, nil
	}
}

func NewMockRedisConn() *MockRedisConn {
	return &MockRedisConn{}
}

func TestRedisGetSet(t *testing.T) {
	conn := NewMockRedisConn()
	store := store.NewRedis(conn)
	key := "foo"
	store.Set(key, "bar")
	value := store.Get(key)
	if value != "bar" {
		t.Errorf("Value = %s; want bar", value)
	}
}
