package http

import (
	"elab-backend/internal/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"log"
	"net/http"
)

func interviewSelect(c *gin.Context) {
	req := new(model.InterviewSelectReq)
	if err := c.BindJSON(req); err != nil {
		slog.Error("interviewSelect c.BindJSON err(%w)", err)
		c.JSON(http.StatusOK, model.CommonResp{Code: 400, Msg: "参数错误,BindJSON失败"})
		return
	}
	if !req.Validate() {
		log.Logger.Infof(c, "interviewSelect req.Validate failed req(%+v)", req)
		c.JSON(http.StatusOK, model.CommonResp{Code: 400, Msg: "参数错误,Validate失败"})
		return
	}

	value, exists := c.Get("openID")
	if !exists {
		log.Logger.Errorf(c, "userSubmit c.Get openID err")
		return
	} else {
		req.OpenID = value.(string)
	}

	c.JSON(http.StatusOK, svc.InterviewSelect(c, req))
}

func interviewGet(c *gin.Context) {
	value, exists := c.Get("openID")
	if !exists {
		log.Logger.Errorf(c, "userSubmit c.Get openID err")
		return
	}
	c.JSON(http.StatusOK, svc.InterviewGet(c, value.(string)))
}
