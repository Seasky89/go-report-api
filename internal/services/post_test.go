package services_test

import (
	"testing"

	"github.com/Seasky89/go-report-api/internal/models"
	"github.com/Seasky89/go-report-api/internal/services"
)

func TestPostService_FindAll(t *testing.T) {
	mock := &mockStore{
		posts: []models.Post{
			{Id: 1, UserId: 1},
			{Id: 2, UserId: 1},
			{Id: 3, UserId: 2},
		},
	}

	services := services.NewPostService(mock)
	posts := services.FindAll()

	if len(posts) != 3 {
		t.Errorf("esperava 3 posts")
	}
}
