package service

import (
	"github.com/tamiat/backend/pkg/domain/content"
)

type ContentService interface {
	//TODO 3
	GetAllContents()([]content.Content,error)
	GetContent(string) (*content.Content,error)
}

type DefaultContentService struct{
	repo content.ContentRepository
}

func (s DefaultContentService) GetAllContents()([]content.Content,error){
	return s.repo.FindAll()
}
func (s DefaultContentService) GetContent(id string)(*content.Content,error){
	return s.repo.ById(id)
}
func NewContentService(repository content.ContentRepository) DefaultContentService {
	return DefaultContentService{repo: repository}
}