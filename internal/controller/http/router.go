package http

import (
	"clear-arch/internal/config"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *HttpController) downloadBmstuLib(ctx echo.Context) error {
	c.log.Info("accepted get connection to bmstu handler")

	err := c.uc.DownloadBmstuLib(ctx.Request().Context())

	if err != nil {
		c.log.Error(err)
		return ctx.String(http.StatusOK, "Error accure, check logs")
	}

	return ctx.String(http.StatusOK, "Done")
}

func (c *HttpController) testFunc(ctx echo.Context) error {
	c.log.Info("accepted get connection to test handler")
	return ctx.String(http.StatusOK, "Hello, world!")
}

func (c *HttpController) Start(cf config.HTTPConfig) {
	c.e.GET("/test", c.testFunc)
	c.e.GET("/bmstu", c.downloadBmstuLib)

	c.e.Start(fmt.Sprintf("0.0.0.0:%s", cf.Port))
}
