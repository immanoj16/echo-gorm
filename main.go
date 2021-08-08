package main

import (
	"github.com/immanoj16/echo-gorm/model"
	"github.com/immanoj16/echo-gorm/repository"
	"github.com/immanoj16/echo-gorm/server"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	s := server.NewServer()
	s.AutoMigrate(&model.Student{})
	s.AddValidator()
	echoClient := s.GetEchoClient()

	echoClient.Use(middleware.Logger())
	echoClient.Use(middleware.Recover())

	studentRepo := repository.NewStudentRepo(s.GetDB(), echoClient)
	studentRepo.SetRoutes("/students")

	echoClient.Logger.Fatal(echoClient.Start(":1323"))
}
