package model

//User 用户模型
type User struct {
	*Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

//TableName 获取用户表的表名
func (u User) TableName() string {
	return "user"
}
