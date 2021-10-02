package contentType

type ContentType struct {
	ID    string    `json:"id"`
	Name string `json:"name"`
}

type ContentTypeRepository interface {

	//TODO 1
	isTableExists(string) (string, error)
	isColExists(string, string) error
	Create(string, string, string) (string, error)
	DeleteById(string, string) error
	UpdateColName(string, string, string, string) error
	AddCol(string, string, string) error
	DeleteCol(string, string, string) error
<<<<<<< HEAD
}
=======
}
>>>>>>> ff64c0ae41d76f7c9491088c13feda7152987af6
