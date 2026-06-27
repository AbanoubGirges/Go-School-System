package classRepo
import (
	"context"
	models "github.com/AbanoubGirges/Go-School-System/internal/models/class"
)
func (r *ClassRepo) UpdateClass(id uint, ctx context.Context) error{
	var class models.Class
	err := r.repo.DB.WithContext(ctx).Model(&class).Where("id = ?", id).Updates(&class).Error
	if err != nil {
		return err
	}
	return nil
}