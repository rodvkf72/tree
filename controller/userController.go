package controller

import (
	"net/http"
	"encoding/json"
	//"strings"
	"fmt"

	"github.com/labstack/echo"

	"tree/service"
	"tree/dto"
)

func ReceiveController(c echo.Context) error {
	reqDiv := c.FormValue("amode")

	result := service.ReceiveService(reqDiv)

	return c.String(http.StatusOK, result)
}

// 비슷한 범위의 기능이라면 하나의 함수에서 처리 가능하도록 하기 위해 amode 라는 string 타입의 변수를 받음
func UserController(c echo.Context, amode string) error {
	switch amode {
	case "userList" :
		page := c.FormValue("page")
		byteResult := service.UserSelectListService(page)
		// key값은 string이고 value 값은 interface{} 인 v 라는 map 타입의 변수를 생성
		var v map[string]interface{}

		// byteResult 라는 값을 v 구조체로 언마샬링(디코딩)
		err := json.Unmarshal(byteResult, &v)
		if (err != nil) {
			// error log
		}

		return c.JSON(http.StatusOK, v)

	case "userInsert":
		// UserInformation 구조체 타입의 메모리 할당
		user := new(dto.UserDTO)
		// json 형태로 넘어온 파라미터 들을 user 변수에 바인딩
		err := c.Bind(user); if err != nil {
			fmt.Println(err)
			return c.String(http.StatusBadRequest, "bad request")
		}

		intResult := service.UserInsertService(user)

		result := ""
		if (intResult < 1) {
			result = "실패"
		} else {
			result = "성공"
		}

		return c.String(http.StatusOK, result)

	case "userUpdate":
		user := new(dto.UserDTO)
		err := c.Bind(user); if err != nil {
			fmt.Println(err)
			return c.String(http.StatusBadRequest, "bad request")
		}

		intResult := service.UserUpdateService(user)

		result := ""
		if (intResult < 1) {
			result = "업데이트 실패"
		} else {
			result = "업데이트 성공"
		}

		return c.String(http.StatusOK, result)

	case "userDelete":
		user := new(dto.UserDTO)
		err := c.Bind(user); if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		intResult := service.UserDeleteService(user)

		result := ""
		if (intResult < 1) {
			result = "삭제 실패"
		} else {
			result = "삭제 성공"
		}

		return c.String(http.StatusOK, result)
	}


	
	return c.String(http.StatusOK, "잘못된 접근입니다.")
}