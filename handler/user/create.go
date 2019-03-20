package user

import (
	// "fmt"
	// "net/http"

	"go-study/model"
	"go-study/pkg/errno"
	. "go-study/handler"

	"github.com/gin-gonic/gin"
	// "github.com/lexkong/log"
)

// Create creates a new user account.
func Create(c *gin.Context) {
	var r CreateRequest

	if err := c.Bind(&r); err != nil {
		// c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
		SendResp(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel {
		Username: r.Username,
		Password: r.Password,
	}

	if err := u.Validate(); err != nil {
		SendResp(c, errno.ErrValidation, CreateResponse {
			Username: r.Password,
		})
		return 
	}

	if err := u.Encrypt(); err != nil {
		SendResp(c, errno.ErrEncrypt, nil)
		return 
	}

	if err := u.Create(); err != nil {
		SendResp(c, errno.ErrDatabase, nil)
		return 
	}

	res := CreateResponse {
		Username: r.Username,
	}

	SendResp(c, nil, res)
}

func (r *CreateRequest) checkParmas() error {
	if r.Username == "" {
		return errno.New(errno.ErrValidation, nil).Add("username is empty.")
	}

	if r.Password == "" {
		return errno.New(errno.ErrValidation, nil).Add("password is empty.")
	}

	return nil
}