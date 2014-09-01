package store

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

type Redis struct {
	conn redis.Conn
}

func (r *Redis) Get(key string) (value string) {
	value, err := redis.String(r.conn.Do("GET", key))
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func (r *Redis) Set(key, value string) {
	_, err := r.conn.Do("SET", key, value)
	if err != nil {
		log.Fatal(err)
	}
}

func NewRedis(c redis.Conn) *Redis {
	return &Redis{c}
}
