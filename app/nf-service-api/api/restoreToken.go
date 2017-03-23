package api

import "github.com/VitaliiHurin/go-newsfeed/entity"

func (a *API) RestoreToken(user *entity.User, token string) error {
	u, err := a.users.GetByToken(entity.UserToken(token))
	if err != nil {
		return err
	}
	if u.ID != user.ID {
		return ErrUnauthorized
	}
	newToken, err := a.SecurityManager.GenerateNewToken(string(user.Email))
	if err != nil {
		return err
	}
	user.Token = entity.UserToken(newToken)
	err = a.users.Store(user)
	if err != nil {
		return err
	}
	return nil
}
