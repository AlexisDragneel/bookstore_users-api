package controllers

import (
	"github.com/alexisdragneel/bookstore_users-api/domain/users"
	"github.com/alexisdragneel/bookstore_users-api/services"
	"github.com/alexisdragneel/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context){
	c.JSON(http.StatusNotImplemented, "not implemented")
}

func FindUser(c *gin.Context) {
	var restErr *errors.RestError
	id, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		restErr = errors.NewRestError("Invalid user id", http.StatusBadRequest)
		c.JSON(restErr.Status, restErr)
		return
	}
	res, err := services.FindUser(id)
	if err != nil {
		restErr = errors.NewRestError(err.Error(), http.StatusInternalServerError)
		c.JSON(restErr.Status, restErr)
	}
	c.JSON(http.StatusFound, res)
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		httpError := errors.NewRestError("invalid json body", http.StatusBadRequest)
		c.JSON(httpError.Status, httpError)
		return
	}
	result, err := services.CreateUser(user)
	if err != nil {
		httpError := errors.NewRestError(err.Error(),http.StatusInternalServerError)
		c.JSON(httpError.Status, httpError)
		return
	}
	c.JSON(http.StatusCreated, result)
}