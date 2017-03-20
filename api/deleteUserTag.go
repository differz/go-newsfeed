package api

import "github.com/VitaliiHurin/go-newsfeed/entity"

func (a *API) DeleteUserTag(user *entity.User, tag string) error {
	tagEntity, err := a.tags.GetByName(entity.TagName(tag))
	if err != nil {
		return err
	}
	return a.userTags.RemoveTagFromUser(user, tagEntity)
}
