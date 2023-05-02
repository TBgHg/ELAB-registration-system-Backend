package middlewares

import (
	"elab-backend/internal/model"
	"github.com/gin-gonic/gin"
)

func NewSpaceIdRequiredError() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "space_id_required",
		Message: "/space/:space_id is required",
	}
}

func SpaceIdRequiredMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		spaceId := c.GetString("space_id")
		if spaceId == "" {
			c.AbortWithStatusJSON(400, NewSpaceIdRequiredError())
			return
		}
		c.Next()
	}
}
