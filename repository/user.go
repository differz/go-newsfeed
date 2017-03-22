package repository

import (
	"errors"

	"github.com/VitaliiHurin/go-newsfeed/entity"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
	"github.com/VitaliiHurin/go-newsfeed/app/nf-service-api/api"
)

type userTable struct {
	ID       int64  `db:"id,omitempty"`
	Token    string `db:"token"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Salt     string `db:"salt"`
}

func assembleUser(t *userTable) *entity.User {
	return &entity.User{
		ID:       entity.UserID(t.ID),
		Token:    entity.UserToken(t.Token),
		Email:    entity.UserEmail(t.Email),
		Password: entity.UserPassword(t.Password),
		Salt:     entity.UserSalt(t.Salt),
	}
}

func newUserTable(r *entity.User) *userTable {
	return &userTable{
		ID:    int64(r.ID),
		Token: string(r.Token),
		Email: string(r.Email),
		Password: string(r.Password),
		Salt: string(r.Salt),
	}
}

type userRepository struct {
	DB sqlbuilder.Database
}

func NewUserRepository(DB sqlbuilder.Database) entity.UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (r *userRepository) GetByToken(token entity.UserToken) (*entity.User, error) {
	res := r.DB.Collection("user").Find(db.Cond{
		"token": token,
	})
	c, err := res.Count()
	if err != nil {
		return nil, err
	}
	if c > 0 {
		var user userTable
		err = res.One(&user)
		if err != nil {
			return nil, err
		}
		return assembleUser(&user), nil
	}
	return nil, errors.New("User not exist")
}

func (r *userRepository) GetByEmail(email entity.UserEmail) (*entity.User, error) {
	res := r.DB.Collection("user").Find(db.Cond{
		"email": email,
	})
	c, err := res.Count()
	if err != nil {
		return nil, err
	}
	if c > 0 {
		var user userTable
		err = res.One(&user)
		if err != nil {
			return nil, err
		}
		return assembleUser(&user), nil
	}
	return nil, api.ErrNotFound
}

func (r *userRepository) Store(user *entity.User) error {
	id, err := r.DB.Collection("user").Insert(newUserTable(user))
	if err != nil {
		return err
	}
	user.ID = entity.UserID(id.(int64))
	return nil
}
