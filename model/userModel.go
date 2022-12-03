package model

type UserRegister struct {
	Email           string `form:"email" binding:"required,email"`
	Username        string `form:"username" binding:"required"`
	Password        string `form:"password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" binding:"required,eqfield=Password"`
}

// UserLogin 此处存在缺陷，此时希望用户既可以使用邮箱登录也可以使用用户名登录，如果添加了binding，则变成强制传参了。
type UserLogin struct {
	Email    string `form:"email"`
	Username string `form:"username"`
	Password string `form:"password" binding:"required"`
}

type UserInfo struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
