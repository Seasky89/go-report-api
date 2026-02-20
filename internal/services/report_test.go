package services_test

import (
	"testing"

	"github.com/Seasky89/go-report-api/internal/models"
	"github.com/Seasky89/go-report-api/internal/services"
)

func TestReportService_GenerateReport(t *testing.T) {
	mock := &mockStore{
		users: []models.User{
			{Id: 1, Name: "João", Email: "joao@email.com"},
			{Id: 2, Name: "Maria", Email: "maria@email.com"},
		},
		posts: []models.Post{
			{UserId: 1},
			{UserId: 1},
			{UserId: 2},
		},
	}

	services := services.NewReportService(mock)
	report := services.GenerateReport()

	if len(report) != 2 {
		t.Errorf("esperava 2 usuários no relatório, recebeu %d", len(report))
	}

	if report[0].TotalPosts != 2 {
		t.Errorf("esperava 2 posts para João, recebeu %d", report[0].TotalPosts)
	}

	if report[1].TotalPosts != 1 {
		t.Errorf("esperava 1 posta para Maria, recebeu %d", report[1].TotalPosts)
	}
}
