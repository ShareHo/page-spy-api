package serve

import (
	"github.com/HuolalaTech/page-spy-api/config"
	"github.com/HuolalaTech/page-spy-api/container"
	"github.com/HuolalaTech/page-spy-api/util"
	"github.com/labstack/echo/v4"
)

func Run() {
	err := container.Container().Invoke(func(e *echo.Echo, config *config.Config) {

		for _, ip := range util.GetLocalIPList() {
			log.Infof("LAN address http://%s:%s", ip, config.Port)
		}

		log.Infof("Local address http://localhost:%s", config.Port)
		e.Logger.Fatal(e.Start(":" + config.Port))
	})

	if err != nil {
		log.Fatal(err)
	}
}
