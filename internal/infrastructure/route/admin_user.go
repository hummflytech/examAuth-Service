package route

import (
	"github.com/Dawit0/examAuth/internal/delivery/handler"
	"github.com/Dawit0/examAuth/internal/delivery/http"
	"github.com/gin-gonic/gin"
)

func AdminUserRoute(hands *handler.AdminUserHandler, r *gin.Engine) {
	api := r.Group("/api/v1/admin")
	{
		api.POST("/login", hands.LoginAdminUser)
		api.POST("/create", hands.CreateAdminUser)
		api.GET("/all", hands.FindAll)
		api.GET("/:id", hands.FindById)
		api.PUT("/update", http.AuthMiddleware(), hands.UpdateAdminUser)
		api.DELETE("/:id", hands.DeleteAdminUser)
	}
}
