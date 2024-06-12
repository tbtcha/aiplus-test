package repository

import (
	"aiplus/internal/app/repository/entity"
	"context"
	"gorm.io/gorm"
)

type Repository interface {
	CreateStaff(ctx context.Context, e *entity.StaffEntity) error
}

type repository struct {
	db *gorm.DB
}

func (r repository) CreateStaff(ctx context.Context, e *entity.StaffEntity) error {
	return r.db.WithContext(ctx).Create(e).Error
}

func NewRepository(db *gorm.DB) (Repository, error) {
	err := db.AutoMigrate(
		&entity.StaffEntity{},
	)
	if err != nil {
		return nil, err
	}

	return &repository{
		db: db,
	}, nil
}
