/**
* Created by GoLand.
* User: link1st
* Date: 2019-07-25
* Time: 12:20
 */

package routers

import (
	"gowebsocket/controllers/home"
	"gowebsocket/controllers/systems"
	"gowebsocket/controllers/user"
	// "gowebsocket/middleware"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	router.LoadHTMLGlob("views/**/*")

	// 用户登陆
	router.POST("/login", user.Login)
	// 用户注册
	router.POST("/register", user.Register)
	// 忘记密码
	router.POST("/forgetPwd", user.ForgetPwd)

	// 中间件,拦截登陆
	// router.Use(middleware.CheckToken())

	// 用户组
	userRouter := router.Group("/user")
	userRouter.GET("/list", user.List)
	userRouter.GET("/online", user.Online)
	userRouter.POST("/sendMessage", user.SendMessage)
	userRouter.POST("/sendMessageAll", user.SendMessageAll)

	// 系统
	systemRouter := router.Group("/system")
	systemRouter.GET("/state", systems.Status)

	// home
	homeRouter := router.Group("/home")
	homeRouter.GET("/index", home.Index)

	// 上传图片、语音、视频

}
