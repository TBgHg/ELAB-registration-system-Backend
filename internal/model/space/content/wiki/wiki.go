package wiki

import (
	"elab-backend/internal/model/space/content"
	"elab-backend/internal/model/space/content/post"
	"github.com/gin-gonic/gin"
)

type ContentMeta struct{}

type HistoryMeta struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
}

type UpdateRequest struct {
	Content string `json:"content" binding:"required"`
	SpaceId string `form:"space_id" binding:"required,uuid"`
	WikiId  string `form:"wiki_id" binding:"required,uuid"`
	Title   string `json:"title" binding:"required"`
}

func Update(ctx *gin.Context, request *UpdateRequest) error {
	postRequest := post.UpdateRequest{
		Content:     request.Content,
		SpaceId:     request.SpaceId,
		PostId:      request.WikiId,
		Title:       request.Title,
		ContentType: content.Wiki,
	}
	return post.Update(ctx, &postRequest)
}

type ContentResponse struct {
	Content string `json:"content"`
}

type GetContentByHistoryIdOptions struct {
	SpaceId   string `form:"space_id" binding:"required,uuid"`
	WikiId    string `form:"wiki_id" binding:"required,uuid"`
	HistoryId string `form:"history_id" binding:"required,uuid"`
}

func GetContentByHistoryId(ctx *gin.Context, options *GetContentByHistoryIdOptions) (string, error) {
	targetOptions := post.GetContentByHistoryIdOptions{
		SpaceId:     options.SpaceId,
		PostId:      options.WikiId,
		ContentType: content.Wiki,
		HistoryId:   options.HistoryId,
	}
	return post.GetContentByHistoryId(ctx, &targetOptions)
}

type SearchResponse struct {
	Wikis []content.Head `json:"wikis"`
}

func Search(ctx *gin.Context, query *content.Query) (*[]content.Head, error) {
	return post.Search(ctx, query, content.Wiki)
}

type CreateRequest struct {
	SpaceId string `form:"space_id" binding:"required,uuid"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func Create(ctx *gin.Context, request *CreateRequest) error {
	targetRequest := post.CreateRequest{
		SpaceId:     request.SpaceId,
		Title:       request.Title,
		Content:     request.Content,
		ContentType: content.Wiki,
	}
	return post.Create(ctx, &targetRequest)
}

func HardDelete(ctx *gin.Context, spaceId string, wikiId string) error {
	return post.HardDelete(ctx, &post.DeleteRequest{
		SpaceId: spaceId,
		PostId:  wikiId,
	}, content.Wiki)
}

func CheckExistsById(ctx *gin.Context, spaceId string, threadId string) (bool, error) {
	return content.CheckExistsById(ctx, &content.CheckExistsByIdOptions{
		SpaceId:     spaceId,
		ContentId:   threadId,
		ContentType: content.Thread,
	})
}

//func CheckExistsByTitle(ctx *gin.Context, spaceId string, title string) (bool, error) {
//	contentHistory := content.History{
//		SpaceId:     spaceId,
//		ContentType: content.Wiki,
//	}
//	svc := service.GetService()
//	var count int64
//	err := svc.MySQL.WithContext(ctx).Model(&contentHistory).Where("meta->'&.title' = ?", title).Count(&count).Error
//	return count > 0, err
//}

func CheckHistoryExistsById(ctx *gin.Context, spaceId string, threadId string, historyId string) (bool, error) {
	return content.CheckHistoryExistsById(ctx, &content.CheckHistoryExistsByIdOptions{
		SpaceId:     spaceId,
		ContentId:   threadId,
		ContentType: content.Thread,
		HistoryId:   historyId,
	})
}
