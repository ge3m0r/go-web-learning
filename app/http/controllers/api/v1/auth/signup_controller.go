package auth

import (
	v1 "golearning/app/http/controllers/api/v1"
	"golearning/app/models/user"
	"golearning/app/requests"
	"golearning/pkg/response"
	"github.com/gin-gonic/gin"
	"golearning/pkg/jwt"
)

type SignupController struct{
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context){
	request := requests.SignupPhoneExistRequest{}

	if ok := requests.Validate(c,&request,requests.ValidateSignupPhoneExist); !ok{
		return
	}

	response.JSON(c, gin.H{
		"exist" : user.IsPhoneExist(request.Phone),
	})
}

func (sc *SignupController) IsEmailExist(c *gin.Context){
	request := requests.SingupEmailExistRequest{}

	if ok := requests.Validate(c,&request,requests.ValidateSignupEmailExist); !ok{
		return
	}

	response.JSON(c, gin.H{
		"exist" : user.IsEmailExist(request.Email),
	})
}

// SignupUsingPhone 使用手机和验证码进行注册
func (sc *SignupController) SignupUsingPhone(c *gin.Context) {

    // 1. 验证表单
    request := requests.SignupUsingPhoneRequest{}
    if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
        return
    }

    // 2. 验证成功，创建数据
    userModel := user.User{
        Name:     request.Name,
        Phone:    request.Phone,
        Password: request.Password,
    }
    userModel.Create()

    if userModel.ID > 0 {
        token := jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Name)
        response.CreatedJSON(c, gin.H{
            "token": token,
            "data":  userModel,
        })
    } else {
        response.Abort500(c, "创建用户失败，请稍后尝试~")
    }
}

// SignupUsingEmail 使用 Email + 验证码进行注册
func (sc *SignupController) SignupUsingEmail(c *gin.Context) {

    // 1. 验证表单
    request := requests.SignupUsingEmailRequest{}
    if ok := requests.Validate(c, &request, requests.SignupUsingEmail); !ok {
        return
    }

    // 2. 验证成功，创建数据
    userModel := user.User{
        Name:     request.Name,
        Email:    request.Email,
        Password: request.Password,
    }
    userModel.Create()

    if userModel.ID > 0 {
        token := jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Name)
        response.CreatedJSON(c, gin.H{
            "token": token,
            "data":  userModel,
        })
    } else {
        response.Abort500(c, "创建用户失败，请稍后尝试~")
    }
}