package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) userIdentity(c *gin.Context) {
	user, ok := h.GetUser(c)
	if !ok {
		return
	}
	c.Set(userCtx, user.Id)
}
