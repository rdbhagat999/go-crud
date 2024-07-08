package repository

import "go-crud/src/model"

type TagRepository interface {
	Save(tag model.Tag) (savedTag model.Tag, err error)
	Update(tag model.Tag) (updatedTag model.Tag, err error)
	Delete(tagId int)
	FindById(tagId int) (foundTag model.Tag, err error)
	FindAll() (tags []model.Tag, err error)
}
