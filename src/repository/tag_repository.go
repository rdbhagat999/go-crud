package repository

import "go-crud/src/model"

type TagRepository interface {
	Save(tag model.Tag) (t model.Tag, err error)
	Update(tag model.Tag) (t model.Tag, err error)
	Delete(tagId int)
	FindById(tagId int) (tag model.Tag, err error)
	FindAll() []model.Tag
}
