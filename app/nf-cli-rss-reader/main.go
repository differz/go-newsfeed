package main

import (
	"github.com/VitaliiHurin/go-newsfeed/config"
	"github.com/VitaliiHurin/go-newsfeed/repository"
	"log"
	"fmt"
	"github.com/VitaliiHurin/go-newsfeed/app/nf-cli-rss-reader/rss-reader"
)

func main(){
	config.DBParams()

	serviceRep := repository.NewServiceRepository(config.DB)
	articleRep := repository.NewArticleRepository(config.DB)
	tagRep := repository.NewTagRepository(config.DB)
	articleTagRep := repository.NewArticleTagRepository(config.DB)

	reader := rss_reader.NewReader(tagRep, articleRep, articleTagRep)

	services, err := serviceRep.GetAll()
	if err != nil {
		log.Panic(err)
	}

	for _, service := range services {
		fmt.Printf("Process: '%s'\n", service.Name)
		c, err := reader.ProcessService(service)
		if err != nil {
			fmt.Printf("%v\n",err)
			continue
		}
		fmt.Printf("\tLoaded: %d\n", c)
	}
}
