package contentType

type ContentTypeRepository interface {

	//TODO 1
	Create(string, string) (string, error)
	DeleteById(string) error
}
