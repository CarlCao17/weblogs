package convert

import (
	"context"
	"time"

	"github.com/weblogs/app/blog/model/entity"
	"github.com/weblogs/app/blog/model/po"
)

type postEntity2POConverter struct{}

var PostEntity2POConverter = &postEntity2POConverter{}

func (cvt *postEntity2POConverter) ToPO(ctx context.Context, e *entity.BlogPost) (*po.BlogPost, error) {
	return &po.BlogPost{
		PostID:      e.PostID,
		Title:       e.Title,
		AuthorID:    e.AuthorID,
		Content:     e.Content,
		Description: e.Description,
		Category:    e.Category,
		Tags:        e.Tags,
		Status:      e.Status,
		PostTime:    e.PostTime,
		Model: po.Model{
			CreatedAt: e.CreatedAt,
			UpdatedAt: e.UpdatedAt,
			IsDel:     e.IsDel,
		},
	}, nil
}

func NewModel() po.Model {
	now := time.Now().UTC()
	return po.Model{
		CreatedAt: now,
		UpdatedAt: now,
		IsDel:     false,
	}
}
