package service

import (
	"go-crud/src/data/request"
	"go-crud/src/data/response"
	"go-crud/src/helper"
	"go-crud/src/model"
	"go-crud/src/repository"

	"github.com/go-playground/validator/v10"
)

type TagServiceImpl struct {
	TagRepository repository.TagRepository
	Validate      *validator.Validate
}

// Create implements TagService.
func (t *TagServiceImpl) Create(tag request.CreateTagRequest) {
	err := t.Validate.Struct(tag)
	helper.ErrorPanic(err)

	tagModel := model.Tag{
		Name: tag.Name,
	}

	t.TagRepository.Save(tagModel)
}

// Delete implements TagService.
func (t *TagServiceImpl) Delete(tagId int) {
	t.TagRepository.Delete(tagId)
}

// FindAll implements TagService.
func (t *TagServiceImpl) FindAll() []response.TagResponse {
	var tags []response.TagResponse
	result := t.TagRepository.FindAll()

	for _, v := range result {
		found := response.TagResponse{
			Id:   v.Id,
			Name: v.Name,
		}

		tags = append(tags, found)
	}

	return tags
}

// FindById implements TagService.
func (t *TagServiceImpl) FindById(tagId int) response.TagResponse {
	result, err := t.TagRepository.FindById(tagId)
	helper.ErrorPanic(err)

	tagReponse := response.TagResponse{
		Id:   result.Id,
		Name: result.Name,
	}

	return tagReponse
}

// Update implements TagService.
func (t *TagServiceImpl) Update(tag request.UpdateTagRequest) {
	found, err := t.TagRepository.FindById(tag.Id)
	helper.ErrorPanic(err)

	found.Name = tag.Name

	t.TagRepository.Update(found)
}

func NewTagServiceImpl(tagRepository repository.TagRepository, validate *validator.Validate) TagService {
	return &TagServiceImpl{
		TagRepository: tagRepository,
		Validate:      validate,
	}
}
