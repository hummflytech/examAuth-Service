package handler

import (
	"net/http"
	"strconv"

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

func (h *UserHandler) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	domain, err := h.usecase.FindByID(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response, err := mapper.MapDomaintoResponse(*domain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) AllUsers(c *gin.Context) {
	out, err := h.usecase.AllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := make([]dto.UserResponse, 0, len(out))

	for _, item := range out {
		val, err := mapper.MapDomaintoResponse(item)
		if err != nil {
			continue
		}

		response = append(response, *val)
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.usecase.DeleteUser(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var dtos dto.UserCreate

	if err := c.ShouldBindJSON(&dtos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	val, err := domain.NewUser(dtos.Email, dtos.Password, dtos.Badge, dtos.Username, dtos.Phone, dtos.IsActive, dtos.Score)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	id := c.GetUint("user_id")
	user, err := h.usecase.UpdateUser(uint(id), val)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.UserResponse{
		ID:       id,
		Email:    user.Email(),
		Badge:    user.Badge(),
		Username: user.Username(),
		Phone:    user.Phone(),
		IsActive: user.IsActive(),
		Score:    user.Score(),
	})

}
