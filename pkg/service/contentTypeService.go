package service

import (
	"github.com/tamiat/backend/pkg/domain/contentType"
)

type ContentTypeService interface {
	CreateContentType(int, string, string) (string, error)
	DeleteContentType(int, string) error
	UpdateColName(int, string, string, string) error
	AddCol(int, string, string) error
	DeleteCol(int, string, string) error
}

type DefaultContentTypeService struct {
	repo contentType.ContentTypeRepository
}

func (s DefaultContentTypeService) CreateContentType(userId int, name, cols string) (string, error) {
	return s.repo.Create(userId, name, cols)
}

func (s DefaultContentTypeService) DeleteContentType(userId int, contentTypeId string) error {
	return s.repo.DeleteById(userId, contentTypeId)
}

func (s DefaultContentTypeService) UpdateColName(userId int, contentTypeId, oldName, newName string) error {
	return s.repo.UpdateColName(userId, contentTypeId, oldName, newName)
}

func (s DefaultContentTypeService) AddCol(userId int, contentTypeId, col string) error {
	return s.repo.AddCol(userId, contentTypeId, col)
}

func (s DefaultContentTypeService) DeleteCol(userId int, contentTypeId, col string) error {
	return s.repo.DeleteCol(userId, contentTypeId, col)
}

func NewContentTypeService(repository contentType.ContentTypeRepository) DefaultContentTypeService {
	return DefaultContentTypeService{repo: repository}
}
