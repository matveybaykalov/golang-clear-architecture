package http

import (
	"clear-arch/internal/config"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *HttpController) downloadBmstuLib(ctx echo.Context) error {
	c.log.Info("/bmstu - get request")

	err := c.uc.DownloadBmstuLib(ctx.Request().Context())

	if err != nil {
		c.log.Error(err)
		return ctx.String(http.StatusOK, "Error accure, check logs")
	}

	return ctx.String(http.StatusOK, "Done")
}

func (c *HttpController) downloadHseLib(ctx echo.Context) error {
	c.log.Info("/hse - get request")

	err := c.uc.DownloadHseLib(ctx.Request().Context())

	if err != nil {
		c.log.Error(err)
		return ctx.String(http.StatusOK, "Error accure, check logs")
	}

	return ctx.String(http.StatusOK, "Done")
}

func (c *HttpController) downloadMsuLib(ctx echo.Context) error {
	c.log.Info("/msu - get request")

	err := c.uc.DownloadMsuLib(ctx.Request().Context())

	if err != nil {
		c.log.Error(err)
		return ctx.String(http.StatusOK, "Error accure, check logs")
	}

	return ctx.String(http.StatusOK, "Done")
}

func (c *HttpController) downloadWithTypeLib(ctx echo.Context) error {
	c.log.Info("/with-type - get request")

	t := ctx.QueryParam("type")

	c.log.Info("get type: " + t)

	err := c.uc.DownloadLibWithType(ctx.Request().Context(), t)

	if err != nil {
		c.log.Error(err)
		return ctx.String(http.StatusOK, "Error accure, check logs")
	}

	return ctx.String(http.StatusOK, "Done")
}

func (c *HttpController) Start(cf config.HTTPConfig) {
	c.e.GET("/bmstu", c.downloadBmstuLib)
	c.e.GET("/msu", c.downloadMsuLib)
	c.e.GET("/hse", c.downloadHseLib)
	c.e.GET("/with-type", c.downloadWithTypeLib)

	c.e.Start(fmt.Sprintf("0.0.0.0:%s", cf.Port))
}
