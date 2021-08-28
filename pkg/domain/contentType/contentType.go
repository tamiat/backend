package contentType

type ContentTypeRepository interface {

	//TODO 1
	isTableExists(string) (string, error)
	Create(string, string) (string, error)
	DeleteById(string) error
	UpdateColName(string, string, string) error
	AddCol(string, string) error
}
