package services

import (
	"github.com/urcop/go-fiber-template/internal/model"
	"github.com/urcop/go-fiber-template/internal/repository"
)

type LinkService interface {
	Create(link *model.Link) error
	Get(id string) (*model.Link, error)
	GetAll() ([]*model.Link, error)
	Update(link *model.Link) error
	Delete(id string) error
}

type LinkServiceImpl struct {
	repos repository.LinkRepository
}

func (s *LinkServiceImpl) Create(link *model.Link) error {
	return s.repos.CreateLink(link)
}

func (s *LinkServiceImpl) Get(id string) (*model.Link, error) {
	return s.repos.GetLink(id)
}

func (s *LinkServiceImpl) GetAll() ([]*model.Link, error) {
	return s.repos.GetAllLinks()
}

func (s *LinkServiceImpl) Update(link *model.Link) error {
	return s.repos.UpdateLink(link)
}

func (s *LinkServiceImpl) Delete(id string) error {
	return s.repos.DeleteLink(id)
}

func NewLinkService(repos repository.LinkRepository) *LinkServiceImpl {
	return &LinkServiceImpl{repos: repos}
}
