package model

type TbUsers struct {
	Userid    uint   `json:"user_id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	RoleId    string `json:"role_id"`
	Branch    string `json:"branch"`
}
