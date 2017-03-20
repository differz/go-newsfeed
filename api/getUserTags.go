package api

import "github.com/VitaliiHurin/go-newsfeed/entity"

func (a *API) GetUserTags(user *entity.User) ([]*entity.Tag, error) {
	return a.tags.GetByUser(user.ID)
}
