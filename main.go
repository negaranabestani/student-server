package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/negaranabestani/students/handler"
	store2 "github.com/negaranabestani/students/store"
	"go.uber.org/zap"
)

func main() {
	app := echo.New()
	var store store2.Student
	logger := &zap.Logger{}
	store = store2.NewStudentMemory(logger)

	hs := handler.Student{Store: store, Logger: logger}
	hs.Register(app.Group("api/student"))
	if err := app.Start(":8080"); err != nil {
		fmt.Println("cannot start server")
	}
}
