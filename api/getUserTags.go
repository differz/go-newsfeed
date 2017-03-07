package api

import "github.com/VitaliiHurin/go-newsfeed/entity"

func (a *API) GetUserTags(user *entity.User) ([]*entity.UserTag, error) {
	return []*entity.UserTag{}, nil
}
