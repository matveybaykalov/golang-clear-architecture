package main

import (
	"clear-arch/internal/app"
	"clear-arch/internal/config/cli"
)

func main() {
	app.Run(cli.ConfigLoader{})
}
