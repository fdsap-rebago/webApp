package model

type User struct {
	Userid   uint   `json:"userid" bson:"userid"`
	Fullname string `json:"fullname" bson:"fullname"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
