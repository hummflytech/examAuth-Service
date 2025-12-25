package main

import (
	"net"

	"github.com/Dawit0/examAuth/internal/delivery/handler"
	"github.com/Dawit0/examAuth/internal/infrastructure/database"
	repo "github.com/Dawit0/examAuth/internal/infrastructure/repository/userRepo"
	"github.com/Dawit0/examAuth/internal/infrastructure/route"
	"github.com/Dawit0/examAuth/internal/pkg/logger"
	"github.com/Dawit0/examAuth/internal/server/middleware"
	"github.com/Dawit0/examAuth/internal/service"
	pb "github.com/Dawit0/examAuth/proto"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	_ "github.com/Dawit0/examAuth/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Auth Service API
// @version         1.0
// @description     This is the authentication service for the Exam application.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      ethioexam.hummflytech.com
// @BasePath  /auth/api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	logger.InitLogger()
	db := database.DBconnection()

	logger.Logger.Info("Database connected successfully")

	userRepos := repo.NewUserRepo(db)
	resetUserRepo := repo.NewResetUserRepo(db)

	usecase := service.NewUserService(userRepos)
	resetUserUseCase := service.NewResetUserService(resetUserRepo, service.NewGmailMailer("workenhdawit@gmail.com", "vlvs ygcl odpe gzee"))

	userHandler := handler.NewUserHandler(usecase)
	resetHandler := handler.NewForgetPasswordHandler(resetUserUseCase)

	routes := gin.New()

	routes.Use(
		middleware.LoggingMiddleware(logger.Logger),
		middleware.RecoveryMiddleware(logger.Logger),
	)

	route.UserRoute(userHandler, routes)
	route.ResetRoute(resetHandler, routes)

	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	logger.Logger.Info("Server started on :8080")

	go func() {
		routes.Run(":8080")
	}()

	grpc_server := grpc.NewServer()
	usecases := handler.NewGrpcHandler(usecase)
	pb.RegisterAuthServiceServer(grpc_server, usecases)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Logger.Error("failed to listen: %v", zap.Error(err))
		return
	}
	logger.Logger.Info("gRPC server started on :50051")
	go func() {
		if err := grpc_server.Serve(lis); err != nil {
			logger.Logger.Error("failed to serve: %v", zap.Error(err))
			return
		}
	}()

	select {}
}
