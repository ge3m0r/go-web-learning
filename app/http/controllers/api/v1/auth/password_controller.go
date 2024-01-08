package auth

import(
	v1 "golearning/app/http/controllers/api/v1"
	"golearning/app/models/user"
	"golearning/app/requests"
	"golearning/pkg/response"

	"github.com/gin-gonic/gin"

)

type PasswordController struct{
	v1.BaseAPIController
}

func (pc *PasswordController)ResetByPhone(c *gin.Context){
	request := requests.ResetByPhoneRequest{}
	if ok := requests.Validate(c , &request, requests.ResetByPhone); !ok{
		return 
	}
	userModel := user.GetByPhone(request.Phone)
    if userModel.ID == 0 {
        response.Abort404(c)
    } else {
        userModel.Password = request.Password
        userModel.Save()

        response.Success(c)
    }
}

func (pc *PasswordController) ResetByEmail(c *gin.Context){
	// 1. 验证表单
    request := requests.ResetByEmailRequest{}
    if ok := requests.Validate(c, &request, requests.ResetByEmail); !ok {
        return
    }

    // 2. 更新密码
    userModel := user.GetByEmail(request.Email)
    if userModel.ID == 0 {
        response.Abort404(c)
    } else {
        userModel.Password = request.Password
        userModel.Save()
        response.Success(c)
    }
}