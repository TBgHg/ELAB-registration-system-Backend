package http

import (
	"ELAB-registration-system-Backend/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signupSubmit(c *gin.Context) {
	var req *model.SignupSubmitReq
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusOK, model.CommonResp{Code: 400, Msg: "参数错误"})
		return
	}
	resp, err := svc.SignupSubmit(c, req)
	if err != nil {
		c.JSON(http.StatusOK, model.CommonResp{Code: 500, Msg: "服务器错误"})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func signupGet(c *gin.Context) {
	req, err := new(model.SignupGetReq), error(nil)
	if err = c.Bind(req); err != nil {
		c.JSON(http.StatusOK, model.CommonResp{Code: 400, Msg: "参数错误"})
		return
	}
	resp, err := svc.SignupGet(c, req)
	if err != nil {
		c.JSON(http.StatusOK, model.CommonResp{Code: 500, Msg: "服务器错误"})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func signupUpdate(c *gin.Context) {
	req, err := new(model.SignupUpdateReq), error(nil)
	if err = c.Bind(req); err != nil {
		c.JSON(http.StatusOK, model.CommonResp{Code: 400, Msg: "参数错误"})
		return
	}
	resp, err := svc.SignupUpdate(c, req)
	if err != nil {
		c.JSON(http.StatusOK, model.CommonResp{Code: 500, Msg: "服务器错误"})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func signupDelete(c *gin.Context) {
	req, err := new(model.SignupDeleteReq), error(nil)
	if err = c.Bind(req); err != nil {
		c.JSON(http.StatusOK, model.CommonResp{Code: 400, Msg: "参数错误"})
		return
	}
	resp, err := svc.SignupDelete(c, req)
	if err != nil {
		c.JSON(http.StatusOK, model.CommonResp{Code: 500, Msg: "服务器错误"})
		return
	}
	c.JSON(http.StatusOK, resp)
}
