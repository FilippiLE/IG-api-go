package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
)

var (
	Router *gin.Engine
)

func main() {
	startApp()
}

func startApp() {
	configureRuntime()
	initGinEngine()
	mapUrlsToControllers()
	Router.Run(":8080")
}

func initGinEngine() {

	Router = gin.Default()
	Router.RedirectFixedPath = false
	Router.RedirectTrailingSlash = false
}

func configureRuntime() {

	numCPU := runtime.NumCPU()
	fmt.Println(numCPU)
	runtime.GOMAXPROCS(numCPU)
}
