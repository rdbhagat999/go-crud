package repository

import "go-crud/src/model"

type PostRepository interface {
	Save(post model.Post) (savedPost model.Post, err error)
	Update(post model.Post) (updatedPost model.Post, err error)
	Delete(postId int)
	FindById(postId int) (foundPost model.Post, err error)
	FindAll() (posts []model.Post, err error)
}
