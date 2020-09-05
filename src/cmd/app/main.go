package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"gtihub.com/gin-websocket/src/service"
	"log"
	"net/http"
)

var (
	configFile = flag.String("config", "config.json", "specify config file")
	debug      = flag.Bool("debug", true, "is in debug env or not")
	root       = flag.String("root", "./", "project root")
)

func main() {
	flag.Parse()

	if !*debug {
		gin.SetMode(gin.ReleaseMode)
	}

	conf := service.ConfigFromFile(*configFile)
	srv := service.FromConfig(conf, context.Background())

	router := gin.Default()

	router.
		Use(service.MakeCORSMiddleware()).
		Static("/static", *root+"/static").
		GET("/ws", srv.Handler).
		GET("/", func(c *gin.Context) {
			c.Redirect(http.StatusPermanentRedirect, "/static/index.html")
		})

	serveAddr := service.ServeAddrFromConfig(conf)

	log.Printf("serving on %s", serveAddr)

	server := http.Server{
		Addr:    serveAddr,
		Handler: router,
	}

	defer func() {
		srv.Shutdown()
	}()

	service.NoError(server.ListenAndServe())
}
