package model

type TbUsers struct {
	Uid       uint   `json:"uid" gorm:"primaryKey"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username" gorm:"unique"`
	Password  string `json:"password"`
	InstiCode int    `json:"insti_code"`
}

type M_Institution struct {
	InstiCode uint   `json:"insti_code"`
	InstiDesc string `json:"insti_desc"`
}
