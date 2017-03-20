package repository

import (
	"github.com/VitaliiHurin/go-newsfeed/entity"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

type userTagTable struct {
	UserID int64 `db:"userID,omitempty"`
	TagID  int64 `db:"tagID,omitempty"`
}

func assembleUserTag(t *userTagTable) *entity.UserTag {
	return &entity.UserTag{
		UserID: entity.UserID(t.UserID),
		TagID:  entity.TagID(t.TagID),
	}
}

func newUserTagTable(r *entity.UserTag) *userTagTable {
	return &userTagTable{
		UserID: int64(r.UserID),
		TagID:  int64(r.TagID),
	}
}

type userTagRepository struct {
	DB *sqlbuilder.Database
}

func NewUserTagRepository(DB *sqlbuilder.Database) entity.UserTagRepository {
	return &userTagRepository{
		DB: DB,
	}
}

func (r *userTagRepository) IsUserHasTag(user *entity.User, tag *entity.Tag) (bool, error) {
	res := (*r.DB).Collection("user_tag_relation").Find(db.Cond{
		"userID": user.ID,
		"tagID":  tag.ID,
	})
	c, err := res.Count()
	if err != nil {
		return false, err
	}
	if c == 0 {
		return false, nil
	}
	return true, nil
}

func (r *userTagRepository) Store(userTag *entity.UserTag) error {
	_, err := (*r.DB).Collection("user_tag_relation").Insert(newUserTagTable(userTag))
	if err != nil {
		return err
	}
	return nil
}

func (r *userTagRepository) RemoveTagFromUser(user *entity.User, tag *entity.Tag) error {
	q := (*r.DB).DeleteFrom("user_tag_relation").Where(db.Cond{
		"userId": user.ID,
		"tagID":  tag.ID,
	})
	_, err := q.Exec()
	return err
}
