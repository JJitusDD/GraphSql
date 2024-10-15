package repository

import "time"

type IRedis interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
	Del(key string) error
	GetWithStruct(key string, structData interface{}) error
	SetWithStruct(key string, value interface{}, expiration time.Duration) error
	SetNX(keylock, value string, expiration time.Duration) (bool, error)
}
