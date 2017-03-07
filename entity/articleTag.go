package entity

type (
	ArticleTagID int64
)

type ArticleTag struct {
	ID        ArticleTagID
	ArticleID ArticleID
	TagID     TagID
}

type ArticleTagRepository interface {
}
