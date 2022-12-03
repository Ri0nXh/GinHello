package handler

import (
	"GinHello/db"
	"GinHello/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {
	/*
		通过PostForm 来解析一个个参数
		email := ctx.PostForm("email")
		username := ctx.PostForm(("username"))
		password := ctx.PostForm("password")
		confirm_password := ctx.PostForm("confirm_password")*/
	registerModel := model.UserRegister{}
	if err := ctx.ShouldBind(&registerModel); err != nil {
		log.Println("error -> ", err.Error())
		ctx.String(http.StatusInternalServerError, "传入参数失败")
		return
	}

	userId := db.SaveUser(&registerModel)

	//ctx.JSON(http.StatusOK, registerModel) 由此可见，gin可以直接renturn一个struct
	DataResp := model.UserResp{UserId: userId}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": &DataResp,
	})

}

func Login(ctx *gin.Context) {
	//var userLoginModel model.UserLogin 和下面这种方式实例化出来的时一样的，具体区别是什么呢？
	userLoginModel := model.UserLogin{}
	if err := ctx.ShouldBind(&userLoginModel); err != nil {
		log.Printf("Params invalid ,error :", err)
		return
	}
	userinfo, err := db.UserLogin(&userLoginModel)
	if err != nil {
		log.Println("Login failed!, error :", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": userinfo,
	})
}
