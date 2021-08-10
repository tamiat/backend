package models

type Content struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Details string `json:"details"`
}

var ContentArr = []Content{
	{Id: 0, Title: "hello world", Details: "Welcome to Tamiat CMS!"},
	{Id: 1, Title: "hello world 1", Details: "Welcome to Tamiat CMS!"},
	{Id: 2, Title: "hello world 2", Details: "Welcome to Tamiat CMS!"},
}
