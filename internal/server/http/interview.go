package http

import "github.com/gin-gonic/gin"

func interviewSelect(c *gin.Context) {
	svc.InterviewSelect()
	c.JSON(200, gin.H{"message": "pong"})
}

func interviewGet(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func interviewUpdate(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
