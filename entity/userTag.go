package entity

type UserTag struct {
	UserID UserID
	TagID  TagID
}

type UserTagRepository interface {
	IsUserHasTag(user *User, tag *Tag) (bool, error)
	Store(userTag *UserTag) error
	RemoveTagFromUser(user *User, tag *Tag) error
}
