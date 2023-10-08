package config

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var rdClient *redis.Client
var expTime = 30 * 24 * 60 * 60 * time.Second

type Client struct {
}

func InitRedis() (*Client, error) {
	rdClient = redis.NewClient(&redis.Options{
		Addr:     ConfigRedis["addr"].(string),
		Password: ConfigRedis["password"].(string),
		DB:       ConfigRedis["db"].(int),
	})
	_, err := rdClient.Ping(context.Background()).Result()
	return &Client{}, err
}
func (r *Client) Rset(key string, val any) error {
	return rdClient.Set(context.Background(), key, val, expTime).Err()
}
func (r *Client) Rincr(key string) error {
	return rdClient.Incr(context.Background(), key).Err()
}
func (r *Client) Rget(key string) (any, error) {
	return rdClient.Get(context.Background(), key).Result()
}
func (r *Client) Rdel(key ...string) error {
	return rdClient.Del(context.Background(), key...).Err()
}
