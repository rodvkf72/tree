package util

import (
	"time"
	"os"
	"fmt"
	"log"
)

type File struct {
	Path 			string
	Name 			string
	PathAndName 	string
}

// 폴더 존재여부 확인 (없으면 생성)
func FolderCheck(folderPath string) {
	if _, err := os.Stat(folderPath); err != nil {
		os.MkdirAll(folderPath, os.ModePerm)
	}
}

// 로그파일명 정의 (ex. 2023-03-01_MainAccess_log.txt)
func LogFileName(fileName string) string {
	// 현재 시간
	now := time.Now()
	// 날짜 포맷방식 지정
	custom := now.Format("2006-01-02")
	// 로그파일명 지정
	logFileName := custom + "_" + fileName + "_log.txt"

	return logFileName
}

// 파일 정보(Path, Name) 을 구조체로 받아 전체경로(All) 설정
func MakeFile(st *File) {
	FolderCheck(st.Path)
	st.PathAndName = st.Path + "/" + LogFileName(st.Name)
}

// 로그 처리
func Logging(st *File, msg string, query string, err error) *os.File {
	MakeFile(st)

	// 파일 열기 (없으면 생성)
	f, file := os.OpenFile(st.PathAndName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	// 파일을 열 때 에러가 있다면
	if file != nil {
		panic(fmt.Sprintf("error opening file : %v", file))
	}

	if err != nil {
		errorLog := log.New(f, "[ERROR] : ", log.LstdFlags)
		errorLog.Print(" " + query)
		errorLog.Println(err)
	} else {
		infoLog := log.New(f, "[INFO] : ", log.LstdFlags)
		infoLog.Println(" " + msg)
		if query != "" {
			infoLog.Println(" " + query)
		}
	}

	return f
}