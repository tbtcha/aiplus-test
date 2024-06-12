package service

import (
	"aiplus/internal/app/domain"
	"aiplus/internal/app/repository"
	"aiplus/internal/app/repository/entity"
	"context"
	"github.com/google/uuid"
)

type Service interface {
	CreateStaff(ctx context.Context, rq domain.StaffRequest) (*domain.StaffResponse, error)
}

type service struct {
	repo repository.Repository
}

func (s service) CreateStaff(ctx context.Context, rq domain.StaffRequest) (*domain.StaffResponse, error) {
	id := uuid.New().String()
	err := s.repo.CreateStaff(ctx, &entity.StaffEntity{
		ID:          id,
		Lastname:    rq.Lastname,
		Firstname:   rq.Firstname,
		Middlename:  rq.Middlename,
		PhoneNumber: rq.PhoneNumber,
		CityName:    rq.CityName,
	})
	return &domain.StaffResponse{
		ID:          id,
		Lastname:    rq.Lastname,
		Firstname:   rq.Firstname,
		Middlename:  rq.Middlename,
		PhoneNumber: rq.PhoneNumber,
		CityName:    rq.CityName,
	}, err
}

func NewService(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}
