package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/mahou-shoujo-sss/bookmark"
)

type Authorization interface {
	CreateUser(user bookmark.User) (int, error)
	GetUser(username, password string) (bookmark.User, error)
}

type List interface {
	Create(userId int, list bookmark.List) (int, error)
	GetAll(userId int) ([]bookmark.List, error)
	GetById(userId, listId int) (bookmark.List, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input bookmark.UpdateListInput) error
}

type Repository struct {
	Authorization
	List
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		List:          NewListPostgres(db),
	}
}
