package entity

type (
	UserTagID int64
)

type UserTag struct {
	ID     UserTagID
	UserID UserID
	TagID  TagID
}

type UserTagRepository interface {
}
