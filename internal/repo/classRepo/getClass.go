package classRepo
import (
	"context"
	models "github.com/AbanoubGirges/Go-School-System/internal/models/class"
)
func (r *ClassRepo) GetClassById(id uint,ctx context.Context)(*models.Class,error){
	var class models.Class
	err := r.repo.DB.WithContext(ctx).First(&class, id).Error
	if err != nil {
		return nil, err
	}
	return &class, nil
}
func (r *ClassRepo) GetAllClasses(ctx context.Context)([]*models.Class,error){
	var classes []*models.Class
	err := r.repo.DB.WithContext(ctx).Find(&classes).Error
	if err != nil {
		return nil, err
	}
	return classes, nil
}