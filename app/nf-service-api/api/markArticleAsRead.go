package api

import "github.com/VitaliiHurin/go-newsfeed/entity"

func (a *API) MarkArticleAsRead(id entity.ArticleID) error {
	return a.articles.ChangeIsRead(id, entity.ArticleIsRead(true))
}
