package repository

import (
	"context"

	"github.com/ibnumei/go-ms-playground/internal/app/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur UserRepository) Create(ctx context.Context, user *domain.User) error {
	return ur.db.WithContext(ctx).Create(user).Error
}