package userRepo

import (
	"context"

	models "github.com/AbanoubGirges/Go-School-System/internal/models/user"
)
func (r *UserRepo) CreateUser(user *models.User,ctx context.Context)error {
	return r.repo.DB.WithContext(ctx).Create(user).Error
}