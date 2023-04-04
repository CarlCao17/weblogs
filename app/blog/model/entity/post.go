package entity

import (
	"time"

	"github.com/weblogs/pkg/conf"
)

type Model struct {
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func NewModel() Model {
	now := time.Now().UTC()
	return Model{
		CreatedAt: now,
		UpdatedAt: now,
	}
}

type BlogPost struct {
	PostID      int64         `json:"post_id" bson:"post_id"`
	Title       string        `json:"title" bson:"title"`
	AuthorID    int64         `json:"author_id" bson:"author_id"`
	Content     string        `json:"content" bson:"content"`
	Description string        `json:"description" bson:"description"`
	Category    string        `json:"category" bson:"category"`
	Tags        []string      `json:"tags" bson:"tags"`
	Status      int64         `json:"status" bson:"status"`
	PostTime    time.Duration `json:"post_time" bson:"post_time"`
	Model
}

const (
	PostCreated   = 1
	PostPublished = 2
	PostDeleted   = 3
)

func (b *BlogPost) CollectionName() string {
	return conf.BlogCollectionName
}
