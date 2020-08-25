package account

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/keller0/xing/server/storage"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

var JWTSigningKey = "tom go"

type loginReq struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func Login(c echo.Context) error {
	req := new(loginReq)
	if err := c.Bind(req); err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, "bad request")
	}

	if len(req.Name) < 3 || len(req.Pass) < 3 {
		log.Info("name pass too short")
		return c.String(http.StatusBadRequest, "bad request")
	}

	// check user pass
	var u storage.User
	u.Name = req.Name
	u.Pass = req.Pass

	if !u.CheckUserAuth() {
		log.Info("wrong password")
		return c.String(http.StatusUnauthorized, "unauthorized")
	}

	token, et := GenUserToken(req.Name)
	if et != nil {
		log.Error(et)
		return c.String(http.StatusInternalServerError, "gen token failed")
	}

	ret := struct {
		Token string
	}{
		token,
	}

	return c.JSON(http.StatusOK, ret)
}

type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func GenUserToken(userName string) (string, error) {

	expTime := time.Now().Add(time.Hour * 24 * 30).Unix()

	claims := JwtClaims{
		Name: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(JWTSigningKey))
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
