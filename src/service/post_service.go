package service

import (
	"go-crud/src/data/request"
	"go-crud/src/data/response"
	"go-crud/src/model"
)

type PostService interface {
	Create(post model.Post) (postdata response.PostResponse, err error)
	Update(id int, post request.UpdatePostRequest) (postdata response.PostResponse, err error)
	Delete(postId int)
	FindById(postId int) (postdata response.PostResponse, err error)
	FindAll() (postList []response.PostResponse, err error)
	FindAllByUserId(userId int) (postList []response.PostResponse, err error)
}
