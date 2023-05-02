package comment

import (
	"elab-backend/internal/model/space/content"
	"elab-backend/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BuildDatabaseQuery(ctx *gin.Context, query *GetRequest) (*gorm.DB, error) {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		return nil, err
	}
	selectQuery := `
contents.content_id as content_id, 
contents.last_update_at as last_update_at,
histories.content as content,
histories.openid as openid,
users.name as name,
users.email as email,
COUNT(likes.content_id) as like_counts
COUNT(likes.openid = ?) = 1 as is_liked
`
	db := svc.MySQL.WithContext(ctx).Select(
		selectQuery, token.Subject,
	).InnerJoins(
		"histories", `(histories.content_id = contents.content_id)
AND (histories.content_type = contents.content_type)
AND (histories.space_id = contents.space_id)
AND (histories.history_id = contents.current_history_id)`,
	).InnerJoins(
		"users", "users.openid = histories.openid",
	).InnerJoins(
		"likes", `
(likes.content_id = contents.content_id) AND
(likes.content_type = contents.content_type)
`,
	).Where(&content.Content{
		SpaceId:     query.SpaceId,
		ContentType: content.Comment,
	}).Where("meta->'$.thread_id' = ?", query.ThreadId).Group("contents.content_id")
	return db, nil
}
