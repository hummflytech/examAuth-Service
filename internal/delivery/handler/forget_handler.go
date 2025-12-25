package handler

import (
	"net/http"

	"github.com/Dawit0/examAuth/internal/delivery/dto"
	"github.com/Dawit0/examAuth/internal/service"
	"github.com/gin-gonic/gin"
)

type ForgetPasswordHandler struct {
	resetUserService *service.ResetUserService
}

func NewForgetPasswordHandler(resetUserService *service.ResetUserService) *ForgetPasswordHandler {
	return &ForgetPasswordHandler{resetUserService: resetUserService}
}

// RequestResetPasswordEmail godoc
// @Summary      Request password reset email
// @Description  Send an OTP to the user's email for password reset
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      dto.RequestResetDTO  true  "Reset Request Data"
// @Success      200      {object}  map[string]string
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /reset [post]
func (h *ForgetPasswordHandler) RequestResetPasswordEmail(c *gin.Context) {
	var req dto.RequestResetDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.resetUserService.RequestResetPasswordEmail(req.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "password reset email sent"})
}

// ResetPassword godoc
// @Summary      Reset password
// @Description  Reset the user's password using the OTP
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        reset  body      dto.ResetPasswordDTO  true  "Reset Password Data"
// @Success      200    {object}  map[string]string
// @Failure      400    {object}  map[string]string
// @Failure      401    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /reset-password [post]
func (h *ForgetPasswordHandler) ResetPassword(c *gin.Context) {
	var req dto.ResetPasswordDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.resetUserService.ResetPassword(req.Email, req.OTP, req.NewPassword); err != nil {
		switch err.Error() {
		case "password reset expired":
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "OTP expired",
			})
		case "invalid credentials", "record not found":
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid OTP",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to reset password",
			})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "password reset successfully"})
}
