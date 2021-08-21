package service

import (
	"github.com/tamiat/backend/pkg/domain/content"
)
type ContentService interface {
	GetAllContents() ([]content.Content, error)
	GetContent(string) (*content.Content, error)
	GetRangeOfContents([]string) ([]content.Content, error)
	PostContent(content.Content) (string, error)
	DeleteContent(string) error
	UpdateContent(string, content.Content) error
}

type DefaultContentService struct {
	repo content.ContentRepository
}

func (s DefaultContentService) GetAllContents() ([]content.Content, error) {
	return s.repo.FindAll()
}
func (s DefaultContentService) GetContent(id string) (*content.Content, error) {
	return s.repo.ById(id)
}
func (s DefaultContentService) GetRangeOfContents(ids []string) ([]content.Content, error) {
	return s.repo.FindRange(ids)
}
func (s DefaultContentService) PostContent(newContent content.Content) (string, error) {
	return s.repo.Post(newContent)
}
func (s DefaultContentService) DeleteContent(id string) error {
	return s.repo.DeleteById(id)
}
func (s DefaultContentService) UpdateContent(id string, updContent content.Content) error {
	return s.repo.UpdateById(id, updContent)
}
func NewContentService(repository content.ContentRepository) DefaultContentService {
	return DefaultContentService{repo: repository}
}
