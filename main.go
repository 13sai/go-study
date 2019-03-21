package main

import (
	"errors"
	// "log"
	"fmt"
	"net/http"
	"time"

	"go-study/config"
	"go-study/router"
	"go-study/model"
	"go-study/router/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "config path")
)

func main() {
	pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	g := gin.New()
	gin.SetMode(viper.GetString("runmode"))

	// middleW := []gin.HandlerFunc{}

	router.Load(
		g,

		middleware.RequestId(),
		middleware.Logging(),
		// middleW...,
	)

	go func() {
		err := pingSer()

		if err != nil {
			fmt.Println("oh no")
		}

		fmt.Println("success")
	}()

	// init db
	model.DB.Init()
	defer model.DB.Close()

	g.Run(viper.GetString("addr"))
	// g.Run(":8077")
	// http.ListenAndServe(":8077", g)
}

func pingSer() error {
	for i := 0; i < 2; i++ {
		res, err := http.Get(viper.GetString("url") + "/sd/health")

		if err == nil && res.StatusCode == 200 {
			return nil
		}

		time.Sleep(time.Second)
	}

	return errors.New("bad route")
}
