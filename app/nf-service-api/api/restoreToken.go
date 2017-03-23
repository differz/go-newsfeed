package api

import "github.com/VitaliiHurin/go-newsfeed/entity"

func (a *API) RestoreToken(user *entity.User) error {
	token, err := a.SecurityManager.GenerateNewToken(string(user.Email))
	if err != nil {
		return err
	}
	user.Token = entity.UserToken(token)
	err = a.users.Store(user)
	if err != nil {
		return err
	}
	return nil
}
