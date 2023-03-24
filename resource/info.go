package resource

var DB = DBInfo {"abc", "123", "localhost:3306"	/*docker의 경우 변경해주어야 함*/, "mysql", "tree"}

type DBInfo struct {
	user 		string
	pwd 		string
	url 		string
	engine 		string
	database 	string
}

type UserInformation struct {
	Sno 				int				`json:"sno" form:"sno" query:"sno"`
	Id 					string			`json:"id" form:"id" query:"id"`
	Pw 					string			`json:"pw" form:"pw" query:"pw"`
	Dkey 				string			`json:"dkey" form:"dkey" query:"dkey"`
	PhoneNum 			string			`json:"phone_num" form:"phone_num" query:"phone_num"`
	Birth 				string			`json:"birth" form:"birth" query:"birth"`
	Gender 				string			`json:"gender" form:"gender" query:"gender"`
	SnsType				string			`json:"sns_type" form:"sns_type" query:"sns_type"`
	NickName			string			`json:"nick_name" form:"nick_name" query:"nick_name"`
	Level 				string			`json:"level" form:"level" query:"level"`
	RegistDate 			string			`json:"regist_date" form:"regist_date" query:"regist_date"`
	PwChangeDate 		string			`json:"pw_change_date" form:"pw_change_date" query:"pw_change_date"`
	FinalConnectDate 	string			`json:"final_connect_date" form :"final_connect_date" query:"final_connect_date"`
	UserStatus 			string			`json:"user_status" form:"user_status" query:"user_status"`
	UserProfile 		string			`json:"user_profile" form:"user_profile" query:"user_profile"`
}
