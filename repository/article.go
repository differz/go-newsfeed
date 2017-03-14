package repository

import (
	"github.com/VitaliiHurin/go-newsfeed/entity"
	"upper.io/db.v3"
)

type ArticleRepository struct {
	db *db.Database
}

func (r *ArticleRepository) GetByUser(uid entity.UserID) ([]*entity.Article, error) {
	(*r.db).Collection("").
}
