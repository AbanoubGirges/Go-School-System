package dto
import(
	"github.com/google/uuid"
	"github.com/AbanoubGirges/Go-School-System/internal/models/class"
)
type CreateClassRequest struct{
	Name string `json:"name"`
	Grade string `json:"grade"`
}
func (r *CreateClassRequest) BindToModel()*models.Class{
	id:=uuid.New()
	return &models.Class{
		ID: id,
		Name:r.Name,
		Grade: r.Grade,
	}
}