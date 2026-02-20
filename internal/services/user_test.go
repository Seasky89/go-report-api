package services_test

import (
	"testing"

	"github.com/Seasky89/go-report-api/internal/models"
	"github.com/Seasky89/go-report-api/internal/services"
)

func TestUserService_FindAll(t *testing.T) {
	mock := &mockStore{
		users: []models.User{
			{Id: 1, Name: "João"},
		},
	}

	service := services.NewUserService(mock)
	users := service.FindAll()

	if len(users) != 1 {
		t.Errorf("esperava 1 usuario")
	}
}
