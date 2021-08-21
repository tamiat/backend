package content

// Content  model
type Content struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Details string `json:"details"`
}
type ContentRepository interface {
	//TODO1
	FindAll() ([]Content, error)
	ById(string) (*Content,error)
	Post(content Content) (string, error)
	FindRange([]string) ([]Content, error)
	DeleteById(string) error
	UpdateById(string, Content) error
}