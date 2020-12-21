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
	"runtime"
	"strings"
)

var (
	serverHost  string
	serverPort  string
	dbPath      string
	allowOrigin []string
)

func init() {
	var configFile = flag.String("conf", "./config.yaml", "config file path")
	flag.StringVar(&serverHost, "host", "127.0.0.1", "listen address")
	flag.StringVar(&serverPort, "port", "1323", "listen port")
	flag.StringVar(&dbPath, "db", "x.db", "database path")
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
	lv, err := log.ParseLevel(configs.App.LogLevel)
	if err != nil {
		panic(err)
	}
	log.SetLevel(lv)

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
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "token"},
	}))

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	router.Register(e)

	// Start server
	e.Logger.Fatal(e.Start(serverHost + ":" + serverPort))
}
