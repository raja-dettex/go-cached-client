package client

import "time"

type GoCacheClient interface {
	Set(string, string, int) (string, error)
	Get(string) (string, error)
	Has(string) (string, error)
	Delete(string) (string, error)
}

type CacheData map[string]CacheObj

type CacheObj struct {
	Value string
	TTL   time.Duration
}
