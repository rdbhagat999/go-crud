package service

import (
	"go-crud/src/data/request"
	"go-crud/src/data/response"
)

type TagService interface {
	Create(tag request.CreateTagRequest) response.TagResponse
	Update(tag request.UpdateTagRequest) response.TagResponse
	Delete(tagId int)
	FindById(tagId int) response.TagResponse
	FindAll() []response.TagResponse
}
