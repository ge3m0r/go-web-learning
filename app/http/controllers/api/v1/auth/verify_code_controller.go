package auth

import (
	v1 "golearning/app/http/controllers/api/v1"
	"golearning/pkg/captcha"
	"golearning/pkg/logger"
	"golearning/pkg/response"
	"golearning/app/requests"
	"golearning/pkg/verifycode"

	"github.com/gin-gonic/gin"
)

type VerifyCodeController struct{
	v1.BaseAPIController
}

func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context){
	id,b64s,_,err := captcha.NewCaptcha().GenerateCaptcha()
	logger.LogIf(err)

	response.JSON(c, gin.H{
		"captcha_id" :    id,
		"captcha_image":  b64s,
	})
}

// SendUsingPhone 发送手机验证码
func (vc *VerifyCodeController) SendUsingPhone(c *gin.Context) {

    // 1. 验证表单
    request := requests.VerifyCodePhoneRequest{}
    if ok := requests.Validate(c, &request, requests.VerifyCodePhone); !ok {
        return
    }

    // 2. 发送 SMS
    if ok := verifycode.NewVerifyCode().SendSMS(request.Phone); !ok {
        response.Abort500(c, "发送短信失败~")
    } else {
        response.Success(c)
    }
}

func (vc *VerifyCodeController) SendUsingEmail(c *gin.Context){
	request := requests.VerifyCodeEmailRequest{}

	if ok := requests.Validate(c, &request, requests.VerifyCodeEmail); !ok{
		return 
	}

	err := verifycode.NewVerifyCode().SendMail(request.Email)
	if err != nil{
		response.Abort500(c, "发送 Email 验证码失败～")
	}else {
		response.Success(c)
	}
}