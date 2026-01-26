package handler

import (
	"context"

	"github.com/Dawit0/examAuth/internal/domain"
	"github.com/Dawit0/examAuth/internal/service"
	pb "github.com/Dawit0/examAuth/proto"
)

type GrpcHandler struct {
	pb.UnimplementedAuthServiceServer
	usecase  *service.UserService
	ausecase *service.AdminUserService
}

func NewGrpcHandler(uc *service.UserService, auc *service.AdminUserService) *GrpcHandler {
	return &GrpcHandler{usecase: uc, ausecase: auc}
}

func (h *GrpcHandler) ValidateUser(ctx context.Context, req *pb.ValidateUserRequest) (*pb.ValidateUserResponse, error) {
	isValid, userID, err := h.usecase.ValidateToke(req.Token)
	if err != nil {
		return &pb.ValidateUserResponse{IsValid: false}, nil
	}

	return &pb.ValidateUserResponse{IsValid: isValid, UserId: userID}, nil
}

func (h *GrpcHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	domain, err := h.usecase.FindByID(req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		Id:       domain.ID(),
		Username: domain.Username(),
		Phone:    domain.Phone(),
		Email:    domain.Email(),
		Badge:    domain.Badge(),
		IsActive: domain.IsActive(),
		Score:    domain.Score(),
	}, nil
}

func (h *GrpcHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	val, err := domain.NewUser(req.Email, req.Password, &req.Badge, req.Username, req.Phone, &req.IsActive, &req.Score)
	if err != nil {
		return nil, err
	}

	val, err = h.usecase.UpdateUser(req.UserId, val)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserResponse{
		Id:       val.ID(),
		Username: val.Username(),
		Phone:    val.Phone(),
		Email:    val.Email(),
		Badge:    val.Badge(),
		IsActive: val.IsActive(),
		Score:    val.Score(),
		Password: val.Password(),
	}, nil
}

func (h *GrpcHandler) UpdateAdminUser(ctx context.Context, req *pb.UpdateAdminUserRequest) (*pb.UpdateAdminUserResponse, error) {
	if req.Password == "" {
		admin, err := h.ausecase.FindById(req.UserId)
		if err != nil {
			return nil, err
		}
		req.Password = admin.Password()
	}
	val, err := domain.NewAdminUser(req.Email, req.Password, &req.IsActive, req.Username, req.Phone)
	if err != nil {
		return nil, err
	}

	val, err = h.ausecase.UpdateAdmins(req.UserId, val)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateAdminUserResponse{
		Id:       val.Id(),
		Username: val.Username(),
		Phone:    val.Phone(),
		Email:    val.Email(),
		IsActive: val.IsActive(),
		Password: val.Password(),
	}, nil
}

func (h *GrpcHandler) GetAdminUser(ctx context.Context, req *pb.GetAdminUserRequest) (*pb.GetAdminUserResponse, error) {
	domain, err := h.ausecase.FindById(req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.GetAdminUserResponse{
		Id:       domain.Id(),
		Username: domain.Username(),
		Phone:    domain.Phone(),
		Email:    domain.Email(),
		IsActive: domain.IsActive(),
	}, nil
}
