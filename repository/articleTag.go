package repository

import (
	"github.com/VitaliiHurin/go-newsfeed/entity"
	"upper.io/db.v3/lib/sqlbuilder"
)

type articleTagTable struct {
	ArticleID int64 `db:"articleID,omitempty"`
	TagID     int64 `db:"tagID,omitempty"`
}

func assembleArticleTag(t *articleTagTable) *entity.ArticleTag {
	return &entity.ArticleTag{
		ArticleID: entity.ArticleID(t.ArticleID),
		TagID:     entity.TagID(t.TagID),
	}
}

func newArticleTagTable(r *entity.ArticleTag) *articleTagTable {
	return &articleTagTable{
		ArticleID: int64(r.ArticleID),
		TagID:     int64(r.TagID),
	}
}

type articleTagRepository struct {
	DB sqlbuilder.Database
}

func NewArticleTagRepository(DB sqlbuilder.Database) entity.ArticleTagRepository {
	return &articleTagRepository{
		DB: DB,
	}
}
