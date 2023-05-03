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
	db := svc.MySQL.WithContext(ctx).Raw(`
SELECT
	c.content_id as content_id, 
	c.last_update_at as last_update_at,
	h.content as content,
	h.openid as openid,
	u.name as name,
	u.email as email,
	COUNT(l.content_id) as like_counts,
	COUNT(l.openid = @openid) = 1 as is_liked
FROM contents c, histories h, users u, likes l
INNER JOIN histories h
	ON (h.content_id = c.content_id) AND
	(h.content_type = c.content_type) AND
	(h.space_id = c.space_id) AND
	(h.history_id = c.current_history_id)
INNER JOIN users u ON u.openid = h.openid
INNER JOIN likes l ON l.content_id = c.content_id AND l.content_type = c.content_type
WHERE
    c.space_id = ? AND c.content_type = ? AND
		c.deleted_at IS NULL AND h.deleted_at IS NULL AND u.deleted_at IS NULL AND l.deleted_at IS NULL AND
		h.meta->'$.thread_id' = ?
GROUP BY c.content_id, c.last_update_at, h.content, h.openid, u.name, u.email

`, token.Subject, query.SpaceId, content.Thread, query.ThreadId)
	return db, nil
}
