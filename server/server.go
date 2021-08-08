package server

import (
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/immanoj16/echo-gorm/storage"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type (
	Server struct {
		EchoClient *echo.Echo
		DB         *gorm.DB
	}

	Context struct {
		echo.Context
		DB *gorm.DB
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewServer() *Server {
	db := storage.NewDB()
	return &Server{
		EchoClient: echo.New(),
		DB:         db,
	}
}

func (s *Server) GetDB() *gorm.DB {
	return s.DB
}

func (s *Server) GetEchoClient() *echo.Echo {
	return s.EchoClient
}

func (s *Server) AutoMigrate(models ...interface{}) {
	err := s.DB.AutoMigrate(models...)
	if err != nil {
		log.Panic("failed to migrate")
	}
}

func (s *Server) AddValidator() {
	s.EchoClient.Validator = &CustomValidator{validator: validator.New()}
}
