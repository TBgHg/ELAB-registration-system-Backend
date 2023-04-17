package service

import (
	log "ELAB-registration-system-Backend/common/logger"
	"ELAB-registration-system-Backend/internal/model"
	"context"
	"github.com/gin-gonic/gin"
)

func (s *Service) SignupSubmit(c context.Context, req *model.SignupSubmitReq) (resp *model.SignupSubmitResp, err error) {
	tokenClaims := c.Value("user").(*model.TokenClaims)
	user := &model.User{
		OpenID:       tokenClaims.OpenID,
		Name:         req.Name,
		StudentID:    req.StudentID,
		Gender:       req.Gender,
		Class:        req.Class,
		Position:     req.Position,
		Mobile:       req.Mobile,
		Mail:         tokenClaims.Email,
		Group:        req.Group,
		Introduction: req.Introduction,
		Awards:       req.Awards,
		Reason:       req.Reason,
	}
	err = s.db.User.WithContext(c).Create(user)
	if err != nil {
		log.Logger.Errorf(c, "SignupSubmit req(%v) openID(%s) err(%v)", req, tokenClaims.OpenID, err)
		resp.Code = 500
		resp.Msg = "数据创建失败"
		return
	}
	return
}

func (s *Service) SignupGet(c *gin.Context, req *model.SignupGetReq) (resp *model.SignupGetResp, err error) {
	tokenClaims := c.Value("user").(*model.TokenClaims)
	openID := tokenClaims.OpenID
	u := s.db.User
	user, err := s.db.User.WithContext(c).Where(u.OpenID.Eq(openID)).First()
	if err != nil {
		log.Logger.Errorf(c, "SignupGet req(%v) openID(%s) err(%v)", req, openID, err)
		return nil, err
	}
	resp = &model.SignupGetResp{
		CommonResp: &model.CommonResp{
			Code: 0,
			Msg:  "success",
		},
		Name:         user.Name,
		StudentID:    user.StudentID,
		Gender:       user.Gender,
		Class:        user.Class,
		Position:     user.Position,
		Mobile:       user.Mobile,
		Group:        user.Group,
		Introduction: user.Introduction,
		Awards:       user.Awards,
		Reason:       user.Reason,
	}
	return
}

func (s *Service) SignupUpdate(c *gin.Context, req *model.SignupUpdateReq) (resp *model.SignupUpdateResp, err error) {
	return
}

func (s *Service) SignupDelete(c *gin.Context, req *model.SignupDeleteReq) (resp *model.SignupDeleteResp, err error) {
	return
}
