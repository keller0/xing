package main

import (
	"flag"
	"github.com/keller0/xing/server/config"
	"github.com/keller0/xing/server/http/router"
	"github.com/keller0/xing/server/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

var (
	serverHost  string
	serverPort  string
	dbPath      string
	allowOrigin []string
	loggerMd    echo.MiddlewareFunc
)

func init() {
	var configFile = flag.String("conf", "./config.yaml", "config file path")
	flag.StringVar(&serverHost, "host", "127.0.0.1", "listen address")
	flag.StringVar(&serverPort, "port", "6666", "listen port")
	flag.StringVar(&dbPath, "db", "", "database path")
	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())

	var configs config.AllConf
	content, err := ioutil.ReadFile(*configFile)
	if err != nil {
		panic("read config file failed")
	}
	if err := yaml.Unmarshal(content, &configs); err != nil {
		panic("unmarshal config content failed")
	}
	lv, err := log.ParseLevel("info")
	if err != nil {
		panic(err)
	}
	log.SetLevel(lv)
	f, err := os.OpenFile(configs.App.LogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)

	loggerMd = middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: f,
	})

	allowOrigin = strings.Split(configs.App.AllowOrigin, " ")
	if len(allowOrigin) <= 0 {
		panic("empty allow_origins")
	}
	log.Info(allowOrigin)

	storage.InitSqlite(dbPath)

}

func main() {

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     allowOrigin,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "token"},
	}))

	// Middleware
	e.Use(loggerMd)
	e.Use(middleware.Recover())

	// Routes
	router.Register(e)

	// Start server
	e.Logger.Fatal(e.Start(serverHost + ":" + serverPort))
}
