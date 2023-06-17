package controller

import "github.com/gin-gonic/gin"

type TodoController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	FindOne(c *gin.Context)
	FindAll(c *gin.Context)
}
