package service

import (
	"fmt"
	"go-crud/src/data/request"
	"go-crud/src/data/response"
	"go-crud/src/model"
	"go-crud/src/repository"

	"github.com/go-playground/validator/v10"
)

type PostServiceImpl struct {
	PostRepository repository.PostRepository
	Validate       *validator.Validate
}

func postServicePrintln(err error) {
	fmt.Println("postServicePrintln: " + err.Error())
}

// Create implements PostService.
func (p *PostServiceImpl) Create(post request.CreatePostRequest) (postdata response.PostResponse, err error) {
	validateErr := p.Validate.Struct(post)
	// helper.ErrorPanic(validateErr)

	if validateErr != nil {
		postServicePrintln(validateErr)
		return response.PostResponse{}, validateErr
	}

	postModel := model.Post{
		Title:  post.Title,
		Body:   post.Body,
		UserID: uint(post.UserID),
	}

	result, resultErr := p.PostRepository.Save(postModel)
	// helper.ErrorPanic(resultErr)
	if resultErr != nil {
		postServicePrintln(resultErr)
		return response.PostResponse{}, resultErr
	}

	postReponse := response.PostResponse{
		ID:     int(result.ID),
		Title:  result.Title,
		Body:   result.Body,
		UserID: int(result.UserID),
	}

	return postReponse, nil
}

// Delete implements PostService.
func (p *PostServiceImpl) Delete(postId int) {
	p.PostRepository.Delete(postId)
}

// FindAll implements PostService.
func (p *PostServiceImpl) FindAll() (postList []response.PostResponse, err error) {
	var posts = []response.PostResponse{}

	result, resultErr := p.PostRepository.FindAll()
	// helper.ErrorPanic(resultErr)

	if resultErr != nil {
		postServicePrintln(resultErr)
		return posts, resultErr
	}

	for _, v := range result {
		found := response.PostResponse{
			ID:     int(v.ID),
			Title:  v.Title,
			Body:   v.Body,
			UserID: int(v.UserID),
		}

		posts = append(posts, found)
	}

	return posts, nil
}

// FindById implements PostService.
func (p *PostServiceImpl) FindById(postId int) (postdata response.PostResponse, err error) {
	var postReponse response.PostResponse

	result, findErr := p.PostRepository.FindById(postId)
	// helper.ErrorPanic(findErr)

	if findErr != nil {
		postServicePrintln(findErr)
		return postReponse, findErr
	}

	postReponse = response.PostResponse{
		ID:     int(result.ID),
		Title:  result.Title,
		Body:   result.Body,
		UserID: int(result.UserID),
	}

	return postReponse, nil
}

// Update implements PostService.
func (p *PostServiceImpl) Update(id int, post request.UpdatePostRequest) (postdata response.PostResponse, err error) {
	updateErr := p.Validate.Struct(post)

	if updateErr != nil {
		postServicePrintln(updateErr)
		return response.PostResponse{}, updateErr
	}

	found, foundErr := p.PostRepository.FindById(id)
	// helper.ErrorPanic(foundErr)

	if foundErr != nil {
		postServicePrintln(foundErr)
		return response.PostResponse{}, foundErr
	}

	found.ID = uint(id)
	found.Title = post.Title
	found.Body = post.Body

	result, resultErr := p.PostRepository.Update(found)
	// helper.ErrorPanic(resultErr)
	if resultErr != nil {
		postServicePrintln(resultErr)
		return response.PostResponse{}, resultErr
	}

	postReponse := response.PostResponse{
		ID:     int(result.ID),
		Title:  result.Title,
		Body:   result.Body,
		UserID: int(result.UserID),
	}

	return postReponse, nil
}

func NewPostServiceImpl(postRepository repository.PostRepository, validate *validator.Validate) PostService {
	return &PostServiceImpl{
		PostRepository: postRepository,
		Validate:       validate,
	}
}
