package repository

import "go-crud/model"

type TagRepository interface {
	Save(tag model.Tag)
	Update(tag model.Tag)
	Delete(tagId int)
	FindById(tagId int) (tag model.Tag, err error)
	FindAll() []model.Tag
}
