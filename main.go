package main

import (
	"errors"
	// "log"
	"net/http"
	"time"
	"fmt"

	"go-study/router"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.New()

	middleW := []gin.HandlerFunc{}

	router.Load(
		g,
		middleW...,
	)

	go func() {
		err := pingSer();

		if err != nil {
			fmt.Println("oh no")
		}

		fmt.Println("success")
	}()

	http.ListenAndServe(":8077", g)
}

func pingSer() error {
	for i := 0; i < 2; i++ {
		res, err := http.Get("http:127.0.0.1:8077" + "/sd/health")

		if err == nil && res.StatusCode == 200 {
			return nil
		}

		time.Sleep(time.Second)
	}

	return errors.New("bad route")
}