package main

import (
	"fmt"
	"os"

	"github.com/VitaliiHurin/go-newsfeed/api"
	"github.com/VitaliiHurin/go-newsfeed/config"
	"github.com/VitaliiHurin/go-newsfeed/repository"
	"github.com/VitaliiHurin/go-newsfeed/server"
	"github.com/VitaliiHurin/go-newsfeed/server/gin"
)

func main() {
	config.ServerParams()
	config.DBParams()

	if config.ServerHTTPPort == "" {
		fmt.Println("ERR - HTTP port is not defined.")
		os.Exit(1)
	}

	var mode server.Mode
	switch config.ServerMode {
	case "release":
		mode = server.ModeRelease
	default:
		mode = server.ModeDebug
	}

	articles := repository.NewArticleRepository(config.DB)
	users := repository.NewUserRepository(config.DB)
	tags := repository.NewTagRepository(config.DB)
	services := repository.NewServiceRepository(config.DB)
	userTags := repository.NewUserTagRepository(config.DB)
	articleTags := repository.NewArticleTagRepository(config.DB)

	a := api.New(articles, articleTags, services, tags, users, userTags)
	gin.New(mode, a).Run(":"+config.ServerHTTPPort)
}
