package service

import (
	"github.com/tamiat/backend/pkg/domain/type"
)

type TypeService interface {
	Create(_type.Type) (int, error)
	Read() ([]_type.Type, error)
	Update(_type.Type, string) error
	Delete(string) error
}

type DefaultTypeService struct {
	repo _type.TypeRepository
}

func (s DefaultTypeService) Create(newType _type.Type) (int, error) {
	return s.repo.Create(newType)
}

func (s DefaultTypeService) Read() ([]_type.Type, error) {
	return s.repo.Read()
}

func (s DefaultTypeService) Update(targetedType _type.Type, id string) error {
	return s.repo.Update(targetedType, id)
}

func (s DefaultTypeService) Delete(id string) error {
	return s.repo.Delete(id)
}

func NewTypeService(repository _type.TypeRepository) DefaultTypeService {
	return DefaultTypeService{repo: repository}
}
