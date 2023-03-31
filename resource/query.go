package resource

import (
	"database/sql"
	"encoding/json"

	"log"
	"tree/util"
)

var fileLog = new(util.File)

func DBConnection(db DBInfo) *sql.DB {
	fileLog.Path = "./logs/SQL"
	fileLog.Name = "DBAccess"

	f := util.Logging(fileLog, "DBConnect", "", nil)

	defer f.Close()

	dataSource := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database

	conn, err := sql.Open(db.engine, dataSource)

	if err != nil {
		errorLog := log.New(f, "[ERROR] : ", log.LstdFlags)
		errorLog.Print(" DBConnect Error!!")
		errorLog.Println(err)
	}

	return conn
}

// INSERT, SELECT, UPDATE, DELETE 모두 커넥션 하는 구조는 동일하지만 여러 개의 함수로 나눈 이유는 서비스 코드에서 호출할 때 어떤 용도로 사용하는지 확인하기 위해서 사용

// INSERT 쿼리문
func InsertQuery(db DBInfo, query string) int64 {
	conn := DBConnection(db)
	defer conn.Close()

	// 연결된 DB에 쿼리 전송
	result, err := conn.Exec(query)

	fileLog.Path = "./logs/SQL"
	fileLog.Name = "Query"

	f := util.Logging(fileLog, "Insert Query", query, nil)

	if err != nil {
		errorLog := log.New(f, "[ERROR] : ", log.LstdFlags)
		errorLog.Print(" Insert Query Error!! -> " + query)
		errorLog.Println(err)
	}

	defer f.Close()

	nRow, err := result.RowsAffected()

	// 실행 결과 리턴
	return nRow
}

// SELECT 쿼리문
func SelectListQuery(db DBInfo, query string) []byte {
	conn := DBConnection(db)
	defer conn.Close()

	fileLog.Path = "./logs/SQL"
	fileLog.Name = "Query"

	f := util.Logging(fileLog, "Select Query", query, nil)

	defer f.Close()
	
	// 연결된 DB에 쿼리 전송
	rows, err := conn.Query(query)
	defer rows.Close()

	if err != nil {
		errorLog := log.New(f, "[ERROR] : ", log.LstdFlags)
		errorLog.Print(" Select Query Error!!")
		errorLog.Println(err)
	}

	columns, _ := rows.Columns()
	// 결과 컬럼의 개수 저장
	count := len(columns)

	var v struct {
		Data []interface{}	//`json : "UserInformation"`
	}

	for rows.Next() {
		values := make([]interface{}, count)
		valuePtrs := make([]interface{}, count)

		// interface 는 타입 문제로 할당이 어려워 포인터를 사용
		for i, _ := range columns {
			valuePtrs[i] = &values[i]
		}

		// rows.Scan(&id, &pw) 와 같이 직접 지정도 가능하나 변수가 많을 수록 불편하기에 가변인자 를 사용
		if err := rows.Scan(valuePtrs...); err != nil {	//Scan parameter same to form the &test 
			log.Fatal(err)
		}

		var m map[string]interface{}
		m = make(map[string]interface{})

		// 해당 컬럼에 해당 값을 string 타입으로 할당
		for i := range columns {
			m[columns[i]] = string(values[i].([]byte))
		}
		v.Data = append(v.Data, m)
	}

	// v 구조체로 마샬링(인코딩) 하여 전달
	jsonMsg, err := json.Marshal(v)
	return jsonMsg
}

// UPDATE 쿼리문
func UpdateQuery(db DBInfo, query string) int64 {
	conn := DBConnection(db)
	defer conn.Close()

	fileLog.Path = "./logs/SQL"
	fileLog.Name = "Query"

	f := util.Logging(fileLog, "Update Query", query, nil)

	defer f.Close()

	// 연결된 DB에 쿼리 전송
	result, err := conn.Exec(query)

	if err != nil {
		errorLog := log.New(f, "[ERROR] : ", log.LstdFlags)
		errorLog.Print(" Update Query Error!!")
		errorLog.Println(err)
	}

	nRow, _ := result.RowsAffected()

	return nRow
}

// DELETE 쿼리문
func DeleteQuery(db DBInfo, query string) int64 {
	conn := DBConnection(db)
	defer conn.Close()

	fileLog.Path = "./logs/SQL"
	fileLog.Name = "Query"

	f := util.Logging(fileLog, "Delete Query", query, nil)

	defer f.Close()

	// 연결된 DB에 쿼리 전송
	result, err := conn.Exec(query)

	if err != nil {
		errorLog := log.New(f, "[ERROR] : ", log.LstdFlags)
		errorLog.Print(" Delete Query Error!!")
		errorLog.Println(err)
	}

	nRow, _ := result.RowsAffected()

	return nRow
}