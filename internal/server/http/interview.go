package http

import (
	"ELAB-registration-system-Backend/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func interviewSelect(c *gin.Context) {
	req, err := new(model.InterviewSelectReq), error(nil)
	if err = c.Bind(req); err != nil {
		c.JSON(http.StatusOK, model.CommonResp{Code: 400, Msg: "参数错误"})
		return
	}
	svc.InterviewSelect()
	c.JSON(200, gin.H{"message": "pong"})
}

func interviewGet(c *gin.Context) {
	req, err := new(model.InterviewGetReq), error(nil)
	if err = c.Bind(req); err != nil {
		c.JSON(http.StatusOK, model.CommonResp{Code: 400, Msg: "参数错误"})
		return
	}
	c.JSON(200, gin.H{"message": "pong"})
}

func interviewUpdate(c *gin.Context) {
	req, err := new(model.InterviewUpdateReq), error(nil)
	if err = c.Bind(req); err != nil {
		c.JSON(http.StatusOK, model.CommonResp{Code: 400, Msg: "参数错误"})
		return
	}
	c.JSON(200, gin.H{"message": "pong"})
}
