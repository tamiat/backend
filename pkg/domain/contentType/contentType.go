package contentType

type ContentType struct {
	ID    string    `json:"id"`
	Name string `json:"name"`
}

type ContentTypeRepository interface {

	//TODO 1
	isTableExists(string) (string, error)
	isColExists(string, string) error
	Create(string, string) (string, error)
	DeleteById(string) error
	UpdateColName(string, string, string) error
	AddCol(string, string) error
	DeleteCol(string, string) error
}
