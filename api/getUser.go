package api

import "github.com/VitaliiHurin/go-newsfeed/entity"

func (a *API) GetUser(token string) (*entity.User, error) {
	return &entity.User{}, nil
}