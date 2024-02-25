package sessions

import (
	"errors"
	"fmt"

	"github.com/jackjohn7/smllnk/db/connections"
	"github.com/jackjohn7/smllnk/db/models"
	"github.com/redis/go-redis/v9"
)

type RedisSessionStore struct {
	conn *redis.Client
}

func NewRedisSessionStore() (store *RedisSessionStore, err error) {
	conn, err := connections.GetRedisConnection()
	store = &RedisSessionStore{
		conn: conn,
	}
	return store, err
}

func (store *RedisSessionStore) Create(user *models.User, userAgent string) (session *Session, err error) {
	err = errors.New("Unimplemented")
	return
}

func (store *RedisSessionStore) Get(id string) (session *Session, err error) {
	fmt.Println("hitting this")
	err = errors.New("Unimplemented")
	return
}

func (store *RedisSessionStore) Delete(id string) (ok bool) {
	return
}

func (store *RedisSessionStore) Refresh(id string) (session *Session, err error) {
	err = errors.New("Unimplemented")
	return
}
