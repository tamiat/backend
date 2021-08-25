package service

import (
	"github.com/tamiat/backend/pkg/domain/contentType"
)

type ContentTypeService interface {
	CreateContentType(string, string) (string, error)
	DeleteContentType(id string) error
}

type DefaultContentTypeService struct {
	repo contentType.ContentTypeRepository
}

func (s DefaultContentTypeService) CreateContentType(name string, cols string) (string, error) {
	return s.repo.Create(name, cols)
}

func (s DefaultContentTypeService) DeleteContentType(id string) error {
	return s.repo.DeleteById(id)
}

func NewContentTypeService(repository contentType.ContentTypeRepository) DefaultContentTypeService {
	return DefaultContentTypeService{repo: repository}
}
