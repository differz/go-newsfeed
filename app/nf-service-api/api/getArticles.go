package api

import "github.com/VitaliiHurin/go-newsfeed/entity"

func (a *API) GetArticles(user *entity.User) ([]*entity.Article, error) {
	return a.articles.GetByUser(user.ID)
}
