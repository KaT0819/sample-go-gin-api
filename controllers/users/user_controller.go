package users

import (
	"fmt"
	"net/http"
	"sample-go-gin-api/domain/users"
	"sample-go-gin-api/services"
	"sample-go-gin-api/utils/errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)

}

func Get(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("ユーザIDは数値で入力してください")
		c.JSON(err.Status, err)
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func Update(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("ユーザIDは数値で入力してください")
		c.JSON(err.Status, err)
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	isPertial := c.Request.Method == http.MethodPatch
	user.Id = userId

	result, err := services.UpdateUser(user, isPertial)
	fmt.Println(result, err)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("ユーザIDは数値で入力してください")
		c.JSON(err.Status, err)
	}

	err := services.DeleteUser(userId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users)
}
