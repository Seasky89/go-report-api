package services

import (
	"github.com/Seasky89/go-report-api/internal/models"
	"github.com/Seasky89/go-report-api/internal/store"
)

type PostService struct {
	store store.DataStore
}

func NewPostService(store store.DataStore) *PostService {
	return &PostService{store: store}
}

func (s *PostService) FindAll() []models.Post {
	return s.store.GetPosts()
}
