package route

import (
	"github.com/Dawit0/examAuth/internal/delivery/handler"
	"github.com/Dawit0/examAuth/internal/delivery/http"
	"github.com/gin-gonic/gin"
)

func UserRoute(handler *handler.UserHandler, route *gin.Engine) {
	api := route.Group("/auth/api/v1")
	{
		api.POST("/create", handler.CreateUser)
		api.POST("/login", handler.LoginUser)
		api.GET("/user/:id", handler.FindByID)
		api.GET("/users", handler.AllUsers)
		api.DELETE("/delete/:id", handler.DeleteUser)
		api.PUT("/update", http.AuthMiddleware(), handler.UpdateUser)
	}
}
