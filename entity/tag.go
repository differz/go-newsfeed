package entity

type (
	TagID   int64
	TagName string
)

type Tag struct {
	ID   TagID
	Name TagName
}

type TagRepository interface {
	GetByUser(uid UserID) ([]*Tag, error)
	GetByName(name TagName) (*Tag, error)
	Store(tag *Tag) error
}
