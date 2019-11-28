package redisdb

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

// RedisDB struct to store and retrive data from redis
type RedisDB struct {
	Client   *redis.Client
	Duration time.Duration
}

var redisDB *RedisDB

func Initial(url, pass string) (*RedisDB, error) {
	cli := NewRedisClient(url, pass, 0)
	if cli == nil {
		return nil, errors.New("Cannot Create Redis Client")
	}
	if err := cli.Ping().Err(); err != nil {
		return nil, err
	}
	redisDB = &RedisDB{
		Client:   cli,
		Duration: time.Hour * 24,
	}

	return redisDB, nil
}

func InitWithOptions(options *redis.Options) (*RedisDB, error) {
	cli := redis.NewClient(options)
	if cli == nil {
		return nil, errors.New("Cannot Create Redis Client")
	}
	if err := cli.Ping().Err(); err != nil {
		return nil, err
	}
	redisDB = &RedisDB{
		Client:   cli,
		Duration: time.Hour * 24,
	}

	return redisDB, nil
}

func (redis *RedisDB) SetDefaultExpired(dr time.Duration) {
	redis.Duration = dr
}

func GetRedisDB() *RedisDB {
	return redisDB
}

// New will create a new Redis client
func NewRedisClient(uri, pass string, db int) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     uri,
		Password: pass,
		DB:       db,
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
