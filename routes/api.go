package routes

import (
	"golearning/app/http/controllers/api/v1/auth"
	controllers "golearning/app/http/controllers/api/v1"
	"golearning/app/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine){
	v1 := r.Group("/v1")
	v1.Use(middlewares.LimitIP("200-H"))
	{
		authGroup := v1.Group("/auth")
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			suc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)

			lgc := new(auth.LoginController)
			authGroup.POST("/signup/using-phone", lgc.LoginByPhone)
			authGroup.POST("/signup/using-email", lgc.LoginByPassword)

			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha" , vcc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone" ,   vcc.SendUsingPhone)
            authGroup.POST("/verify-codes/email" ,   vcc.SendUsingEmail)

			authGroup.POST("/login/refresh-token", lgc.RefreshToken)

			pwc := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-phone", pwc.ResetByPhone)
			authGroup.POST("/password-reset/using-email", pwc.ResetByEmail)

			
		}
		uc := new(controllers.UsersController)
		v1.GET("/user", middlewares.AuthJWT(),uc.CurrentUser)
        usersGroup := v1.Group("/users")
        {
            usersGroup.GET("", uc.Index)
        }

		cgc := new(controllers.CategoriesController)
        cgcGroup := v1.Group("/categories")
        {
			cgcGroup.GET("", cgc.Index)
            cgcGroup.POST("", middlewares.AuthJWT(), cgc.Store)
			cgcGroup.PUT("/:id", middlewares.AuthJWT(), cgc.Update)
			cgcGroup.DELETE("/:id", middlewares.AuthJWT(), cgc.Delete)
        }

		tpc := new(controllers.TopicsController)
        tpcGroup := v1.Group("/topics")
        {
			tpcGroup.GET("", tpc.Index)
            tpcGroup.POST("", middlewares.AuthJWT(), tpc.Store)
			tpcGroup.PUT("/:ID", middlewares.AuthJWT(), tpc.Update)
			tpcGroup.DELETE("/:id", middlewares.AuthJWT(), tpc.Delete)
			tpcGroup.GET("/:id", tpc.Show)
        }

		lsc := new(controllers.LinksController)
        linksGroup := v1.Group("/links")
        {
            linksGroup.GET("", lsc.Index)
        }
	}
	
}