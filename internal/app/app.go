package app

import (
	"clear-arch/internal/config/cli"
	httpC "clear-arch/internal/controller/http"
	bmsturepo "clear-arch/internal/repository/bmstuRepo"
	clickhouserepo "clear-arch/internal/repository/clickhouseRepo"
	msurepo "clear-arch/internal/repository/msuRepo"
	bmstuformat "clear-arch/internal/service/BMSTUformat"
	book "clear-arch/internal/service/Book"
	msuformat "clear-arch/internal/service/MSUformat"
	"clear-arch/internal/usecase"
	"os"
	"time"

	"fmt"

	"dev.solidwall.io/solidlab/NGPlatform/ost/go-clickhouse.git/ch"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

func Run() {
	cl := cli.ConfigLoader{}
	config := cl.Load()

	logger := &logrus.Logger{
		Out:   os.Stdout,
		Level: logrus.DebugLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %time% - %msg%\n",
		},
	}

	e := echo.New()

	chConnection := ch.Connect(
		ch.WithDSN(config.ClickHouseConfig.DSN),
		ch.WithReadTimeout(120*time.Second),
		ch.WithWriteTimeout(120*time.Second),
	)
	//chConnection.AddQueryHook(chdebug.NewQueryHook(
	//	chdebug.WithVerbose(true),
	//	chdebug.FromEnv("CHDEBUG"),
	//))

	//sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(config.PostgresConfig.DSN)))
	//pgConnection := bun.NewDB(sqldb, pgdialect.New())

	clickHouseRepo := clickhouserepo.New(chConnection, logger)
	//postgresRepo := postgresrepo.New(pgConnection, logger)
	bmstuRepo := bmsturepo.New(logger)
	msuRepo := msurepo.New(logger)

	bmstuParser := bmstuformat.New(bmstuRepo, logger)
	msuParser := msuformat.New(msuRepo, logger)
	hseParser := &msuformat.MsuFormatService{}
	bookService := book.New(clickHouseRepo, logger)
	//bookService := book.New(postgresRepo, logger)

	usecase := usecase.New(msuParser, bmstuParser, hseParser, bookService, logger)

	httpController := httpC.New(e, usecase, logger)

	httpController.Start(config.HttpConfig)

	fmt.Println(config)
}
