package http

import (
	"elab-backend/internal/model"
	log "elab-backend/logger"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
)

func userSubmit(c *gin.Context) {
	req := new(model.UserSubmitReq)
	if err := c.BindJSON(req); err != nil {
		log.Logger.Errorf(c, "userSubmit c.BindJSON err(%v)", err)
		c.JSON(http.StatusOK, model.CommonResp{Code: 400, Msg: "参数错误"})
		return
	}

	value, exists := c.Get("openID")
	if !exists {
		log.Logger.Errorf(c, "userSubmit c.Get openID err")
		return
	} else {
		req.OpenID = value.(string)
	}

	value, exists = c.Get("email")
	if !exists {
		log.Logger.Errorf(c, "userSubmit c.Get email err")
		return
	} else {
		req.Email = value.(string)
	}

	c.JSON(http.StatusOK, svc.UserSubmit(c, req))
}

func getUser(c *gin.Context) {
	value, exists := c.Get("openID")
	if !exists {
		log.Logger.Errorf(c, "getUser c.Get openID err")
		return
	}
	openID := c.Param("openid")
	log.Logger.Infof(c, "getUser openID() ctx.openID(%s)", openID, value.(string))
	c.JSON(http.StatusOK, svc.UserGet(c, openID))
}

func uploadAvatar(c *gin.Context) {
	value, exists := c.Get("openID")
	if !exists {
		log.Logger.Errorf(c, "uploadAvatar c.Get openID err")
		return
	}
	openID := c.Param("openid")
	log.Logger.Infof(c, "uploadAvatar openID() ctx.openID(%s)", openID, value.(string))
	avatar, err := c.FormFile("data")
	if err != nil {
		log.Logger.Errorf(c, "uploadAvatar c.FormFile err(%v)", err)
		c.JSON(http.StatusOK, model.CommonResp{Code: 400, Msg: "参数错误"})
		return
	}

	fileHandle, err := avatar.Open()
	if err != nil {
		log.Logger.Errorf(c, "uploadAvatar data.Open err(%v)", err)
		c.JSON(http.StatusOK, model.CommonResp{Code: 400, Msg: "参数错误"})
	}

	defer func(fileHandle multipart.File) {
		err = fileHandle.Close()
		if err != nil {
			log.Logger.Errorf(c, "uploadAvatar fileHandle.Close err(%v)", err)
		}
	}(fileHandle)

	avatarBytes, err := io.ReadAll(fileHandle)
	if err != nil {
		log.Logger.Errorf(c, "uploadAvatar io.ReadAll err(%v)", err)
		c.JSON(http.StatusOK, model.CommonResp{Code: 500, Msg: "服务器错误"})
	}

	c.JSON(http.StatusOK, svc.UploadAvatar(c, openID, avatarBytes))
}

func userUpdate(c *gin.Context) {
	req := new(model.UserUpdateReq)
	if err := c.BindJSON(req); err != nil {
		log.Logger.Errorf(c, "userUpdate c.BindJSON err(%v)", err)
		c.JSON(http.StatusOK, model.CommonResp{Code: 400, Msg: "参数错误"})
		return
	}
	value, exists := c.Get("openID")
	if !exists {
		log.Logger.Errorf(c, "userUpdate c.Get openID err")
		return
	} else {
		req.OpenID = value.(string)
	}
	c.JSON(http.StatusOK, svc.UserUpdate(c, req))
}
