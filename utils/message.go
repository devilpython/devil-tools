package utils

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

//显示错误消息
func ShowErrorMessage(context *gin.Context, message string) {
	context.Abort()
	context.JSON(http.StatusUnauthorized, gin.H{
		"successful": false,
		"message":    message,
	})
}

//显示错误消息
func ShowMessage(context *gin.Context, successful bool, message string) {
	context.JSON(200, gin.H{
		"successful": successful,
		"message":    message,
	})
}

//显示错误消息
func ShowIdMessage(context *gin.Context, successful bool, message string, id interface{}) {
	context.JSON(200, gin.H{
		"successful": successful,
		"message":    message,
		"id":         id,
	})
}

//显示数据消息
func ShowDataMessage(context *gin.Context, successful bool, message string, data interface{}) {
	context.JSON(200, gin.H{
		"successful": successful,
		"message":    message,
		"data":       data,
	})
}

//显示数据消息
func ShowQueryDataMessage(context *gin.Context, successful bool, message string, data interface{}, total int64) {
	context.JSON(200, gin.H{
		"successful": successful,
		"message":    message,
		"data":       data,
		"total":      total,
	})
}

var messageMap map[string]string
var once sync.Once

//获得消息
func GetMessage(messageId string) (string, bool) {
	once.Do(func() {
		loadMessageConfig("config/message.xml")
	})
	message, hasMessage := messageMap[messageId]
	if !hasMessage {
		message = "error message"
	}
	return message, hasMessage
}

//消息配置对象
type MessageConfig struct {
	XMLName      xml.Name  `xml:"message-config"` //顶层的消息配置名称
	MessageArray []Message `xml:"message"`        //消息数组
}

//账号模型
type Message struct {
	Id      string `xml:"id,attr"`      //消息标识
	Message string `xml:"message,attr"` //消息信息
}

//加载消息配置
func loadMessageConfig(path string) {
	messageMap = make(map[string]string)
	xconfig := MessageConfig{}
	success := LoadXmlObject(path, &xconfig)
	if success {
		for index := range xconfig.MessageArray {
			messageMap[xconfig.MessageArray[index].Id] = xconfig.MessageArray[index].Message
		}
	}
}
