package configs

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/mysql/v2"
)

var (
	SesseionStore  *session.Store
	sessionStorage fiber.Storage
	// CsrfMiddleware fiber.Handler
)

// const HeaderName = "X-Csrf-Token"
//
// func GetSessionStoreForCsrf() *session.Store {
// 	return SesseionStore
// }
//
// func CsrfMiddleware() fiber.Handler {
// 	sessionStorage = sqlite3.New(
// 		sqlite3.Config{Database: "./storage/sessions/sessions.sqlite3", Table: "session"},
// 	)
// 	store := session.New(session.Config{
// 		Storage:        sessionStorage,
// 		CookieSecure:   true,
// 		CookieSameSite: "Strict",
// 		// CookieHTTPOnly: true,
// 	})
// 	SesseionStore = store
//
// 	ConfigDefault := csrf.Config{
// 		KeyLookup:         "header:" + HeaderName,
// 		ContextKey:        "csrf",
// 		CookieName:        "csrf_",
// 		CookieSameSite:    "Lax",
// 		CookieSessionOnly: true,
// 		CookieHTTPOnly:    true,
// 		// ErrorHandler:      defaultErrorHandler,
// 		// Extractor:         CsrfFromHeader(HeaderName),
// 		Session:           store,
// 		SessionKey:        "fiber.csrf.token",
// 		HandlerContextKey: "fiber.csrf.handler",
// 	}
// 	csrfMiddleware := csrf.New(ConfigDefault)
// 	return csrfMiddleware
// }

func SessionsInit() {
	sessionStorage = mysql.New(
		mysql.Config{
			Host:       "127.0.0.1",
			Port:       3306,
			Username:   "root",
			Database:   "events_db",
			Table:      "sessions",
			Reset:      false,
			GCInterval: 10 * time.Second,
		},
	)
	store := session.New(session.Config{
		Storage:        sessionStorage,
		CookieSecure:   true,
		CookieSameSite: "Strict",
		CookieHTTPOnly: true,
	})
	SesseionStore = store
}

func GetSessionStore(ctx *fiber.Ctx) (session.Session, error) {
	sess, err := SesseionStore.Get(ctx)
	if err != nil {
		fmt.Println(err)
		return *sess, err
	}
	return *sess, nil
}
