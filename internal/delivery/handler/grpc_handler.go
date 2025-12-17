package handler

import (
	"context"
	"time"

	"github.com/Dawit0/examAuth/internal/domain"
	"github.com/Dawit0/examAuth/internal/service"
	pb "github.com/Dawit0/examAuth/proto"
)

type GrpcHandler struct {
	usecase *service.UserService
	pb.UnimplementedAuthServiceServer
}

func NewGrpcHandler(uc *service.UserService) *GrpcHandler {
	return &GrpcHandler{usecase: uc}
}

func (h *GrpcHandler) ValidateUser(ctx context.Context, req *pb.ValidateUserRequest) (*pb.ValidateUserResponse, error) {
	valid, id, err := h.usecase.ValidateToke(req.Token)
	if err != nil {
		return &pb.ValidateUserResponse{
			IsValid: valid,
			UserId:  int64(id),
		}, err
	}
	return &pb.ValidateUserResponse{
		IsValid: valid,
		UserId:  int64(id),
	}, nil
}

func (h *GrpcHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	domain, err := h.usecase.FindByID(uint(req.UserId))

	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		Id:       int64(domain.ID()),
		Email:    domain.Email(),
		Badge:    domain.Badge(),
		Username: domain.Username(),
		Phone:    domain.Phone(),
		IsActive: domain.IsActive(),
		Score:    domain.Score(),
	}, nil

}

func (h *GrpcHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	domain_out, err := domain.WithoutValidation(req.Email, req.Password, req.Badge, req.Username, req.Phone, req.IsActive, req.Score, time.Now())
	if err != nil {
		return nil, err
	}

	val, err := h.usecase.UpdateUser(uint(req.UserId), domain_out)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserResponse{
		Id:       int64(val.ID()),
		Username: val.Username(),
		Phone:    val.Phone(),
		Email:    val.Email(),
		Badge:    val.Badge(),
		IsActive: val.IsActive(),
		Score:    val.Score(),
	}, nil
}
