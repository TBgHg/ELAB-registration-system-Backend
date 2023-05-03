package thread

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
	Content  string `json:"content" binding:"required"`
	SpaceId  string `form:"space_id" binding:"required,uuid"`
	ThreadId string `form:"thread_id" binding:"required,uuid"`
	Title    string `json:"title" binding:"required"`
}

func Update(ctx *gin.Context, request *UpdateRequest) error {
	postRequest := post.UpdateRequest{
		Content:     request.Content,
		SpaceId:     request.SpaceId,
		PostId:      request.ThreadId,
		Title:       request.Title,
		ContentType: content.Thread,
	}
	return post.Update(ctx, &postRequest)
}

type ContentResponse struct {
	Content string `json:"content"`
}

type GetContentByHistoryIdOptions struct {
	SpaceId   string
	ThreadId  string
	HistoryId string
}

func GetContentByHistoryId(ctx *gin.Context, options *GetContentByHistoryIdOptions) (string, error) {
	targetOptions := post.GetContentByHistoryIdOptions{
		SpaceId:     options.SpaceId,
		PostId:      options.ThreadId,
		ContentType: content.Thread,
		HistoryId:   options.HistoryId,
	}
	return post.GetContentByHistoryId(ctx, &targetOptions)
}

type SearchResponse struct {
	Threads []content.Head `json:"threads"`
}

func Search(ctx *gin.Context, query *content.Query) (*[]content.Head, error) {
	return post.Search(ctx, query, content.Thread)
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
		ContentType: content.Thread,
	}
	return post.Create(ctx, &targetRequest)
}

func HardDelete(ctx *gin.Context, spaceId string, threadId string) error {
	return post.HardDelete(ctx, &post.DeleteRequest{
		SpaceId: spaceId,
		PostId:  threadId,
	}, content.Thread)
}

func CheckExistsById(ctx *gin.Context, spaceId string, threadId string) (bool, error) {
	return content.CheckExistsById(ctx, &content.CheckExistsByIdOptions{
		SpaceId:     spaceId,
		ContentId:   threadId,
		ContentType: content.Thread,
	})
}

func CheckHistoryExistsById(ctx *gin.Context, spaceId string, threadId string, historyId string) (bool, error) {
	return content.CheckHistoryExistsById(ctx, &content.CheckHistoryExistsByIdOptions{
		SpaceId:     spaceId,
		ContentId:   threadId,
		ContentType: content.Thread,
		HistoryId:   historyId,
	})
}
