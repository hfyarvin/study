package main

import (
	"./router"
	"io"
	"log"
	"os"
	// "net/http"
	// "time"

	"github.com/gin-gonic/gin"
)

func main() {
	//写log
	// gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	log.SetOutput(gin.DefaultWriter)

	r := gin.Default()
	router.Router(r) //路由
	// http.ListenAndServe(":555", r)
	// s := http.Server{
	// 	Addr:           ":8080",
	// 	Handler:        r,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// s.ListenAndServe()
	r.Run(":555")
}
