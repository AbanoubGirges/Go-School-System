package userRepo

import (
	"context"
	models "github.com/AbanoubGirges/Go-School-System/internal/models/user"
)
func (r *UserRepo) DeleteUserById(id string, ctx context.Context) error {
	return r.repo.DB.WithContext(ctx).Where("id = ?", id).Delete(&models.User{}).Error
}