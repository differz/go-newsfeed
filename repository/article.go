package repository

import (
	"errors"
	"time"

	"github.com/VitaliiHurin/go-newsfeed/entity"
	"upper.io/db.v3/lib/sqlbuilder"
	"fmt"
	"log"
)

type articleTable struct {
	ID          int64     `db:"id,omitempty"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	URL         string    `db:"url"`
	DateCreated int64     `db:"dateCreated"`
	DateIndexed int64     `db:"dateIndexed"`
	IsRead      bool      `db:"isRead"`
	ServiceID   int64     `db:"serverID"`
}

func assembleArticle(t *articleTable) *entity.Article {
	return &entity.Article{
		ID:          entity.ArticleID(t.ID),
		Title:       entity.ArticleTitle(t.Title),
		Description: entity.ArticleDescription(t.Description),
		URL:         entity.ArticleURL(t.URL),
		DateCreated: entity.ArticleDateCreated(time.Unix(t.DateCreated, 0)),
		DateIndexed: entity.ArticleDateIndexed(time.Unix(t.DateIndexed, 0)),
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
		DateCreated: time.Time(r.DateCreated).Unix(),
		DateIndexed: time.Time(r.DateIndexed).Unix(),
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
	if uid <= 0 {
		return nil, errors.New("Invalid argument")
	}
	q := r.DB.
		Select("a.id", "a.title", "a.description", "a.url", "a.dateCreated", "a.dateIndexed", "a.isRead", "a.serverID").
		From("article as a").
		Join("article_tag_relation as at").
		On("at.articleID = a.id").
		Join("user_tag_relation as ut").
		On("ut.tagID = at.tagID").
		Where("ut.userID = ?", uid).
		OrderBy("a.dateIndexed desc")
	var rows []articleTable
	err := q.All(&rows)
	if err != nil {
		return nil, err
	}
	var articles []*entity.Article
	for _, v := range rows {
		articles = append(articles, assembleArticle(&v))
	}
	if articles == nil {
		articles = []*entity.Article{}
	}
	return articles, nil
}

func (r *articleRepository) GetByTag(tid entity.TagID) ([]*entity.Article, error) {
	if tid <= 0 {
		return nil, errors.New("Invalid argument")
	}
	q := r.DB.
		Select("a.*").
		From("article a", "article_tag_relation at").
		Where("a.id = at.articleID and at.tagID = ?", tid)
	var rows []articleTable
	err := q.All(&rows)
	if err != nil {
		return nil, err
	}
	var articles []*entity.Article
	for _, v := range rows {
		articles = append(articles, assembleArticle(&v))
	}
	if articles == nil {
		articles = []*entity.Article{}
	}
	return articles, nil
}

func (r *articleRepository) Store(a *entity.Article) error {
	id, err := r.DB.Collection("article").Insert(newArticleTable(a))
	if err != nil {
		return err
	}
	a.ID = entity.ArticleID(id.(int64))
	return nil
}

func (r *articleRepository) ChangeIsRead(aid entity.ArticleID, isRead entity.ArticleIsRead) error {
	return r.DB.Collection("article").Find("id", aid).Update(map[string]interface{}{
		"isRead": isRead,
	})
}

func (r *articleRepository) GetAll(){
	var birthdays []articleTable

	err := r.DB.Collection("article").Find().All(&birthdays)
	if err != nil {
		log.Panic("res.All(): %q\n", err)
	}

	// Printing to stdout.
	for _, birthday := range birthdays {
		fmt.Printf("%s was born in %s.\n",
			birthday.Title,
			time.Unix(birthday.DateCreated, 0).Format("January 2, 2006"),
		)
	}
}
