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
	selectQuery := `
contents.content_id as content_id, 
contents.current_history_id as current_history_id, 
contents.last_update_at as last_update_at,
histories.openid as openid,
users.name as name,
users.email as email,
histories.meta as meta`
	switch query.OrderBy {
	case content.TimeDesc:
		selectQuery += "FROM contents"
		break
	case content.LikeDesc:
		selectQuery += ", COUNT(likes.content_id) as _counts FROM contents"
		break
	}
	db = db.Select(
		selectQuery,
	).InnerJoins(
		"histories", `(histories.content_id = contents.content_id)
AND (histories.content_type = contents.content_type)
AND (histories.space_id = contents.space_id)
AND (histories.history_id = contents.current_history_id)`,
	).InnerJoins(
		"users", "users.openid = histories.openid",
	)
	if query.OrderBy == content.LikeDesc {
		db = db.InnerJoins(
			"likes",
			`
(likes.content_id = contents.content_id) AND
(likes.content_type = contents.content_type)
`,
		)
	}
	db = db.Where(&content.Content{
		SpaceId:     query.SpaceId,
		ContentType: contentType,
	})
	if query.Name != "" {
		db = db.Where(
			"meta->'$.title' LIKE ? AND space_id = ? AND content_type = ?",
			"%"+query.Name+"%",
			query.SpaceId, contentType,
		)
	}
	switch query.OrderBy {
	case content.TimeDesc:
		db = db.Order("last_update_at DESC")
		break
	case content.LikeDesc:
		db = db.Order("_counts DESC")
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
