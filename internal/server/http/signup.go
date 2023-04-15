package http

import "github.com/gin-gonic/gin"

func signupSubmit(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func signupGet(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func signupUpdate(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func signupDelete(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
