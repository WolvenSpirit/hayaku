package cache

import (
	"errors"
	"log"

	"github.com/go-redis/redis"
)

// Redis represents the Redis configuration
type Redis struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	DBI      int    `yaml:"DBI"`
}

var (
	// Client should be an initialized Redis instance
	Client *redis.Client
)

// ConnectRedis creates the connection
func ConnectRedis(host, password string, dbi int) {
	opt := &redis.Options{
		Addr:     host,
		Password: password,
		DB:       dbi,
	}
	Client = redis.NewClient(opt)
	if Client == nil {
		panic(errors.New("error: failed to connect to Redis instance"))
	}
	r, err := Client.Ping().Result()
	log.Printf("Redis connected: %+v", r)
	panic(err)
}
