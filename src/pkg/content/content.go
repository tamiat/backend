package content

var Contents []Content

type Content struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Details string `json:"details"`
}
