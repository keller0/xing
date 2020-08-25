package router

import (
	h "github.com/keller0/xing/server/http/handler"
	"github.com/keller0/xing/server/http/handler/account"
	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo) {
	e.GET("/", h.Hello)
	e.POST("/login", account.Login)
}
