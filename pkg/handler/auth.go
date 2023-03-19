package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if input.Username != "dima" || input.Password != "123" {
		newErrorResponse(c, http.StatusUnauthorized, "wrong login or password")
		return
	}
	response := "success"
	c.JSON(http.StatusOK, map[string]interface{}{
		"response": response,
	})
}

func (h *Handler) signUp(c *gin.Context) {

}
