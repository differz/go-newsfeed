package api

import (
	"github.com/VitaliiHurin/go-newsfeed/entity"
	"upper.io/db.v3"
)

func (a *API) AddUserTag(user *entity.User, tag string) error {
	tagEntity, err := a.tags.GetByName(tag)
	if err != nil {
		if err == db.ErrNoMoreRows {
			tagEntity = &entity.Tag{
				Name: entity.TagName(tag),
			}
			err := a.tags.StoreTag(tagEntity)
			if err != nil {
				return err
			}
		}
		return err
	}
	ok, err := a.userTags.IsUserHasTag(user, tagEntity)
	if err != nil {
		return err
	}
	if !ok {
		err = a.userTags.Store(&entity.UserTag{
			UserID: user.ID,
			TagID: tagEntity.ID,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
