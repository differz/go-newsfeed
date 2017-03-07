package main

import (
	"fmt"
	"os"

	"github.com/VitaliiHurin/go-newsfeed/api"
	"github.com/VitaliiHurin/go-newsfeed/config"
	"github.com/VitaliiHurin/go-newsfeed/server"
	"github.com/VitaliiHurin/go-newsfeed/server/gin"
)

func main() {
	config.ServerParams()
	config.Load()

	if *config.ServerHTTPAddr == "" {
		fmt.Println("ERR - HTTP port is not defined.")
		os.Exit(1)
	}

	var mode server.Mode
	switch *config.ServerMode {
	case "release":
		mode = server.ModeRelease
	default:
		mode = server.ModeDebug
	}

	a := &api.API{}
	gin.New(mode, a).Run(*config.ServerHTTPAddr)
}
