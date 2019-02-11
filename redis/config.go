package redis

// session redis config

import (
	"errors"
	"strings"
)

type Config struct {

	// Addr=Host:Port
	Addr string

	// Maximum number of idle connections in the redis server pool.
	MaxIdle int

	// Close connections after remaining idle for this duration. If the value
	// is zero, then idle connections are not closed. Applications should set
	// the timeout to a value less than the server's timeout.
	// (s)
	IdleTimeout int64

	// redis server conn auth, default ""
	Password string

	// select db number, default 0
	DbNumber int

	// sessionId as redis key prefix
	KeyPrefix string

	// session value serialize func
	SerializeFunc func(data map[string]interface{}) ([]byte, error)

	// session value unSerialize func
	UnSerializeFunc func(data []byte) (map[string]interface{}, error)
}

func (mc *Config) Name() string {
	return ProviderName
}

//tcp@127.0.0.1:6379/1/
func ParceRedisAddr(Params string) (proto, ipaddr, db string, err error) {
	arr := strings.Split(Params, "@")
	if len(arr) != 2 {
		return "", "", "", errors.New("Failed full redis queue format")
	}
	in := strings.Split(arr[1], "/")
	if len(in) != 3 {
		return "", "", "", errors.New("Failed full redis queue format")
	}

	return arr[0], in[0], in[1], nil
}
