package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	error2 "github.com/negaranabestani/students/error"
	"github.com/negaranabestani/students/model"
	request2 "github.com/negaranabestani/students/request"
	"github.com/negaranabestani/students/store"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type Student struct {
	Store  store.Student
	Logger *zap.Logger
}

func (s Student) Get(c echo.Context) error {
	Id, err := strconv.ParseInt(c.Param("id"), 0, 64)
	if err != nil {
		s.Logger.Error("invalid id format. id is not parsable", zap.Error(err), zap.Any("id", c.Param("id")))
		return echo.ErrBadRequest
	}
	result, e := s.Store.Get(Id)
	if e != nil {
		var NotFoundErr error2.StudentNotFoundError
		if ok := errors.As(e, &NotFoundErr); ok {
			s.Logger.Error("student not found", zap.Error(e))
			return echo.ErrNotFound
		}
		s.Logger.Error("database failed to response", zap.Error(e))
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, result)

}
func (s Student) GetAll(c echo.Context) error {
	result, err := s.Store.GetAll()
	if err != nil {
		var EmptyErr error2.EmptyStoreError
		if errors.As(err, &EmptyErr) {
			return c.JSON(http.StatusOK, err)
		}
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, result)
}
func (s Student) save(c echo.Context) error {
	var request request2.Student
	if err := c.Bind(&request); err != nil {
		return echo.ErrBadRequest
	}
	if err := request.Validate(); err != nil {
		return echo.ErrBadRequest
	}
	newStudent := model.Student{Id: request.ID, Name: request.FirstName, LastName: request.LastName}
	if err := s.Store.Save(newStudent); err != nil {
		var DuplicateErr error2.DuplicateStudentError
		if ok := errors.As(err, &DuplicateErr); ok {
			return echo.ErrBadRequest
		}
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusCreated, newStudent)
}
func (s Student) Register(app *echo.Group) {
	app.GET("/:id", s.Get)
	app.GET("/students", s.GetAll)
	app.POST("/new", s.save)
}
