// Package requests 处理请求数据和表单验证
package requests

import (
	"golearning/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ValidatorFunc func(interface{}, *gin.Context)map[string][]string

func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool{

	if err := c.ShouldBindJSON(obj); err != nil{
		response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式")
		return false
	}
	errs := handler(obj,c)

	if len(errs) > 0{
		response.ValidationError(c, errs)
		return false
	}
	return true

}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {

    // 配置选项
    opts := govalidator.Options{
        Data:          data,
        Rules:         rules,
        TagIdentifier: "valid", // 模型中的 Struct 标签标识符
        Messages:      messages,
    }

    // 开始验证
    return govalidator.New(opts).ValidateStruct()
}