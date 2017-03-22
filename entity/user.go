package entity

type (
	UserID       int64
	UserToken    string
	UserEmail    string
	UserPassword string
	UserSalt     string
)

type User struct {
	ID       UserID
	Token    UserToken
	Email    UserEmail
	Password UserPassword
	Salt     UserSalt
}

type UserRepository interface {
	GetByToken(token UserToken) (*User, error)
	GetByEmail(email UserEmail) (*User, error)
	Store(user *User) error
}
