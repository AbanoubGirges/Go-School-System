package dto
import (
	"github.com/google/uuid"
	"github.com/AbanoubGirges/Go-School-System/internal/models/user"
)
type RegisterRequest struct {
	Name string `json:"name" binding:"required"`
	PhoneNumber  string  `json:"phoneNumber" binding:"required,min=11"`
	Password	string	`json:"password" binding:"required"`
	Email	string	`json:"email" binding:"required,email"`

}
func (r *RegisterRequest) BindToModel()*models.User{
	id:=uuid.New()
	role:=models.UserRoles.User
	return &models.User{
		ID: id,
		Name:r.Name,
		PhoneNumber: r.PhoneNumber,
		Password: r.Password,
		Email: r.Email,
		Role: role,
	}

}
type LoginRequest struct{
	PhoneNumber  string  `json:"phoneNumber" binding:"required,min=11"`
	Password	string	`json:"password" binding:"required"`
}