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
func (t *TagServiceImpl) Create(tag request.CreateTagRequest) response.TagResponse {
	err := t.Validate.Struct(tag)
	helper.ErrorPanic(err)

	tagModel := model.Tag{
		Name:   tag.Name,
		UserID: uint(tag.UserID),
	}

	result, resultErr := t.TagRepository.Save(tagModel)

	helper.ErrorPanic(resultErr)

	tagReponse := response.TagResponse{
		ID:     int(result.ID),
		Name:   result.Name,
		UserID: int(result.UserID),
	}

	return tagReponse
}

// Delete implements TagService.
func (t *TagServiceImpl) Delete(tagId int) {
	t.TagRepository.Delete(tagId)
}

// FindAll implements TagService.
func (t *TagServiceImpl) FindAll() []response.TagResponse {
	var tags = []response.TagResponse{}
	result, err := t.TagRepository.FindAll()
	helper.ErrorPanic(err)

	for _, v := range result {
		found := response.TagResponse{
			ID:     int(v.ID),
			Name:   v.Name,
			UserID: int(v.UserID),
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
		ID:     int(result.ID),
		Name:   result.Name,
		UserID: int(result.UserID),
	}

	return tagReponse
}

// Update implements TagService.
func (t *TagServiceImpl) Update(tag request.UpdateTagRequest) response.TagResponse {
	found, err := t.TagRepository.FindById(tag.ID)
	helper.ErrorPanic(err)

	found.Name = tag.Name

	result, resultErr := t.TagRepository.Update(found)

	helper.ErrorPanic(resultErr)

	tagReponse := response.TagResponse{
		ID:     int(result.ID),
		Name:   result.Name,
		UserID: int(result.UserID),
	}

	return tagReponse
}

func NewTagServiceImpl(tagRepository repository.TagRepository, validate *validator.Validate) TagService {
	return &TagServiceImpl{
		TagRepository: tagRepository,
		Validate:      validate,
	}
}
