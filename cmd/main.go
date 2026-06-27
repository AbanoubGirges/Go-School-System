package main

import (
	"github.com/AbanoubGirges/Go-School-System/internal/config"
	authControllers "github.com/AbanoubGirges/Go-School-System/internal/controllers/authControllers"
	profileControllers "github.com/AbanoubGirges/Go-School-System/internal/controllers/profileControllers"
	"github.com/AbanoubGirges/Go-School-System/internal/middleware"
	userRepo "github.com/AbanoubGirges/Go-School-System/internal/repo/userRepo"
	"github.com/AbanoubGirges/Go-School-System/internal/routes"
	"github.com/gin-contrib/cors"
	classControllers "github.com/AbanoubGirges/Go-School-System/internal/controllers/classControllers"
	classRepo "github.com/AbanoubGirges/Go-School-System/internal/repo/classRepo"
	"github.com/gin-gonic/gin"
	"github.com/AbanoubGirges/Go-School-System/internal/repo/studentRepo"
	studentControllers "github.com/AbanoubGirges/Go-School-System/internal/controllers/studentControllers"
)

func main(){
	cfg:=config.Load()
	repo:=config.SetupDB(cfg.DBUrl)
	userRepo:=userRepo.NewUserRepo(repo)
	classRepo:=classRepo.NewClassRepo(repo)
	studentRepo:=studentRepo.NewStudentRepo(repo)
	authController:=authControllers.NewAuthController(userRepo,cfg.JwtSecret)
	editProfileController:=profileControllers.NewEditProfileController(userRepo,cfg.JwtSecret)
	classController:=classControllers.NewClassController(classRepo,cfg.JwtSecret)
	studentController:=studentControllers.NewStudentController(studentRepo,cfg.JwtSecret)
	userMiddleware:=middleware.NewUserMiddleware(cfg.JwtSecret)
	router:=gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	routes.RouterSetup(router,userMiddleware,authController,editProfileController,classController,studentController)
	router.Run(":"+cfg.Port)
}