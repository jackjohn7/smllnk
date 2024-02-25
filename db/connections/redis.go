package connections

import (
	"errors"

	"github.com/jackjohn7/smllnk/environment"
	"github.com/redis/go-redis/v9"
)

var conn *redis.Client

func GetRedisConnection() (*redis.Client, error) {
	err := error(nil)
	if conn == nil {
		// create new connection
		conn = redis.NewClient(&redis.Options{
			Addr:     environment.Env.DbEnv.REDIS_URL,
			Password: environment.Env.DbEnv.REDIS_PW,
			DB:       0,
		})
		if conn == nil {
			err = errors.New("Failed to connect")
		}
	}
	return conn, err
}
