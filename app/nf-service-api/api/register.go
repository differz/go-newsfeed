package api

import (
	"github.com/VitaliiHurin/go-newsfeed/entity"
)

func (a *API) Register(email string, password string) (*entity.User, error) {
	u, err := a.users.GetByEmail(entity.UserEmail(email))
	if err != nil && err != ErrNotFound {
		return nil, err
	}
	if u != nil {
		return nil, ErrUserAlreadyExist
	}
	salt, err := a.SecurityManager.GenerateSalt()
	if err != nil {
		return nil, err
	}
	pass, err := a.SecurityManager.GetPasswordHash(password, salt)
	if err != nil {
		return nil, err
	}
	token, err := a.SecurityManager.GenerateNewToken(email)
	if err != nil {
		return nil, err
	}
	u = &entity.User{
		Email: entity.UserEmail(email),
		Salt: entity.UserSalt(salt),
		Password: entity.UserPassword(pass),
		Token: entity.UserToken(token),
	}
	err = a.users.Store(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}