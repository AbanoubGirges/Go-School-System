package userRepo

import (
	"context"

	models "github.com/AbanoubGirges/Go-School-System/internal/models/user"
)
func (r *UserRepo) EditUserById(id string, fieldsToUpdate map[string]interface{}, ctx context.Context) error {
	return r.repo.DB.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Updates(fieldsToUpdate).Error
}