package datastore

import (
	"log"

	"github.com/Shuto-san/4babe_api/config"
	"github.com/gomodule/redigo/redis"
)

func NewKVS() redis.Conn {

	conn, err := redis.Dial(config.C.Redis.Net, config.C.Redis.Addr)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}
