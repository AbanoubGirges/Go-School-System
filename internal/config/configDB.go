package config

import (
	models "github.com/AbanoubGirges/Go-School-System/internal/models/user"
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
	db.AutoMigrate(&models.User{})
	return &Repo{db}
}