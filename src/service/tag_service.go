package service

import (
	"go-crud/src/data/request"
	"go-crud/src/data/response"
)

type TagService interface {
	Create(tag request.CreateTagRequest)
	Update(tag request.UpdateTagRequest)
	Delete(tagId int)
	FindById(tagId int) response.TagResponse
	FindAll() []response.TagResponse
}
