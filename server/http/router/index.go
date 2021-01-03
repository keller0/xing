package router

import (
	h "github.com/keller0/xing/server/http/handler"
	"github.com/keller0/xing/server/http/handler/account"
	"github.com/keller0/xing/server/http/handler/note"
	md "github.com/keller0/xing/server/http/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Register(e *echo.Echo) {
	e.Use(middleware.BodyLimit("1M"))
	e.GET("/", h.Hello)
	e.POST("/login", account.Login)
	e.POST("/notes", note.NewNotes, md.JwtMd)
	e.GET("/notes", note.GetNotes, md.JwtMd)
	e.GET("/notes/:id", note.GetNoteById, md.JwtMd)
}
