package service

import (
	"github.com/mahou-shoujo-sss/bookmark"
	"github.com/mahou-shoujo-sss/bookmark/pkg/repository"
)

type ListService struct {
	repo repository.List
}

func NewListService(repo repository.List) *ListService {
	return &ListService{repo: repo}
}

func (s *ListService) Create(userId int, list bookmark.List) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *ListService) GetAll(userId int) ([]bookmark.List, error) {
	return s.repo.GetAll(userId)
}

func (s *ListService) GetById(userId, listId int) (bookmark.List, error) {
	return s.repo.GetById(userId, listId)
}

func (s *ListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *ListService) Update(userId, listId int, input bookmark.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, listId, input)
}
