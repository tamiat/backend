package contentType

type ContentType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ContentTypeRepository interface {

	//TODO 1
	isTableExists(string) (string, error)
	isColExists(string, string) error
	Create(int, string, string) (string, error)
	DeleteById(int, string) error
	UpdateColName(int, string, string, string) error
	AddCol(int, string, string) error
	DeleteCol(int, string, string) error
}
