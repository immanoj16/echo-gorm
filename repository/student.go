package repository

import (
	"github.com/immanoj16/echo-gorm/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Student struct {
	Base
}

func NewStudentRepo(db *gorm.DB, client *echo.Echo) *Student {
	repo := new(Student)
	repo.DB = db
	repo.Client = client
	repo.Model = &model.Student{}
	return repo
}
