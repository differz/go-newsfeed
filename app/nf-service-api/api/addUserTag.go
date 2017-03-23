package api

import (
	"github.com/VitaliiHurin/go-newsfeed/entity"
	"upper.io/db.v3"
)

func (a *API) AddUserTag(user *entity.User, tag string) error {
	tagName := entity.TagName(tag)
	t, err := a.tags.GetByName(tagName)
	if err != nil {
		if err == db.ErrNoMoreRows {
			t = &entity.Tag{
				Name: tagName,
			}
			t, err = a.tags.Store(t)
			if err != nil {
				return err
			}
		}
		return err
	}
	ok, err := a.userTags.IsUserHasTag(user, t)
	if err != nil {
		return err
	}
	if !ok {
		err = a.userTags.Store(&entity.UserTag{
			UserID: user.ID,
			TagID:  t.ID,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
