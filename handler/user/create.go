package user

import (
	"fmt"
	// "net/http"

	"go-study/pkg/errno"
	. "go-study/handler"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// Create creates a new user account.
func Create(c *gin.Context) {
	// var r struct {
	// 	Username string `json:"username"`
	// 	Password string `json:"password"`
	// }

	var r CreateRequest

	// var err error
	if err := c.Bind(&r); err != nil {
		// c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
		SendResp(c, errno.ErrBind, nil)
		return
	}

	admin := c.Param("username")
	// pw := c.PostForm("password")

	// r = CreateRequest {
	// 	Username: admin,
	// 	Password: pw,
	// }

	log.Infof("URL username: %s", admin)

	desc := c.Query("desc")
	log.Infof("URL params desc: %s", desc)

	// log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)
	if r.Username == "" {
		SendResp(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username is empty")), nil)
		return 
	}

	if r.Password == "" {
		SendResp(c, fmt.Errorf("password is empty"), nil)
		return
	}

	res := CreateResponse {
		Username: admin,
	}
	

	SendResp(c, nil, res)

	// code, message := errno.DecodeErr(err)
	// c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}