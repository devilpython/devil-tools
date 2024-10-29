package middle_ware

import (
	"io"

	"github.com/devilpython/devil-tools/utils"
	"github.com/gin-gonic/gin"
)

var KeyPostData = 0

// 数据绑定中间件
func PostDataBinding() gin.HandlerFunc {
	return func(context *gin.Context) {
		//防止数据溢出
		defer utils.RemoveAllGlobalData()
		data, _ := io.ReadAll(context.Request.Body)
		param := make(map[string]interface{})
		paramObj, _ := utils.BindingData(string(data), param)
		dataMap, ok := paramObj.(map[string]interface{})
		if !ok {
			dataMap = param
		}
		err := context.Request.ParseForm()
		if err == nil {
			for key, value := range context.Request.PostForm {
				dataMap[key] = value
			}
		}
		for key, value := range context.Request.URL.Query() {
			if len(value) == 1 {
				dataMap[key] = value[0]
			} else {
				dataMap[key] = value
			}
		}
		utils.SetGlobalData(KeyPostData, dataMap)
		context.Next()
	}
}
