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

// CreateUser godoc
// @Summary      Create a new user
// @Description  Create a new user with the input payload
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      dto.UserCreate  true  "User Create Data"
// @Success      200   {object}  dto.UserResponse
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /create [post]
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

// LoginUser godoc
// @Summary      Login a user
// @Description  Login a user and return a JWT token
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        login  body      dto.UserLogin  true  "User Login Data"
// @Success      200    {object}  map[string]string
// @Failure      400    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /login [post]
func (h *UserHandler) LoginUser(c *gin.Context) {
	var dtos dto.UserLogin
	if err := c.ShouldBindJSON(&dtos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.usecase.UserLogin(dtos.Email, dtos.Password)
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

// FindByID godoc
// @Summary      Get a user by ID
// @Description  Get details of a specific user by their ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  dto.UserResponse
// @Failure      500  {object}  map[string]string
// @Router       /user/{id} [get]
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

// AllUsers godoc
// @Summary      Get all users
// @Description  Get a list of all users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string][]dto.UserResponse
// @Failure      500  {object}  map[string]string
// @Router       /users [get]
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

// DeleteUser godoc
// @Summary      Delete a user
// @Description  Delete a user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /delete/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.usecase.DeleteUser(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}

// UpdateUser godoc
// @Summary      Update a user
// @Description  Update user details. Requires JWT token.
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        user  body      dto.UserCreate  true  "User Update Data"
// @Success      200   {object}  dto.UserResponse
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /update [put]
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
