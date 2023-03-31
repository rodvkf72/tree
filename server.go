package main

import (
	"net/http"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"tree/controller"
	"tree/util"
	//"tree/resource"
)

func main() {
	e := echo.New()

	var mainLog = new(util.File)
	mainLog.Path = "./logs/MainAccess"
	mainLog.Name = "MainAccess"

	f := util.Logging(mainLog, "Run On Server", "", nil)

	defer f.Close()

	// 로그 데이터 포맷 방식을 지정하고 파일에 기록
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: ` - "time":"${time_rfc3339}", "remote_ip":"${remote_ip}", ` +
			`"host":"${host}", "method":"${method}", "uri":"${uri}", "user_agent":"${user_agent}",` +
			`"status":${status}, ` + "\n",
		Output: f,
	}))

	// 미들웨어에서 복구를 사용
	e.Use(middleware.Recover())

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

	// 서버 가동
	err := e.Start(":9090")
	// 서버 시작 시 문제가 있다면
	if err != nil {
		panic(fmt.Sprintf("error opening file : %v", err))
	}
}