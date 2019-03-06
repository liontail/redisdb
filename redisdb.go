package redisdb

import (
	"log"
	"time"

	"github.com/go-redis/redis"
)

// RedisDB struct to store and retrive data from redis
type RedisDB struct {
	Client   *redis.Client
	Duration time.Duration
}

var redisDB *RedisDB

func Initial(url, pass string) {
	cli := NewRedisClient(url, pass)
	if cli == nil {
		log.Panic("Redis error: Cannot connect to redis")
		return
	}
	redisDB = &RedisDB{
		Client:   cli,
		Duration: time.Hour * 24,
	}

}

func (redis *RedisDB) SetDefaultExpired(dr time.Duration) {
	redis.Duration = dr
}

func GetRedisDB() *RedisDB {
	return redisDB
}

// New will create a new Redis client
func NewRedisClient(uri, pass string) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     uri,
		Password: pass,
	})

	return client
}

// Close the cache
func (c *RedisDB) Close() {
	c.Client.Close()
}

// Get will return the value from redis
func (c *RedisDB) Get(key string) (string, error) {
	str := c.Client.Get(key)
	return str.Val(), str.Err()
}

// Set the value
func (c *RedisDB) Set(key string, value interface{}) error {
	_, err := c.Client.Set(key, value, c.Duration).Result()
	return err
}
