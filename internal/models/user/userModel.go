package models
import (
	"github.com/google/uuid"
)
var UserRoles = struct {
	User  string
	Admin string
	Sudo  string
}{
	User:  "USER",
	Admin: "ADMIN",
	Sudo:  "SUDO",
}
var UserStatus = struct {
	Approved   string
	Pending string
	Rejected  string
}{
	Approved:   "APPROVED",
	Pending: "PENDING",
	Rejected:  "REJECTED",
}

type User struct {
	ID          uuid.UUID	`gorm:"primaryKey"`
	Name    string
	PhoneNumber string	`gorm:"uniqueIndex"`
	Password    string
	Email       string	`gorm:"uniqueIndex"`
	Role        string
	Class       uuid.UUID	`gorm:"foreignKey:Class"`
	Status      string
}
type UserToken struct {
	ID          uuid.UUID	
	Name	string
	PhoneNumber string
	Email       string
	Role        string
	Class	   uuid.UUID	`gorm:"foreignKey:Class"`
	Status      string
}
func (u *User) BindToToken() *UserToken {
	return &UserToken{
		ID: u.ID,
		Name: u.Name,
		PhoneNumber: u.PhoneNumber,
		Email: u.Email,
		Role: u.Role,
		Class: u.Class,
		Status: u.Status,
	}
}