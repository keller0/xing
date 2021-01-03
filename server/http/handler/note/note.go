package note

import (
	"github.com/google/uuid"
	"github.com/keller0/xing/server/storage"
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
		return c.String(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	if len(req.Content) <= 0 || len(req.Content) > 1024 {
		return c.String(http.StatusNotAcceptable, http.StatusText(http.StatusNotAcceptable))
	}

	user := c.Get(storage.CtxUserKey).(storage.CtxUserInfo)
	id, _ := uuid.NewRandom()

	notes := storage.Notes{
		Id:      id.String(),
		Content: req.Content,
		Uid:     user.Id,
	}

	err := storage.AddNotes(notes)
	if err != nil {
		log.Error(err)
		return c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return c.String(http.StatusCreated, notes.Id)
}

func GetNotes(c echo.Context) error {
	user := c.Get(storage.CtxUserKey).(storage.CtxUserInfo)
	notes, err := storage.GetNotesByUid(user.Id)
	if err != nil {
		log.Error(err)
		return c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return c.JSON(http.StatusOK, notes)

}

func GetNoteById(c echo.Context) error {
	id := c.Param("id")
	if len(id) <= 0 {
		return c.String(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	user := c.Get(storage.CtxUserKey).(storage.CtxUserInfo)
	notes, err := storage.GetNotesById(id, user.Id)
	if err != nil {
		log.Error(err)
		return c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return c.JSON(http.StatusOK, notes)

}
