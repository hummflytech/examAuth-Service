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

type AdminUserHandler struct {
	service *service.AdminUserService
}

func NewAdminUserHandler(service *service.AdminUserService) *AdminUserHandler {
	return &AdminUserHandler{service: service}
}

func (auh *AdminUserHandler) CreateAdminUser(c *gin.Context) {
	var adminuser dto.AdminUserCreate
	if err := c.ShouldBindJSON(&adminuser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isActive := true

	adminusers, err := domain.NewAdminUser(adminuser.Email, adminuser.Password, &isActive, adminuser.Username, adminuser.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	out, err := auh.service.CreateAdmins(adminusers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	val, err := mapper.MApAdminDomainToResponse(out)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, val)
}

func (auh *AdminUserHandler) FindById(c *gin.Context) {
	id := c.Param("id")

	out, err := auh.service.FindById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	val, err := mapper.MApAdminDomainToResponse(out)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, val)
}

func (auh *AdminUserHandler) FindAll(c *gin.Context) {
	out, err := auh.service.AllAdmins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var outs []dto.AdminUserResponse

	for _, item := range out {
		val, err := mapper.MApAdminDomainToResponse(&item)
		if err != nil {
			continue
		}
		outs = append(outs, *val)
	}

	c.JSON(http.StatusOK, gin.H{"data": outs})
}

func (auh *AdminUserHandler) UpdateAdminUser(c *gin.Context) {
	var adminuser dto.AdminUserCreate
	if err := c.ShouldBindJSON(&adminuser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isActive := true

	adminusers, err := domain.NewAdminUser(adminuser.Email, adminuser.Password, &isActive, adminuser.Username, adminuser.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	id := c.GetString("user_id")
	if id == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: missing user_id"})
		return
	}

	out, err := auh.service.UpdateAdmins(id, adminusers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	val, err := mapper.MApAdminDomainToResponse(out)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, val)
}

func (auh *AdminUserHandler) DeleteAdminUser(c *gin.Context) {
	id := c.Param("id")

	err := auh.service.DeleteAdmins(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Successfully Deleted"})
}

func (auh *AdminUserHandler) LoginAdminUser(c *gin.Context) {
	var adminuser dto.AdminUserLogin
	if err := c.ShouldBindJSON(&adminuser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	toke, err := auh.service.AdminLogin(adminuser.Email, adminuser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	val, err := security.GenerateToken(toke.Id())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": val})
}
