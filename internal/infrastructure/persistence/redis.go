package persistence

import (
	"context"
	"encoding/json"
	"time"

	"project-test/internal/domain/repository"
	//"project-test/internal/infrastructure/utils/helpers"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(client *redis.Client) repository.IRedis {
	return &Redis{
		client: client,
	}
}

func (pr *Redis) Set(key string, value interface{}, expiration time.Duration) error {
	return pr.client.Set(context.Background(), key, value, expiration).Err()
}

func (pr *Redis) Get(key string) (string, error) {
	val, err := pr.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (pr *Redis) SetWithStruct(key string, value interface{}, expiration time.Duration) error {
	valueJson := JSONMarshalString(value)
	return pr.client.Set(context.Background(), key, valueJson, expiration).Err()
}

func (pr *Redis) GetWithStruct(key string, structData interface{}) error {
	val, err := pr.client.Get(context.Background(), key).Result()
	if err != nil {
		return err
	}
	if val != "" {
		json.Unmarshal([]byte(val), &structData)
		return nil
	}
	return nil
}

func (pr *Redis) Del(key string) error {
	_, err := pr.client.Del(context.Background(), key).Result()
	if err != nil {
		return err
	}
	return nil
}

func (pr *Redis) SetNX(keylock, value string, expiration time.Duration) (bool, error) {
	// lấy lock
	return pr.client.SetNX(context.Background(), keylock, value, expiration).Result()
}

func JSONMarshalString(v interface{}) string {
	result := "{}"
	b, err := json.Marshal(v)
	if err != nil {
		return result
	}
	result = string(b)
	return result
}
