package redis

import (
	"github.com/go-redis/redis"
	"github.com/somprabhsharma/the-lazy-traveler/constants"
	"github.com/somprabhsharma/the-lazy-traveler/utils/logger"
	"time"
)

const (
	maxRetries      = 10   //maximum number of retries if connection is lost
	maxRetryBackOff = 3000 //time after which each retry will happen
)

// Client redis client
type Client struct {
	client *redis.Client
}

//NewClient initializes redis client with proper configuration
func NewClient() *Client {
	var redisClient = &Client{}
	opt, err := redis.ParseURL(constants.Env.RedisURL)
	if err != nil {
		logger.Err("Redis", "Error while parsing redis url", err)
	}

	redisClient.client = redis.NewClient(&redis.Options{
		Addr:     opt.Addr,
		DB:       opt.DB,
		Password: opt.Password,
		OnConnect: func(conn *redis.Conn) error {
			logger.Info("Redis", "successfully connected to redis.")
			return nil
		},
		MaxRetries:      maxRetries,
		MaxRetryBackoff: maxRetryBackOff,
	})

	_, err = redisClient.client.Ping().Result()
	if err != nil {
		logger.Err("Redis", "Error while connecting to redis.", err)
	}
	return redisClient
}

// Put value corresponding to key in redis
func (r *Client) Put(key, value string, ttl time.Duration) error {
	err := r.client.Set(key, value, ttl).Err()
	return err
}

// Get value for given key
func (r *Client) Get(key string) (string, error) {
	value, err := r.client.Get(key).Result()
	return value, err
}

// Delete value for given key
func (r *Client) Delete(key string) error {
	_, err := r.client.Del(key).Result()
	return err
}
