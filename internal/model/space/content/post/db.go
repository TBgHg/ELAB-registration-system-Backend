package post

import (
	"crypto/md5"
	"elab-backend/internal/model/space/content"
	"elab-backend/internal/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

func BuildDatabaseQuery(ctx *gin.Context, query *content.Query, contentType content.Type) *gorm.DB {
	svc := service.GetService()
	db := svc.MySQL.WithContext(ctx)
	db = db.Raw(`
SELECT
	c.content_id as content_id, 
	c.current_history_id as current_history_id, 
	c.last_update_at as last_update_at,
	h.openid as openid,
	u.name as name,
	u.email as email,
	h.meta as meta,
	COUNT(l.content_id) as like_counts
FROM contents c, histories h, users u, likes l
INNER JOIN histories h ON
    (h.content_id = c.content_id) AND
    (h.content_type = c.content_type) AND
    (h.space_id = c.space_id) AND
    (h.history_id = c.current_history_id)
INNER JOIN users u ON u.openid = h.openid
INNER JOIN likes l ON l.content_id = c.content_id AND l.content_type = c.content_type
WHERE (c.deleted_at IS NULL AND h.deleted_at IS NULL AND u.deleted_at IS NULL AND l.deleted_at IS NULL)
AND h.meta->'$.title' LIKE '%?%' AND h.space_id = ? AND h.content_type = ?
GROUP BY c.content_id, c.current_history_id, c.last_update_at, h.openid, u.name, u.email, h.meta
`, query.Name, query.SpaceId, contentType)
	switch query.OrderBy {
	case content.TimeDesc:
		db = db.Order("last_update_at DESC")
		break
	case content.LikeDesc:
		db = db.Order("like_counts DESC")
		break
	}
	return db
}

func ParseDatabaseSearchResult(rawResult *[]content.SearchDatabaseResult) (*[]content.Head, error) {
	resultChan := make(chan *content.Head, len(*rawResult))
	var eg errgroup.Group
	for _, c := range *rawResult {
		eg.Go(func() error {
			var meta HistoryMeta
			err := json.Unmarshal([]byte(c.Meta), &meta)
			if err != nil {
				return err
			}
			md5Hash := md5.New()
			md5Hash.Write([]byte(c.Email))
			md5HashValue := md5Hash.Sum(nil)
			avatar := "https://gravatar.loli.net/avatar/" + string(md5HashValue)
			// 根据history.openid获取用户信息（email和name）
			head := content.Head{
				ContentId:        c.ContentId,
				CurrentHistoryId: c.CurrentHistoryId,
				LastUpdateAt:     c.LastUpdateAt,
				Title:            meta.Title,
				Summary:          meta.Summary,
				Author: content.Author{
					OpenId: c.OpenId,
					Avatar: avatar,
					Name:   c.Name,
				},
			}
			resultChan <- &head
			return nil
		})
	}
	err := eg.Wait()
	if err != nil {
		return nil, err
	}
	var resultGroup []content.Head
	for result, ok := <-resultChan; ok; {
		resultGroup = append(resultGroup, *result)
	}
	return &resultGroup, nil
}
