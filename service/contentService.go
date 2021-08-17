package service

import (
	"github.com/tamiat/backend/domain/content"
)

type ContentService interface {
	//TODO 3
	GetAllCustomers()([]content.Content,error)
	GetCustomer(string) (*content.Content,error)
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
func NewCustomerService(repository content.ContentRepository) DefaultContentService{
	return DefaultContentService{repo: repository}
}