package resource

var DB = DBInfo {"abc", "123", "localhost:3306"	/*docker의 경우 변경해주어야 함*/, "mysql", "tree"}

type DBInfo struct {
	user 		string
	pwd 		string
	url 		string
	engine 		string
	database 	string
}
