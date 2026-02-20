package models

type Post struct {
	Id     uint   `json:"id"`
	UserId uint   `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
