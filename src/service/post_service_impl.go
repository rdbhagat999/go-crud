package service

import (
	"fmt"
	"go-crud/src/data/request"
	"go-crud/src/data/response"
	"go-crud/src/helper"
	"go-crud/src/model"
	"go-crud/src/repository"

	"github.com/go-playground/validator/v10"
)

type PostServiceImpl struct {
	PostRepository repository.PostRepository
	Validate       *validator.Validate
}

// Create implements PostService.
func (p *PostServiceImpl) Create(post request.CreatePostRequest) response.PostResponse {
	err := p.Validate.Struct(post)
	fmt.Println(err.Error())
	helper.ErrorPanic(err)

	postModel := model.Post{
		Title:  post.Title,
		Body:   post.Body,
		UserID: uint(post.UserID),
	}

	result, resultErr := p.PostRepository.Save(postModel)

	helper.ErrorPanic(resultErr)

	postReponse := response.PostResponse{
		ID:     int(result.ID),
		Title:  result.Title,
		Body:   result.Body,
		UserID: int(result.UserID),
	}

	return postReponse
}

// Delete implements PostService.
func (p *PostServiceImpl) Delete(postId int) {
	p.PostRepository.Delete(postId)
}

// FindAll implements PostService.
func (p *PostServiceImpl) FindAll() []response.PostResponse {
	var posts = []response.PostResponse{}
	result, err := p.PostRepository.FindAll()
	fmt.Println(err.Error())
	helper.ErrorPanic(err)

	for _, v := range result {
		found := response.PostResponse{
			ID:     int(v.ID),
			Title:  v.Title,
			Body:   v.Body,
			UserID: int(v.UserID),
		}

		posts = append(posts, found)
	}

	return posts
}

// FindById implements PostService.
func (p *PostServiceImpl) FindById(postId int) response.PostResponse {
	result, err := p.PostRepository.FindById(postId)
	fmt.Println(err.Error())
	helper.ErrorPanic(err)

	postReponse := response.PostResponse{
		ID:     int(result.ID),
		Title:  result.Title,
		Body:   result.Body,
		UserID: int(result.UserID),
	}

	return postReponse
}

// Update implements PostService.
func (p *PostServiceImpl) Update(post request.UpdatePostRequest) response.PostResponse {
	err := p.Validate.Struct(post)
	fmt.Println(err.Error())

	found, foundErr := p.PostRepository.FindById(post.ID)
	fmt.Println(foundErr.Error())
	helper.ErrorPanic(foundErr)

	found.Title = post.Title
	found.Body = post.Body

	result, resultErr := p.PostRepository.Update(found)
	helper.ErrorPanic(resultErr)

	postReponse := response.PostResponse{
		ID:     int(result.ID),
		Title:  result.Title,
		Body:   result.Body,
		UserID: int(result.UserID),
	}

	return postReponse
}

func NewPostServiceImpl(postRepository repository.PostRepository, validate *validator.Validate) PostService {
	return &PostServiceImpl{
		PostRepository: postRepository,
		Validate:       validate,
	}
}
