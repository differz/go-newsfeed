package api

import (
	"github.com/VitaliiHurin/go-newsfeed/entity"
	"github.com/VitaliiHurin/go-newsfeed/app/nf-service-api/security"
)

type API struct {
	articles        entity.ArticleRepository
	articleTags     entity.ArticleTagRepository
	services        entity.ServiceRepository
	tags            entity.TagRepository
	users           entity.UserRepository
	userTags        entity.UserTagRepository
	SecurityManager *security.SecurityManager
}

func New(
	articles entity.ArticleRepository,
	articleTags entity.ArticleTagRepository,
	services entity.ServiceRepository,
	tags entity.TagRepository,
	users entity.UserRepository,
	userTags entity.UserTagRepository,
	securityManager *security.SecurityManager,
) *API {
	return &API{
		articles:        articles,
		articleTags:     articleTags,
		services:        services,
		tags:            tags,
		users:           users,
		userTags:        userTags,
		SecurityManager: securityManager,
	}
}
