package cache

import "time"

type ICache interface{
	Get(key interface{})(interface{},error)
	Set(key, value interface{})error
	SetEx(key, value interface{}, expiration time.Duration)error
}

