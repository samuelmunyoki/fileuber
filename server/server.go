package server

import (
	"os"
	"fmt"

	"github.com/gin-gonic/gin"
)


type LocalServer struct{

}

func NewLocalServer() *LocalServer {
	return &LocalServer{}
}


func (ls *LocalServer) StartLocalServer() {

	gin.SetMode(gin.DebugMode)

	addr := fmt.Sprintf(":%v", ShareCfg.Port )
	pid := os.Getpid()

	ShareCfg.SetShareConfigPID(int(pid))

	router := gin.Default()

	router.LoadHTMLGlob("./server/templates/*")

	router.Static("./static", "./server/static/")

	router.GET("/", GetHandler)

	router.POST("/", PostTHandler)
	
	router.GET("/home", HomeHandler)

	router.Run(addr)

	
}