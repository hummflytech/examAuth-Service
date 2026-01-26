package handler

import (
	"net/http"

	"github.com/Dawit0/examAuth/internal/delivery/dto"
	"github.com/Dawit0/examAuth/internal/service"
	"github.com/gin-gonic/gin"
)

type AdminForgetHandler struct {
	service *service.AdminResetService
}

func NewAdminForgetHandler(service *service.AdminResetService) *AdminForgetHandler {
	return &AdminForgetHandler{service: service}
}

func (h *AdminForgetHandler) RequestResetPasswordEmail(c *gin.Context) {
	var request dto.RequestResetDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.RequestResetPasswordEmail(request.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "password reset email sent"})
}

func (h *AdminForgetHandler) ResetPassword(c *gin.Context) {
	var request dto.ResetPasswordDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.ResetPassword(request.Email, request.OTP, request.NewPassword); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "password reset successfully"})
}
