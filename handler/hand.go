package handler

import (
	"net/http"

	"go-study/pkg/errno"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int `json:"code"`
	Message string `json:"msg"`
	Data interface{} `json:"data"`
}

func SendResp(c *gin.Context, err error, data interface{}) {
	code, msg := errno.DecodeErr(err)

	c.JSON(http.StatusOK, Response {
		Code: code,
		Message: msg,
		Data: data,
	})
}