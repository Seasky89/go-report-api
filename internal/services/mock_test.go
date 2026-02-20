package services_test

import "github.com/Seasky89/go-report-api/internal/models"

type mockStore struct {
	users []models.User
	posts []models.Post
}

func (m *mockStore) GetUsers() []models.User {
	return m.users
}

func (m *mockStore) GetPosts() []models.Post {
	return m.posts
}
