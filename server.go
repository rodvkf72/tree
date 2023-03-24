package main

import (
	"net/http"
	"time"
	"os"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"tree/controller"
)

func main() {
	e := echo.New()

	now := time.Now()
	custom := now.Format("2006-01-02")
	fileName := custom + "_log.txt"

	if _, err := os.Stat("./logs"); err != nil {	//폴더가 존재하지 않는 경우
		os.MkdirAll("./logs", os.ModePerm)
	}

	f, file := os.OpenFile("./logs/" + fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if file != nil {
		panic(fmt.Sprintf("error opening file : %v", file))
	}

	defer f.Close()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339}", "remote_ip":"${remote_ip}", ` +
			`"host":"${host}", "method":"${method}", "uri":"${uri}", "user_agent":"${user_agent}",` +
			`"status":${status}, ` + "\n",
		Output: f,
	}))

	e.Use(middleware.Recover()) //미들웨어에서 복구를 사용

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