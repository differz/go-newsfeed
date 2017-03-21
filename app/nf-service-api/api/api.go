package api

import "github.com/VitaliiHurin/go-newsfeed/entity"

type API struct {
	articles    entity.ArticleRepository
	articleTags entity.ArticleTagRepository
	services    entity.ServiceRepository
	tags        entity.TagRepository
	users       entity.UserRepository
	userTags    entity.UserTagRepository
}

func New(
	articles entity.ArticleRepository,
	articleTags entity.ArticleTagRepository,
	services entity.ServiceRepository,
	tags entity.TagRepository,
	users entity.UserRepository,
	userTags entity.UserTagRepository,
) *API {
	return &API{
		articles:    articles,
		articleTags: articleTags,
		services:    services,
		tags:        tags,
		users:       users,
		userTags:    userTags,
	}
}
