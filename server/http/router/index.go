package router

import (
	"github.com/dgrijalva/jwt-go"
	h "github.com/keller0/xing/server/http/handler"
	"github.com/keller0/xing/server/http/handler/account"
	"github.com/keller0/xing/server/http/handler/note"
	"github.com/keller0/xing/server/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Register(e *echo.Echo) {
	e.Use(middleware.BodyLimit("1M"))
	e.GET("/", h.Hello)
	e.POST("/login", account.Login)
	e.POST("/notes", note.NewNotes, jwtMd)
}

/*

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()

			// Based on content length
			if req.ContentLength > config.limit {
				return echo.ErrStatusRequestEntityTooLarge
			}

			// Based on content read
			r := pool.Get().(*limitedReader)
			r.Reset(req.Body, c)
			defer pool.Put(r)
			req.Body = r

			return next(c)
		}
	}
*/
func jwtMd(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("x-jwt")

		//fmt.Println(tokenString)
		claims := &storage.JwtClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return storage.JWTSigningKey, nil
		})
		if err != nil {
			log.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		if !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		return next(c)

	}
}
