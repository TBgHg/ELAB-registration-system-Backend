package wiki

import "github.com/gin-gonic/gin"

func ApplyGroup(engine *gin.RouterGroup) {
	group := engine.Group("/wiki")
	{
		group.GET("", searchWiki)
		group.POST("", newWiki)
		group.GET("/:wiki_id", getWikiByID)
		group.PATCH("/:wiki_id", updateWikiByID)
		group.DELETE("/:wiki_id", deleteWikiByID)
		group.GET("/:wiki_id/revision", searchWikiRevision)
	}
}
