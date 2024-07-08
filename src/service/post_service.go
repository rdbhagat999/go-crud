package service

import (
	"go-crud/src/data/request"
	"go-crud/src/data/response"
)

type PostService interface {
	Create(post request.CreatePostRequest) response.PostResponse
	Update(post request.UpdatePostRequest) response.PostResponse
	Delete(postId int)
	FindById(postId int) response.PostResponse
	FindAll() []response.PostResponse
}
