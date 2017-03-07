package entity

type (
	UserID    int64
	UserToken string
)

type User struct {
	ID    UserID
	Token UserToken
}

type UserRepository interface {
}
