package services

import (
	"github.com/urcop/go-fiber-template/internal/model"
	"github.com/urcop/go-fiber-template/internal/repository"
)

type LinkService interface {
	Create(link *model.Link) error
	Get(short_link string) (*model.Link, error)
	GetAll() ([]*model.Link, error)
	Update(link *model.Link) error
	Delete(id string) error
	UpdateClicked(link *model.Link) error
}

type LinkServiceImpl struct {
	repos repository.LinkRepository
}

func (s *LinkServiceImpl) Create(link *model.Link) error {
	return s.repos.CreateLink(link)
}

func (s *LinkServiceImpl) Get(short_link string) (*model.Link, error) {
	return s.repos.GetLink(short_link)
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

func (s *LinkServiceImpl) UpdateClicked(link *model.Link) error {
	link.Clicked++

	err := s.Update(link)
	if err != nil {
		return err
	}

	return nil
}

func NewLinkService(repos repository.LinkRepository) *LinkServiceImpl {
	return &LinkServiceImpl{repos: repos}
}
