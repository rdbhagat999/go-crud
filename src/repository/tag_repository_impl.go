package repository

import (
	"errors"
	"go-crud/src/data/request"
	"go-crud/src/helper"
	"go-crud/src/model"

	"gorm.io/gorm"
)

type TagRepositoryImpl struct {
	Db *gorm.DB
}

// Delete implements TagRepository.
func (t *TagRepositoryImpl) Delete(tagId int) {
	deleteTag := request.DeleteTagRequest{
		ID: tagId,
	}
	var tag model.Tag

	result := t.Db.Model(&tag).Delete(deleteTag)
	// result := t.Db.Where("id=?", tagId).Delete(&tag)
	helper.ErrorPanic(result.Error)
}

// FindAll implements TagRepository.
func (t *TagRepositoryImpl) FindAll() []model.Tag {
	var tags []model.Tag
	result := t.Db.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags
}

// FindById implements TagRepository.
func (t *TagRepositoryImpl) FindById(tagId int) (tag model.Tag, err error) {
	var foundTag model.Tag
	result := t.Db.Find(&foundTag, tagId)

	// fmt.Println(foundTag)

	if result != nil {
		return foundTag, nil
	} else {
		return foundTag, errors.New("tag not found")
	}
}

// Save implements TagRepository.
func (t *TagRepositoryImpl) Save(tag model.Tag) {

	result := t.Db.Create(&tag)
	helper.ErrorPanic(result.Error)
}

// Update implements TagRepository.
func (t *TagRepositoryImpl) Update(tag model.Tag) {
	var updateTag = request.UpdateTagRequest{
		ID:   tag.ID,
		Name: tag.Name,
	}

	result := t.Db.Model(&tag).Updates(updateTag)
	helper.ErrorPanic(result.Error)
}

func NewTagRepositoryImpl(Db *gorm.DB) TagRepository {
	return &TagRepositoryImpl{Db: Db}
}
