package cli

import (
	"clear-arch/internal/config"
	"flag"
	"log"
	"os"
)

type ConfigLoader struct{}

var isHelp = flag.Bool("help", false, "show this help message")
var httpPort = flag.String("httpPort", "5678", "set up port for http handler")
var chDSN = flag.String("chDSN", "", "set up DSN for clickhouse")
var pgDSN = flag.String("pgDSN", "", "set up DSN for postgres")

func (ConfigLoader) Load() config.Config {
	flag.Parse()
	if *isHelp {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *chDSN == "" {
		log.Fatal("please, set ClickHouse DSN  (use flag -chDSN)")
	}

	if *pgDSN == "" {
		log.Fatal("please, set Postgres DSN  (use flag -pgDSN)")
	}

	return config.Config{
		ClickHouseConfig: config.ClickHouseConfig{
			DSN: *chDSN,
		},
		PostgresConfig: config.PostgresConfig{
			DSN: *pgDSN,
		},
		HttpConfig: config.HTTPConfig{
			Port: *httpPort,
		},
	}
}
