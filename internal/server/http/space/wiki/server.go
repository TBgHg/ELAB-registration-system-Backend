package wiki

import (
	"elab-backend/internal/server/middlewares"
	"github.com/gin-gonic/gin"
)

func ApplyGroup(engine *gin.RouterGroup) {
	group := engine.Group("/wiki")
	{
		group.GET("", middlewares.SpacePublicPermissionMiddleware(), searchWiki)
		group.POST("", middlewares.SpacePrivatePermissionMiddleware(), createWiki)
		wikiIdSpecifiedGroup := group.Group("/:wiki_id")
		wikiIdSpecifiedGroup.Use(middlewares.WikiIdRequiredMiddleware())
		wikiIdSpecifiedGroup.PATCH("", middlewares.SpacePrivatePermissionMiddleware(), updateWikiById)
		wikiIdSpecifiedGroup.DELETE("", middlewares.SpacePrivatePermissionMiddleware(), deleteWikiById)
		wikiIdSpecifiedGroup.GET("/history/:history_id", middlewares.SpacePublicPermissionMiddleware(), getWikiHistoryById)
	}
}
