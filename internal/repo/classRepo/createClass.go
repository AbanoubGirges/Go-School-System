package classRepo
import (
	"context"
	models "github.com/AbanoubGirges/Go-School-System/internal/models/class"
)
func (r *ClassRepo) CreateClass(class *models.Class,ctx context.Context)error {
	return r.repo.DB.WithContext(ctx).Create(class).Error
}