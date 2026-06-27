package userRepo

import (
	"github.com/AbanoubGirges/Go-School-System/internal/config"
)
type UserRepo struct{
	repo *config.Repo
}
func NewUserRepo(repo *config.Repo) *UserRepo{
	return &UserRepo{repo}
}
