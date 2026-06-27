package userRepo

import (
	"context"

	models "github.com/AbanoubGirges/Go-School-System/internal/models/user"
)
func (r *UserRepo) GetUserByPhone(phone string,ctx context.Context)(*models.User,error){
	var user models.User
	err := r.repo.DB.WithContext(ctx).Where("phone_number = ?", phone).First(&user).Error
	return &user, err
}