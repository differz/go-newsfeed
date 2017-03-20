package repository

import (
	"errors"

	"github.com/VitaliiHurin/go-newsfeed/entity"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

type userTable struct {
	ID    int64     `db:"id,omitempty"`
	Token string    `db:"token"`
}

func assembleUser(t *userTable) *entity.User {
	return &entity.User{
		ID:    entity.UserID(t.ID),
		Token: entity.UserToken(t.Token),
	}
}

func newUserTable(r *entity.User) *userTable {
	return &userTable{
		ID:    int64(r.ID),
		Token: string(r.Token),
	}
}

type userRepository struct {
	DB *sqlbuilder.Database
}

func NewUserRepository(DB *sqlbuilder.Database) entity.UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (r *userRepository) GetByToken(token entity.UserToken) (*entity.User, error) {
	res := (*r.DB).Collection("user").Find(db.Cond{
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
