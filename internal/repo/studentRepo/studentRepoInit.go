package studentRepo

import (
	"github.com/AbanoubGirges/Go-School-System/internal/config"
)
type StudentRepo struct{
	repo *config.Repo
}
func NewStudentRepo(repo *config.Repo) *StudentRepo{
	return &StudentRepo{repo}
}
