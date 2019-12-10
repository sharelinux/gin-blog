package main

import (
	"fmt"
	"net/http"

	"gin-blog/pkg/setting"
	"gin-blog/routers"
)

func main() {

	//router := gin.Default()
	//router.GET("/test", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "test",
	//	})
	//})
	routersInit := routers.InitRouter()

	s := &http.Server{
		Addr:             fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:           routersInit,
		ReadTimeout:       setting.ReadTimeout,
		WriteTimeout:      setting.WriteTimeout,
		MaxHeaderBytes:    1 <<20,
	}

	s.ListenAndServe()
}

