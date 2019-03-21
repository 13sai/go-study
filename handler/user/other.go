package user

import (
	// "fmt"
	"strconv"
	// "net/http"

	"go-study/util"
	"go-study/model"
	"go-study/service"
	"go-study/pkg/errno"
	. "go-study/handler"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// Create creates a new user account.
func View(c *gin.Context) {
	username := c.Param("username")
	// Get the user by the `username` from the database.
	user, err := model.GetUser(username)
	if err != nil {
		SendResp(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResp(c, nil, user)
}

// Delete delete an user by the user identifier.
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResp(c, errno.ErrDatabase, nil)
		return
	}

	SendResp(c, nil, nil)
}


// Create creates a new user account.
func Update(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	// Get the user id from the url parameter.
	userId, _ := strconv.Atoi(c.Param("id"))

	// Binding the user data.
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResp(c, errno.ErrBind, nil)
		return
	}

	// We update the record based on the user id.
	u.Id = uint64(userId)

	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResp(c, errno.ErrValidation, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResp(c, errno.ErrEncrypt, nil)
		return
	}

	// Save changed fields.
	if err := u.Update(); err != nil {
		SendResp(c, errno.ErrDatabase, nil)
		return
	}

	SendResp(c, nil, nil)
}


// Create creates a new user account.
func Index(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		SendResp(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		SendResp(c, err, nil)
		return
	}

	SendResp(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}