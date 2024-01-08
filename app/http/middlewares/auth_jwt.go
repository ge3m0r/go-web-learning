package middlewares

import(
	"fmt"
	"golearning/app/models/user"
	"golearning/pkg/config"
	"golearning/pkg/jwt"
	"golearning/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc{
	return func(c *gin.Context){
		claims, err := jwt.NewJWT().ParserToken(c)
		if err != nil{
			response.Unauthorized(c, fmt.Sprintf("请查看 %v 相关的接口认证文档", config.GetString("app.name")))
			return 
		}
		userModel := user.Get(claims.UserID)
		if userModel.ID == 0 {
			response.Unauthorized(c, "找不到对应用户，用户可能已经删除")
			return
		}
		c.Set("current_user_id", userModel.GetStringID())
        c.Set("current_user_name", userModel.Name)
        c.Set("current_user", userModel)

        c.Next()
	}
}