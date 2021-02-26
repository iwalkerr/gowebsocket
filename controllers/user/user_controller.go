/**
* Created by GoLand.
* User: link1st
* Date: 2019-07-25
* Time: 12:11
 */

package user

import (
	"fmt"
	"gowebsocket/common"
	"gowebsocket/controllers"
	"gowebsocket/lib/cache"
	"gowebsocket/lib/jwt"
	"gowebsocket/models"
	"gowebsocket/servers/websocket"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 用户登陆
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 验证用户名密码
	fmt.Println(username, password)
	uid := 1

	// 生成token
	token, err := jwt.GenerateToken(uid)
	if err != nil {
		c.JSON(200, common.Response(common.ServerError, "", nil))
		return
	}

	// 返回token和用户信息
	c.JSON(200, common.Response(common.OK, "", gin.H{
		"token":        token,
		"username":     username,
		"head_picture": "http://www.baidu.com",
		"uuid":         "uzr3234",
	}))
}

// 用户注册
func Register(c *gin.Context) {

}

// 忘记密码
func ForgetPwd(c *gin.Context) {

}

// 查看全部在线用户
func List(c *gin.Context) {

	appIdStr := c.Query("appId")
	appIdUint64, _ := strconv.ParseInt(appIdStr, 10, 32)
	appId := uint32(appIdUint64)

	fmt.Println("http_request 查看全部在线用户", appId)

	data := make(map[string]interface{})

	userList := websocket.UserList(appId)
	data["userList"] = userList
	data["userCount"] = len(userList)

	controllers.Response(c, common.OK, "", data)
}

// 查看用户是否在线
func Online(c *gin.Context) {

	userId := c.Query("userId")
	appIdStr := c.Query("appId")
	appIdUint64, _ := strconv.ParseInt(appIdStr, 10, 32)
	appId := uint32(appIdUint64)

	fmt.Println("http_request 查看用户是否在线", userId, appIdStr)

	data := make(map[string]interface{})

	online := websocket.CheckUserOnline(appId, userId)
	data["userId"] = userId
	data["online"] = online

	controllers.Response(c, common.OK, "", data)
}

// 给用户发送消息
func SendMessage(c *gin.Context) {
	// 获取参数
	appIdStr := c.PostForm("appId")
	userId := c.PostForm("userId")
	msgId := c.PostForm("msgId")
	message := c.PostForm("message")
	appIdUint64, _ := strconv.ParseInt(appIdStr, 10, 32)
	appId := uint32(appIdUint64)

	fmt.Println("http_request 给用户发送消息", appIdStr, userId, msgId, message)

	// TODO::进行用户权限认证，一般是客户端传入TOKEN，然后检验TOKEN是否合法，通过TOKEN解析出来用户ID
	// 本项目只是演示，所以直接过去客户端传入的用户ID(userId)

	data := make(map[string]interface{})

	if cache.SeqDuplicates(msgId) {
		fmt.Println("给用户发送消息 重复提交:", msgId)
		controllers.Response(c, common.OK, "", data)

		return
	}

	sendResults, err := websocket.SendUserMessage(appId, userId, msgId, message)
	if err != nil {
		data["sendResultsErr"] = err.Error()
	}

	data["sendResults"] = sendResults

	controllers.Response(c, common.OK, "", data)
}

// 给全员发送消息
func SendMessageAll(c *gin.Context) {
	// 获取参数
	appIdStr := c.PostForm("appId")
	userId := c.PostForm("userId")
	msgId := c.PostForm("msgId")
	message := c.PostForm("message")
	appIdUint64, _ := strconv.ParseInt(appIdStr, 10, 32)
	appId := uint32(appIdUint64)

	fmt.Println("http_request 给全体用户发送消息", appIdStr, userId, msgId, message)

	data := make(map[string]interface{})
	if cache.SeqDuplicates(msgId) {
		fmt.Println("给用户发送消息 重复提交:", msgId)
		controllers.Response(c, common.OK, "", data)

		return
	}

	sendResults, err := websocket.SendUserMessageAll(appId, userId, msgId, models.MessageCmdMsg, message)
	if err != nil {
		data["sendResultsErr"] = err.Error()

	}

	data["sendResults"] = sendResults

	controllers.Response(c, common.OK, "", data)

}
