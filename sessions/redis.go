package sessions

import (
	"context"
	"errors"
	"time"

	"github.com/jackjohn7/smllnk/db/connections"
	"github.com/jackjohn7/smllnk/db/models"
	"github.com/jackjohn7/smllnk/utils"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

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

func (store *RedisSessionStore) Create(user *models.User, userAgent string) (*Session, error) {
	sessionId, err := utils.GenerateSessionId()
	if err != nil {
		return nil, err
	}
	expiry := time.Now().Add(time.Hour * 24 * 10)
	session := &Session{
		Id:          sessionId,
		UserId:      user.Id,
		UserAgent:   userAgent,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
		ExpiresAt:   expiry,
	}
	// create session with expiry
	err = store.conn.HSet(ctx, sessionId, session).Err()
	if err == nil {
		err = store.conn.ExpireAt(ctx, sessionId, expiry).Err()
	}
	return session, err
}

func (store *RedisSessionStore) Get(id string) (*Session, error) {
	session := Session{}
	err := store.conn.HGetAll(ctx, id).Scan(&session)
	return &session, err
}

func (store *RedisSessionStore) Delete(id string) (ok bool) {
	return
}

func (store *RedisSessionStore) Refresh(id string) (session *Session, err error) {
	err = errors.New("Unimplemented")
	return
}
