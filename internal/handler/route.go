package handler

import "github.com/gin-gonic/gin"

// TODO: add handlers for this requests
func UserRoutes(req *gin.Engine) {
	req.POST("/users/signup")
	req.POST("/users/login")
	req.GET("/users/search")
	req.POST("/admin/addproduct")
	req.GET("/admin/viewproduct")
}
