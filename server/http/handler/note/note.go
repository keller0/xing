package note

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type noteReq struct {
	Content string `json:"content"`
}

func NewNotes(c echo.Context) error {
	req := new(noteReq)
	if err := c.Bind(req); err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, "bad request")
	}
	log.Debug(req)

	ret := struct {
		Token string
	}{
		"asd",
	}

	return c.JSON(http.StatusOK, ret)
}
