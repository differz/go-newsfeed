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
	GetByUser(uid int64) ([]*Tag, error)
	GetByName(name string) (*Tag, error)
	StoreTag(tag *Tag) error
}
