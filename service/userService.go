package service

import (
	"fmt"
	"strconv"
	"reflect"
	"strings"

	"tree/resource"
	"tree/dto"
)

func ReceiveService(amode string) string {
	switch amode {
	case "login" : 
		return "Hi! This is Login Page"
	}
	return "잘못된 접근!!"
}

func UserSelectListService(page string) []byte {
	intPage, _ := strconv.Atoi(page)
	if (intPage < 1) {
		intPage = 1
	}
	offset := (intPage - 1) * 10;
	query := "SELECT * FROM user_information LIMIT 10 OFFSET " + strconv.Itoa(offset) + ";"
	return resource.SelectListQuery(resource.DB, query)
}

func UserInsertService(user *dto.UserDTO) int64 {
	query := "INSERT INTO user_information ("
	query += "id, pw, dkey, phone_num, birth, "
	query += "gender, sns_type, nick_name, level, "
	query += "regist_date, pw_change_date, "
	query += "final_connect_date, user_status, user_profile) "
	query += "VALUES ('"
	query += user.Id + "', '"
	query += user.Pw + "', '"
	query += user.Dkey + "', '"
	query += user.PhoneNum + "', '"
	query += user.Birth + "', '"
	query += user.Gender + "', '"
	query += user.SnsType + "', '"
	query += user.NickName + "', '"
	query += user.Level + "', '"
	query += user.RegistDate + "', '"
	query += user.PwChangeDate + "', '"
	query += user.FinalConnectDate + "', '"
	query += user.UserStatus + "', '"
	query += user.UserProfile + "')"

	return resource.InsertQuery(resource.DB, query)
}

func UserUpdateService(user *dto.UserDTO) int64 {
	query := "UPDATE user_information SET "
	
	//구조체로 선언된 UserDTO 를 reflect 하여 값을 가져옴
	v := reflect.ValueOf(user)
	//값으로 선언된 v의 요소에 접근
	elements := v.Elem()

	var cnt = 0
	var sno = ""

	//접근한 요소의 필드 수 만큼 반복 -> 구조체에서 선언한 변수만큼 반복
	for i := 0; i < elements.NumField(); i++ {
		//구조체에서 i 번째로 선언한 변수의 태그를 가져옴 (`` 사이에 선언된 값)
		structTag := elements.Type().Field(i).Tag
		//구조체에서 i 번째로 선언한 변수의 값을 가져옴
		structValue := elements.Field(i)
		
		//가져온 태그를 문자열로 변환
		stringTag := fmt.Sprintf("%v", structTag)
		//문자열로 변환된 태그에서 컬럼명만 추출
		column := stringTag[6:6 + strings.Index(stringTag[6:len(stringTag)], "\"")]
		//가져온 변수의 값을 문자열로 변환
		value := fmt.Sprintf("%v", structValue)

		//컬럼명만 추출했을 때 태그 설정에 타입이 포함된 경우
		if (strings.Contains(column, ",")) {
			//타입을 제외하고 컬럼 명만 추출
			column = column[:strings.Index(column[:len(column)], ",")]
		}

		if (column == "sno") {
			sno = value
		}

		if value != "0" || value != "" {
			if cnt != 0 {
				query += ", "
			}

			query += column + " = '" + value + "'"

			cnt++
		}
	}

	query += " WHERE sno = '" + sno + "';"
	
	if sno == "0" || sno == "" {
		return 0
	} else {
		return resource.UpdateQuery(resource.DB, query);
	}
}

func UserDeleteService(user *dto.UserDTO) int64 {
	query := "DELETE FROM user_information WHERE sno = '" + strconv.Itoa(user.Sno) + "';"

	return resource.DeleteQuery(resource.DB, query)
}