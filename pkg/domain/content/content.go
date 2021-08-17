package content

//content model
type Content struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Details string `json:"details"`
}
type ContentRepository interface {
	//TODO 1
	FindAll() ([]Content, error)
	ById(string) (*Content,error)
}