package service

import (
	"ELAB-registration-system-Backend/internal/model"
	log "ELAB-registration-system-Backend/logger"
	"context"
	"github.com/gin-gonic/gin"
)

func (s *Service) SignupSubmit(c context.Context, req *model.SignupSubmitReq) (resp *model.SignupSubmitResp) {
	//tokenClaims := c.Value("user").(*model.TokenClaims)
	// todo:待删除
	tokenClaims := new(model.TokenClaims)
	tokenClaims.Email = "312"
	tokenClaims.OpenID = "123"

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
	err := s.db.User.WithContext(c).Create(user)
	if err != nil {
		log.Logger.Errorf(c, "SignupSubmit req(%v) openID(%s) err(%v)", req, tokenClaims.OpenID, err)
		resp = &model.SignupSubmitResp{
			CommonResp: &model.CommonResp{
				Code: 101,
				Msg:  "数据库查询错误",
			},
		}
		return
	}
	resp = &model.SignupSubmitResp{
		CommonResp: &model.CommonResp{
			Code: 0,
			Msg:  "success",
		},
	}
	return
}

func (s *Service) SignupGet(c *gin.Context) (resp *model.SignupGetResp) {

	//tokenClaims := c.Value("user").(*model.TokenClaims)
	// todo:待删除
	tokenClaims := new(model.TokenClaims)
	tokenClaims.OpenID = "123"

	openID := tokenClaims.OpenID
	u := s.db.User
	user, err := s.db.User.WithContext(c).Where(u.OpenID.Eq(openID)).First()
	if err != nil {
		log.Logger.Errorf(c, "SignupGet openID(%s) err(%v)", openID, err)
		resp = &model.SignupGetResp{
			CommonResp: &model.CommonResp{
				Code: 500,
				Msg:  "服务器错误",
			},
		}
		return
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

func (s *Service) SignupUpdate(c *gin.Context, req *model.SignupUpdateReq) (resp *model.SignupUpdateResp) {
	resp = new(model.SignupUpdateResp)
	//tokenClaims := c.Value("user").(*model.TokenClaims)
	// todo:待删除
	tokenClaims := new(model.TokenClaims)
	tokenClaims.OpenID = "123"
	openID := tokenClaims.OpenID
	user := &model.User{
		OpenID:       tokenClaims.OpenID,
		Name:         req.Name,
		StudentID:    req.StudentID,
		Gender:       req.Gender,
		Class:        req.Class,
		Position:     req.Position,
		Mobile:       req.Mobile,
		Group:        req.Group,
		Introduction: req.Introduction,
		Awards:       req.Awards,
		Reason:       req.Reason,
	}
	u := s.db.User
	_, err := s.db.User.WithContext(c).Where(u.OpenID.Eq(openID)).Updates(user)
	if err != nil {
		log.Logger.Errorf(c, "SignupSubmit req(%v) openID(%s) err(%v)", req, tokenClaims.OpenID, err)
		resp.CommonResp = &model.CommonResp{
			Code: 101,
			Msg:  "数据库查询错误",
		}
		return
	}
	resp.CommonResp = &model.CommonResp{
		Code: 0,
		Msg:  "success",
	}
	return
}
