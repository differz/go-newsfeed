package entity

import "time"

type (
	ArticleID          int64
	ArticleTitle       string
	ArticleDescription string
	ArticleURL         string
	ArticleDateCreated time.Time
	ArticleDateIndexed time.Time
	ArticleIsRead      bool
)

type Article struct {
	ID          ArticleID
	Title       ArticleTitle
	Description ArticleDescription
	URL         ArticleURL
	DateCreated ArticleDateCreated
	DateIndexed ArticleDateIndexed
	IsRead      ArticleIsRead
	ServiceID   ServiceID
}

type ArticleRepository interface {
	GetByUser(uid UserID) ([]*Article, error)
	GetByTag(tid TagID) ([]*Article, error)
	Store(a *Article) (*Article, error)
	ChangeIsRead(aid ArticleID, isRead ArticleIsRead) error
	GetAll()
	FindById(id ArticleID) (*Article, error)
	FindByUrlAndSource(url string, sId ServiceID) (*Article, error)
	AddTag(article *Article, tag *Tag) error
}
