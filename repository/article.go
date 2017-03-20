package repository

import (
	"time"

	"github.com/VitaliiHurin/go-newsfeed/entity"
	"upper.io/db.v3/lib/sqlbuilder"
)

type articleTable struct {
	ID          int64     `db:"id,omitempty"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	URL         string    `db:"url"`
	DateCreated time.Time `db:"dateCreated"`
	DateIndexed time.Time `db:"dateIndexed"`
	IsRead      bool      `db:"isRead"`
	ServiceID   int64     `db:"serverID"`
}

func assembleArticle(t *articleTable) *entity.Article {
	return &entity.Article{
		ID:          entity.ArticleID(t.ID),
		Title:       entity.ArticleTitle(t.Title),
		Description: entity.ArticleDescription(t.Description),
		URL:         entity.ArticleURL(t.URL),
		DateCreated: entity.ArticleDateCreated(t.DateCreated),
		DateIndexed: entity.ArticleDateIndexed(t.DateIndexed),
		IsRead:      entity.ArticleIsRead(t.IsRead),
		ServiceID:   entity.ServiceID(t.ServiceID),
	}
}

func newArticleTable(r *entity.Article) *articleTable {
	return &articleTable{
		ID:          int64(r.ID),
		Title:       string(r.Title),
		Description: string(r.Description),
		URL:         string(r.URL),
		DateCreated: time.Time(r.DateCreated),
		DateIndexed: time.Time(r.DateIndexed),
		IsRead:      bool(r.IsRead),
		ServiceID:   int64(r.ServiceID),
	}
}

type articleRepository struct {
	DB sqlbuilder.Database
}

func NewArticleRepository(DB sqlbuilder.Database) entity.ArticleRepository {
	return &articleRepository{
		DB: DB,
	}
}

func (r *articleRepository) GetByUser(uid entity.UserID) ([]*entity.Article, error) {
	return nil, nil
}
