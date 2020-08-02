package v1

import "github.com/gin-gonic/gin"

type User struct{}

//NewUser 获取新的user结构体
func NewUser() User {
	return User{}
}

//Create 创建新用户
func (u User) Create(c *gin.Context) {}

//Delete 删除指定用户
func (u User) Delete(c *gin.Context) {}

//Update 更新用户信息
func (u User) Update(c *gin.Context) {}

//Retrieve 获取指定用户信息
func (u User) Retrieve(c *gin.Context) {}
