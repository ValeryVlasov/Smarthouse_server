package handler

import (
	"fmt"
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input Smarthouse_server.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	/*	var input signInInput

		if err := c.BindJSON(&input); err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	*/
	tokenString := c.GetHeader(authorizationHeader)
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	fmt.Println(tokenString)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(claims["username"], claims["password"])

	//Сравнить с данными из бд
	//Сравнение логина и пароля
	user, ok := h.services.Authorization.IsSameUser(claims["username"], claims["password"])
	fmt.Println("ok = " + cast.ToString(ok))
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "incorrect login or password")
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"name": user.Name,
	})
}

func (h *Handler) GetUser(c *gin.Context) (Smarthouse_server.User2, bool) {
	var user Smarthouse_server.User2
	tokenString := c.GetHeader(authorizationHeader)
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	fmt.Println(tokenString)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return user, false
	}

	fmt.Println(claims["username"], claims["password"])

	//Сравнить с данными из бд
	//Сравнение логина и пароля
	user, ok := h.services.Authorization.IsSameUser(claims["username"], claims["password"])
	fmt.Println("ok = " + cast.ToString(ok))
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "incorrect login or password")
		return user, false
	}
	return user, true
}
