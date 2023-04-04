package convert

import (
	"github.com/weblogs/app/blog/model/entity"
	"github.com/weblogs/kitex_gen/github/com/weblogs/blog"
)

func CreatePostReq2Entity(req *blog.CreatePostRequest) (*entity.BlogPost, error) {
	post := &entity.BlogPost{
		Title:    req.Title,
		AuthorID: req.AuthorID,
		Content:  req.Content,
		Category: req.GetCategory(),
		Tags:     req.Tags,
	}
	return post, nil
}

func UpdatePostReq2Entity(req *blog.UpdatePostRequest) (*entity.BlogPost, error) {
	post := &entity.BlogPost{
		PostID:   req.BlogID,
		AuthorID: req.AuthorID,
		Title:    req.GetTitle(),
		Content:  req.GetContent(),
		Category: req.GetCategory(),
		Tags:     req.GetTags(),
	}
	return post, nil
}
