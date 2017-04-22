package rss_reader

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/VitaliiHurin/go-newsfeed/entity"
	"github.com/mmcdole/gofeed"
)

type Reader struct {
	tagRep     entity.TagRepository
	articleRep entity.ArticleRepository
	atRep      entity.ArticleTagRepository
}


func NewReader(tag entity.TagRepository, article entity.ArticleRepository, at entity.ArticleTagRepository) *Reader {
	return &Reader{
		tagRep: tag,
		articleRep: article,
		atRep: at,
	}
}

func (r *Reader) ProcessService(service *entity.Service) (int, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(string(service.Host))
	if err != nil {
		return 0, err
	}
	if len(feed.Items) == 0 {
		return 0, errors.New("No items")
	}

	newArticles := []*entity.Article{}

	for _, item := range feed.Items {
		a, err := r.articleRep.FindByUrlAndSource(item.Link, service.ID)
		if err != nil {
			return 0, err
		}
		if a == nil {
			a, tags := r.newArticleFromFeed(item)
			a, err = r.articleRep.Store(a)
			if err != nil {
				return 0, err
			}
			for _, t := range tags {
				tag, err := r.tagRep.GetByName(entity.TagName(t))
				if err != nil {
					fmt.Printf("Tag '%s': %v\n", t, err)
				}
				if tag == nil {
					tag = &entity.Tag{
						Name: entity.TagName(t),
					}
					tag, err = r.tagRep.Store(tag)
					if err != nil {
						log.Panic(err)
						return 0, err
					}
				}
				err = r.articleRep.AddTag(a, tag)
				if err != nil {
					return 0, err
				}
			}
			newArticles = append(newArticles, a)
		}
	}
	return len(newArticles), nil
}

func (r *Reader) newArticleFromFeed(item *gofeed.Item) (*entity.Article, []string) {
	a := &entity.Article{
		Title: entity.ArticleTitle(item.Title),
		Description: entity.ArticleDescription(item.Description),
		URL: entity.ArticleURL(item.Link),
		DateCreated: entity.ArticleDateCreated(*item.PublishedParsed),
		DateIndexed: entity.ArticleDateIndexed(time.Now()),
		IsRead: entity.ArticleIsRead(false),
	}
	return a, item.Categories
}