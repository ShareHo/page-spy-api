package handler

import (
	"net/http"
	"log"

	"github.com/ShareHo/page-spy-api/config"
	"github.com/HuolalaTech/page-spy-api/container"
	"github.com/labstack/echo/v4"
)

var srv *echo.Echo

func init() {
	container := container.Container()
	err := container.Provide(func() *config.StaticConfig {
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	// serve.Run()

	errR := container.Invoke(func(e *echo.Echo, config *config.Config, staticConfig *config.StaticConfig) {
		// if staticConfig != nil {
		// 	hash := staticConfig.GitHash
		// 	version := staticConfig.Version
		// 	if hash == "" {
		// 		hash = "local"
		// 	}
		// 	if version == "" {
		// 		version = "local"
		// 	}
		// 	log.Infof("server info: %s@%s", version, hash)
		// }

		// for _, ip := range util.GetLocalIPList() {
		// 	log.Infof("LAN address http://%s:%s", ip, config.Port)
		// }

		// log.Infof("Local address http://localhost:%s", config.Port)
		// e.Logger.Fatal(e.Start(":" + config.Port))

		srv = e
	})

	if errR != nil {
		log.Fatal(errR)
	}
}

func Handler(w http.ResponseWriter,r *http.Request) {
	srv.ServeHTTP(w,r)
 }

