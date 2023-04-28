package http

import (
	log "elab-backend/logger"
	"fmt"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

// Callback, 重定向至前端页面
func callback(c *gin.Context) {
	// get code from query string
	code := c.Query("code")
	if code == "" {
		log.Logger.Error("callback code is empty")
		c.AbortWithStatusJSON(400, gin.H{"message": "code is empty"})
	}
	token, err := svc.OAuth2Config.Exchange(c, code)
	if err != nil {
		log.Logger.Error("callback OAuth2Config.Exchange failed err: " + err.Error())
		c.AbortWithStatusJSON(400, gin.H{"message": "OAuth2Config.Exchange failed"})
		return
	}
	url, err := url.Parse(svc.Conf.Mobile.Endpoint)
	if err != nil {
		log.Logger.Error("callback url.Parse failed err:" + err.Error())
		c.AbortWithStatusJSON(400, gin.H{"message": "url.Parse failed"})
		return
	}
	url.Path = "/callback"
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	query := url.Query()
	query.Set("access_token", token.AccessToken)
	query.Set("refresh_token", token.RefreshToken)
	query.Set("token_type", token.TokenType)
	query.Set("expires_in", (fmt.Sprintf("%d", (token.Expiry.Unix() - time.Now().In(timezone).Unix()))))
	query.Set("state", c.Query("state"))
	url.RawQuery = query.Encode()
	// redirect to frontend
	c.Redirect(301, url.String())
}
