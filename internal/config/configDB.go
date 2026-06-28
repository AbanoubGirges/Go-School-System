package config

import (
	userModels "github.com/AbanoubGirges/Go-School-System/internal/models/user"
	classModels "github.com/AbanoubGirges/Go-School-System/internal/models/class"
	studentModels "github.com/AbanoubGirges/Go-School-System/internal/models/student"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
type Repo struct{
	DB	*gorm.DB
}
func SetupDB(dbUrl string) *Repo{
	print(dbUrl)
	db,err:=gorm.Open(postgres.Open(dbUrl),&gorm.Config{})
	if err !=nil{
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&userModels.User{})
	db.AutoMigrate(&classModels.Class{})
	db.AutoMigrate(&classModels.Class_User{})
	db.AutoMigrate(&studentModels.Student{})
	db.AutoMigrate(&studentModels.Attendance{})
	return &Repo{db}
}