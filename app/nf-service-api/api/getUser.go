package api

import "github.com/VitaliiHurin/go-newsfeed/entity"

func (a *API) GetUser(email string) (*entity.User, error) {
	return a.users.GetByEmail(entity.UserEmail(email))
}