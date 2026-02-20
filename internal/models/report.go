package models

type Report struct {
	UserId     uint   `json:"user_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	TotalPosts int    `json:"total_posts"`
}
