package po

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/weblogs/pkg/conf"
)

type Model struct {
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type BlogPost struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	PostID      int64              `json:"post_id" bson:"post_id"`
	Title       string             `json:"title" bson:"title"`
	AuthorID    int64              `json:"author_id" bson:"author_id"`
	Content     string             `json:"content" bson:"content"`
	Description string             `json:"description" bson:"description"`
	Category    string             `json:"category" bson:"category"`
	Tags        []string           `json:"tags" bson:"tags"`
	Status      int64              `json:"status" bson:"status"`
	PostTime    time.Duration      `json:"post_time" bson:"post_time"`
	Model
}

func (b *BlogPost) CollectionName() string {
	return conf.BlogCollectionName
}

var schema = bson.M{
	"$jsonSchema": bson.M{
		"bsonType": "object",
		"required": []string{"name", "color"},
		"properties": bson.M{
			"name": bson.M{"bsonType": "string"},
			"color": bson.M{
				"bsonType": "string",
				"enum":     bson.A{"Red", "Green", "Blue"},
			},
		},
	},
}
