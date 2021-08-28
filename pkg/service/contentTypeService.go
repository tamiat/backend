package service

import (
	"github.com/tamiat/backend/pkg/domain/contentType"
)

type ContentTypeService interface {
	CreateContentType(string, string) (string, error)
	DeleteContentType(string) error
	UpdateColName(string, string, string) error
	AddCol(string, string) error
	DeleteCol(string, string) error
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

func (s DefaultContentTypeService) UpdateColName(id string, oldName string, newName string) error {
	return s.repo.UpdateColName(id, oldName, newName)
}

func (s DefaultContentTypeService) AddCol(id string, col string) error {
	return s.repo.AddCol(id, col)
}

func (s DefaultContentTypeService) DeleteCol(id string, col string) error {
	return s.repo.DeleteCol(id, col)
}

func NewContentTypeService(repository contentType.ContentTypeRepository) DefaultContentTypeService {
	return DefaultContentTypeService{repo: repository}
}
