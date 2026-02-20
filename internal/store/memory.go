package store

import "github.com/Seasky89/go-report-api/internal/models"

type DataStore interface {
	GetUsers() []models.User
	GetPosts() []models.Post
}

type MemoryStore struct {
	users []models.User
	posts []models.Post
}

func NewMemoryStore(users []models.User, posts []models.Post) *MemoryStore {
	return &MemoryStore{
		users: users,
		posts: posts,
	}
}

func (m *MemoryStore) GetUsers() []models.User {
	return m.users
}

func (m *MemoryStore) GetPosts() []models.Post {
	return m.posts
}
