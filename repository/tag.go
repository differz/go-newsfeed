package repository

import (
	"errors"

	"github.com/VitaliiHurin/go-newsfeed/entity"
	"upper.io/db.v3/lib/sqlbuilder"
)

type tagTable struct {
	ID   int64  `db:"id,omitempty"`
	Name string `db:"name"`
}

func assembleTag(t *tagTable) *entity.Tag {
	return &entity.Tag{
		ID:   entity.TagID(t.ID),
		Name: entity.TagName(t.Name),
	}
}

func newTagTable(r *entity.Tag) *tagTable {
	return &tagTable{
		ID:   int64(r.ID),
		Name: string(r.Name),
	}
}

type tagRepository struct {
	DB sqlbuilder.Database
}

func NewTagRepository(DB sqlbuilder.Database) entity.TagRepository {
	return &tagRepository{
		DB: DB,
	}
}

func (r *tagRepository) GetByUser(uid int64) ([]*entity.Tag, error) {
	if uid <= 0 {
		return nil, errors.New("Invalid argument")
	}
	q := r.DB.Select("t.id", "t.name").
		From("tag AS t", "user_tag_relation AS ut").
		Where("t.id = ut.tagID and ut.userID = ?", uid)
	var rows []tagTable
	err := q.All(&rows)
	if err != nil {
		return nil, err
	}
	var tags []*entity.Tag
	for _, v := range rows {
		tags = append(tags, assembleTag(&v))
	}
	if tags == nil {
		tags = []*entity.Tag{}
	}
	return tags, nil
}

func (r *tagRepository) GetByName(name string) (*entity.Tag, error) {
	res := r.DB.Collection("tag").Find("name", name)
	var t tagTable
	err := res.One(&t)
	if err != nil {
		return nil, err
	}
	return assembleTag(&t), nil
}

func (r *tagRepository) StoreTag(tag *entity.Tag) error {
	id, err := r.DB.Collection("tag").Insert(newTagTable(tag))
	if err != nil {
		return err
	}
	tag.ID = entity.TagID(id.(int64))
	return nil
}
