package repository

import (
	"fmt"
	"go-crud/src/data/request"
	"go-crud/src/helper"
	"go-crud/src/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostRepositoryImpl struct {
	Db *gorm.DB
}

func postRepositoryPrintln(err error) {
	fmt.Println("postRepositoryPrintln: " + err.Error())
}

// Delete implements PostRepository.
func (p *PostRepositoryImpl) Delete(postId int) {
	deletePost := request.DeletePostRequest{
		ID: postId,
	}
	var post model.Post

	delErr := p.Db.Model(&post).Delete(deletePost).Error
	// result := t.Db.Where("id=?", postId).Delete(&post)
	// helper.ErrorPanic(delErr)

	if delErr != nil {
		postRepositoryPrintln(delErr)
		helper.ErrorPanic(delErr)
	}
}

// FindAll implements PostRepository.
func (p *PostRepositoryImpl) FindAll() (posts []model.Post, err error) {
	var postList []model.Post

	postListErr := p.Db.Model(&model.Post{}).Preload(clause.Associations).Find(&postList).Error

	if postListErr != nil {
		postRepositoryPrintln(postListErr)
		return postList, postListErr
	}

	return postList, nil
}

// FindById implements PostRepository.
func (p *PostRepositoryImpl) FindById(postId int) (post model.Post, err error) {
	var foundPost model.Post

	findErr := p.Db.Model(&model.Post{}).Preload(clause.Associations).Find(&foundPost, postId).Error
	fmt.Println("foundPost", foundPost)

	if findErr != nil {
		postRepositoryPrintln(findErr)
		return foundPost, findErr
	}

	// if foundPost.ID == 0 {
	// 	return foundPost, errors.New("post not found")
	// }

	return foundPost, nil
}

// Save implements PostRepository.
func (p *PostRepositoryImpl) Save(post model.Post) (ps model.Post, err error) {

	createPost := request.CreatePostRequest{
		Title: post.Title,
		Body:  post.Body,
	}

	createErr := p.Db.Model(&post).Create(createPost).Error
	// helper.ErrorPanic(createErr)

	if createErr != nil {
		postRepositoryPrintln(createErr)
		return model.Post{}, createErr
	}

	return p.FindById(int(post.ID))
}

// Update implements PostRepository.
func (p *PostRepositoryImpl) Update(post model.Post) (ps model.Post, err error) {

	updatePost := request.UpdatePostRequest{
		Title: post.Title,
		Body:  post.Body,
	}

	updateErr := p.Db.Model(&post).Updates(updatePost).Error
	// helper.ErrorPanic(updateErr)

	if updateErr != nil {
		postRepositoryPrintln(updateErr)
		return model.Post{}, updateErr
	}

	return p.FindById(int(post.ID))
}

func NewPostRepositoryImpl(Db *gorm.DB) PostRepository {
	return &PostRepositoryImpl{Db: Db}
}
