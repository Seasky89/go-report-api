package services

import (
	"github.com/Seasky89/go-report-api/internal/models"
	"github.com/Seasky89/go-report-api/internal/store"
)

type UserService struct {
	store store.DataStore
}

func NewUserService(store store.DataStore) *UserService {
	return &UserService{store: store}
}

func (s *UserService) FindAll() []models.User {
	return s.store.GetUsers()
}
