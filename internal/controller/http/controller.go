package http

import (
	"clear-arch/internal/interfaces"
	"context"

	"github.com/labstack/echo/v4"
)

type httpUsecase interface {
	DownloadBmstuLib(context.Context) error
	DownloadFilterBmstuLib(context.Context) error
	DownloadHseLib(context.Context) error
	DownloadLibWithType(context.Context, string) error
	DownloadMsuLib(context.Context) error
}

type HttpController struct {
	e   *echo.Echo
	uc  httpUsecase
	log interfaces.Logger
}

func New(echo *echo.Echo, uc httpUsecase, l interfaces.Logger) *HttpController {
	return &HttpController{
		e:   echo,
		uc:  uc,
		log: l,
	}
}
