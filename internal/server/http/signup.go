package http

import (
	"ELAB-registration-system-Backend/internal/model"
	log "ELAB-registration-system-Backend/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signupSubmit(c *gin.Context) {
	req := new(model.SignupSubmitReq)
	if err := c.BindJSON(req); err != nil {
		log.Logger.Errorf(c, "signupSubmit c.BindJSON err(%v)", err)
		c.JSON(http.StatusOK, model.CommonResp{Code: 400, Msg: "参数错误"})
		return
	}

	value, exists := c.Get("openID")
	if !exists {
		log.Logger.Errorf(c, "signupSubmit c.Get openID err")
		return
	} else {
		req.OpenID = value.(string)
	}

	value, exists = c.Get("email")
	if !exists {
		log.Logger.Errorf(c, "signupSubmit c.Get email err")
		return
	} else {
		req.Email = value.(string)
	}

	c.JSON(http.StatusOK, svc.SignupSubmit(c, req))
}

func signupGet(c *gin.Context) {
	value, exists := c.Get("openID")
	if !exists {
		log.Logger.Errorf(c, "signupSubmit c.Get openID err")
		return
	}
	c.JSON(http.StatusOK, svc.SignupGet(c, value.(string)))
}

func signupUpdate(c *gin.Context) {
	req := new(model.SignupUpdateReq)
	if err := c.BindJSON(req); err != nil {
		log.Logger.Errorf(c, "signupUpdate c.BindJSON err(%v)", err)
		c.JSON(http.StatusOK, model.CommonResp{Code: 400, Msg: "参数错误"})
		return
	}
	value, exists := c.Get("openID")
	if !exists {
		log.Logger.Errorf(c, "signupSubmit c.Get openID err")
		return
	} else {
		req.OpenID = value.(string)
	}
	c.JSON(http.StatusOK, svc.SignupUpdate(c, req))
}
