package route

import (
	"github.com/Dawit0/examAuth/internal/delivery/handler"
	"github.com/gin-gonic/gin"
)


func ResetAdminPasswordRoute(hands *handler.AdminForgetHandler, r *gin.Engine) {	  
	api := r.Group("/api/v1/reset-admin")
	api.POST("/reset-password", hands.ResetPassword)
	api.POST("/request-reset-password", hands.RequestResetPasswordEmail)
}