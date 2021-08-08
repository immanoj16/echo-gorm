package repository

import (
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type (
	IBase interface {
		Get(e echo.Context) error
		GetAll(e echo.Context) error
		Create(e echo.Context) error
		Update(e echo.Context) error
		// Delete(e echo.Context) error
		SetRoutes(route string)
	}

	Base struct {
		DB     *gorm.DB
		Client *echo.Echo
		Model  interface{}
		dType  reflect.Type
	}
)

func (b *Base) setReflectType() {
	b.dType = reflect.TypeOf(b.Model).Elem()
}

func (b *Base) getInstance() interface{} {
	return reflect.New(b.dType).Interface()
}

func (b *Base) getInstances() interface{} {
	return reflect.New(reflect.SliceOf(b.dType)).Interface()
}

func (b *Base) Get(c echo.Context) error {
	id := c.Param("id")
	instance := b.getInstance()
	if err := b.DB.Debug().Model(instance).First(instance, id).Error; err != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, instance)
}

func (b *Base) GetAll(c echo.Context) error {
	instances := b.getInstances()
	if err := b.DB.Debug().Model(b.Model).Find(instances).Error; err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, instances)
}

func (b *Base) Create(c echo.Context) error {
	instance := b.getInstance()

	if err := c.Bind(instance); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(instance); err != nil {
		return err
	}

	if err := b.DB.Debug().Model(b.Model).Create(instance).Error; err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, instance)
}

func (b *Base) Update(c echo.Context) error {
	id := c.Param("id")
	instance := b.getInstance()

	if err := b.DB.Debug().Model(b.Model).Where("id = ?", id).First(instance).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(instance); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(instance); err != nil {
		return err
	}

	if err := b.DB.Debug().Model(b.Model).Save(instance).Error; err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, instance)
}

func (b *Base) SetRoutes(route string) {
	b.setReflectType()
	g := b.Client.Group(route)
	g.GET("", b.GetAll)
	g.GET("/:id", b.Get)
}

var _ IBase = (*Base)(nil)
