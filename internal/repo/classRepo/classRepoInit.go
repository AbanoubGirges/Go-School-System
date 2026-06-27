package classRepo
import (
	"github.com/AbanoubGirges/Go-School-System/internal/config"
)
type ClassRepo struct{
	repo *config.Repo
}
func NewClassRepo(repo *config.Repo) *ClassRepo{
	return &ClassRepo{repo}
}
