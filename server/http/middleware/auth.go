package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/keller0/xing/server/storage"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// check jwt token

func JwtMd(next echo.HandlerFunc) echo.HandlerFunc {
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

		cu := storage.CtxUserInfo{Id: claims.Id, Name: claims.Name}
		c.Set(storage.CtxUserKey, cu)

		return next(c)

	}
}
