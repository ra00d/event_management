package utils

import (
	"github.com/gofiber/fiber/v2/middleware/session"
)

type SessionHelper struct {
	sess session.Session
}

func (s SessionHelper) GetItem(key string) interface{} {
	// sess, err := configs.GetSessionStore(ctx)
	return s.sess.Get(key)
}

func (s SessionHelper) SetItem(key string, value interface{}) {
	// sess, err := configs.GetSessionStore(ctx)
	s.sess.Set(key, value)
}
