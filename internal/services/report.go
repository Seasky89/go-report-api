package services

import (
	"github.com/Seasky89/go-report-api/internal/models"
	"github.com/Seasky89/go-report-api/internal/store"
)

type ReportService struct {
	store store.DataStore
}

func NewReportService(store store.DataStore) *ReportService {
	return &ReportService{store: store}
}

func (s *ReportService) GenerateReport() []models.Report {
	users := s.store.GetUsers()
	posts := s.store.GetPosts()

	postCount := make(map[uint]int)
	for _, post := range posts {
		postCount[post.UserId]++
	}

	var report []models.Report
	for _, user := range users {
		r := models.Report{
			UserId:     user.Id,
			Name:       user.Name,
			Email:      user.Email,
			TotalPosts: postCount[user.Id],
		}
		report = append(report, r)
	}

	return report
}
