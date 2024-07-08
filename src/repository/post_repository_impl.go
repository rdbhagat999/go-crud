package repository

import (
	"errors"
	"go-crud/src/data/request"
	"go-crud/src/helper"
	"go-crud/src/model"

	"gorm.io/gorm"
)

type PostRepositoryImpl struct {
	Db *gorm.DB
}

// Delete implements PostRepository.
func (p *PostRepositoryImpl) Delete(postId int) {
	deletePost := request.DeletePostRequest{
		ID: postId,
	}

	var post model.Post

	result := p.Db.Model(&post).Delete(deletePost)
	// result := t.Db.Where("id=?", postId).Delete(&post)
	helper.ErrorPanic(result.Error)
}

// FindAll implements PostRepository.
func (p *PostRepositoryImpl) FindAll() (posts []model.Post, err error) {
	var postList []model.Post
	err = p.Db.Model(&model.Post{}).Find(&postList).Error
	return postList, err
}

// FindById implements PostRepository.
func (p *PostRepositoryImpl) FindById(postId int) (post model.Post, err error) {
	var foundPost model.Post
	result := p.Db.Find(&foundPost, postId)

	// fmt.Println(foundPost)

	if result != nil {
		return foundPost, nil
	}

	return foundPost, errors.New("post not found")
}

// Save implements PostRepository.
func (p *PostRepositoryImpl) Save(post model.Post) (ps model.Post, err error) {
	result := p.Db.Create(&post)
	helper.ErrorPanic(result.Error)

	return p.FindById(int(post.ID))
}

// Update implements PostRepository.
func (p *PostRepositoryImpl) Update(post model.Post) (ps model.Post, err error) {
	var updatePost = request.UpdatePostRequest{
		ID:    int(post.ID),
		Title: post.Title,
		Body:  post.Body,
	}

	result := p.Db.Model(&post).Updates(updatePost)
	helper.ErrorPanic(result.Error)

	return p.FindById(int(post.ID))
}

func NewPostRepositoryImpl(Db *gorm.DB) PostRepository {
	return &PostRepositoryImpl{Db: Db}
}
