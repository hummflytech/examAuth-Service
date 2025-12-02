package handler

import (
	"net/http"

	"github.com/Dawit0/examAuth/internal/delivery/dto"
	"github.com/Dawit0/examAuth/internal/delivery/mapper"
	"github.com/Dawit0/examAuth/internal/domain"
	"github.com/Dawit0/examAuth/internal/infrastructure/security"
	"github.com/Dawit0/examAuth/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase *service.UserService
}

func NewUserHandler(uc *service.UserService) *UserHandler {
	return &UserHandler{usecase: uc}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var dtos dto.UserCreate
	if err := c.ShouldBindJSON(&dtos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domain, errs := domain.NewUser(dtos.Email, dtos.Password, dtos.Badge, dtos.Username, dtos.Phone, dtos.IsActive, dtos.Score)

	if errs != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errs.Error()})
		return
	}

	val, err := h.usecase.CreateUser(domain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := mapper.MapDomaintoResponse(*val)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var dtos dto.UserLogin
	if err := c.ShouldBindJSON(&dtos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.usecase.UserLogin(dtos.Phone, dtos.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := security.GenerateToken(user.ID())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
