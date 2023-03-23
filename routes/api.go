package routes

import (
	"github.com/gin-gonic/gin"
	controllers "gohub/app/http/controllers/api/v1"
	"gohub/app/http/controllers/api/v1/auth"
	"gohub/app/http/middlewares"
)

func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.Use(middlewares.LimitIP("200-H"))
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsEmailExist)
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), suc.SignupUsingPhone)
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), suc.SignupUsingEmail)

			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("50-H"), vcc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone)
			authGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("20-H"), vcc.SendUsingEmail)

			lgc := new(auth.LoginController)
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), lgc.LoginByPhone)
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), lgc.RefreshToken)

			pwc := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-phone", middlewares.GuestJWT(), pwc.ResetByPhone)
			authGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), pwc.ResetByEmail)
		}

		uc := new(controllers.UsersController)
		v1.GET("/user", middlewares.AuthJWT(), uc.CurrentUser)
		usersGroup := v1.Group("/users")
		{
			usersGroup.GET("", uc.Index)
		}
	}
}
