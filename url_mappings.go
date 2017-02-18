package main

import (
	"github.com/FilippiLE/IG-api-go/controllers/insta_controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func mapUrlsToControllers() {

	Router.GET("/ping", Ping)
	Router.GET("/permiso", insta_controller.permiso)
	Router.GET("/acces_token", insta_controller.accestoken)
}

func Ping(c *gin.Context) {

	c.String(http.StatusOK, "pong")
}
