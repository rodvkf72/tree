package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"

	"tree/controller"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Tree-Api World!")
	})

	// 사용자 데이터 제어
	e.GET("/users", func(c echo.Context) error {
		return controller.UserController(c, "userList")
	})
	e.POST("/users", func(c echo.Context) error {
		return controller.UserController(c, "userInsert")
	})
	e.PUT("/users", func(c echo.Context) error {
		return controller.UserController(c, "userUpdate")
	})
	e.DELETE("/users", func(c echo.Context) error {
		return controller.UserController(c, "userDelete")
	})

	err := e.Start(":9090")
	if err != nil {
		
	}
}