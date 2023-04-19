package http

import (
	"ELAB-registration-system-Backend/internal/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JwtMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 从请求头中获取Authorization字段
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			// 如果没有Authorization字段，返回401错误
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.CommonResp{Code: -1, Msg: "缺少token"})
			return
		}
		// 从Authorization字段中提取AccessToken
		var tokenStr string
		fmt.Sscanf(authHeader, "Bearer %s", &tokenStr)

		token, err := jwt.ParseWithClaims(tokenStr, &model.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("ELAB2023"), nil
		})
		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 { //token格式错误
					c.JSON(http.StatusOK, model.CommonResp{Code: -1, Msg: "token 格式错误"})
					c.Abort() //阻止执行
					return
				} else if ve.Errors&jwt.ValidationErrorExpired != 0 { //token过期
					c.JSON(http.StatusOK, model.CommonResp{Code: -1, Msg: "登录状态已失效，请重新登录"})
					c.Abort() //阻止执行
					return
				} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 { //token未激活
					c.JSON(http.StatusOK, model.CommonResp{Code: -1, Msg: "token 尚未激活"})
					c.Abort() //阻止执行
					return
				} else {
					c.JSON(http.StatusOK, model.CommonResp{Code: -1, Msg: "无法处理此Token"})
					c.Abort() //阻止执行
					return
				}
			}
		}

		if claims, ok := token.Claims.(*model.TokenClaims); ok && token.Valid {
			openID := claims.OpenID
			email := claims.Email
			c.Set("openID", openID)
			c.Set("email", email)
			c.Next()
			return
		}
		//失效的token
		c.JSON(http.StatusOK, model.CommonResp{Code: -1, Msg: "无效 token"})
		c.Abort() //阻止执行
	}
}
