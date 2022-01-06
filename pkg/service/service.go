package service

import (
	"github.com/mahou-shoujo-sss/bookmark"
	"github.com/mahou-shoujo-sss/bookmark/pkg/repository"
)

type Authorization interface {
	CreateUser(user bookmark.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type List interface {
	Create(userId int, list bookmark.List) (int, error)
	GetAll(userId int) ([]bookmark.List, error)
	GetById(userId, listId int) (bookmark.List, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input bookmark.UpdateListInput) error
}

type Service struct {
	Authorization
	List
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		List:          NewListService(repos.List),
	}
}
