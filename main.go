package main

import (
	"fmt"
	"log"
	"syscall"

	"github.com/fvbock/endless"

	"gin-blog/pkg/setting"
	"gin-blog/routers"
	"gin-blog/pkg/clean"
)

func main() {

	//router := gin.Default()
	//router.GET("/test", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "test",
	//	})
	//})

	//routersInit := routers.InitRouter()
	//
	//s := &http.Server{
	//	Addr:             fmt.Sprintf(":%d", setting.HTTPPort),
	//	Handler:           routersInit,
	//	ReadTimeout:       setting.ReadTimeout,
	//	WriteTimeout:      setting.WriteTimeout,
	//	MaxHeaderBytes:    1 <<20,
	//}
	//
	//s.ListenAndServe()

	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	// Add cron
	go func() {
		fmt.Println("Start clear cron task.")
		clean.CleanMySQLDeleted()
	}()

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}

