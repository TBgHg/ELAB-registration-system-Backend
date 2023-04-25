package service

import (
	"ELAB-registration-system-Backend/internal/model"
	log "ELAB-registration-system-Backend/logger"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Service) UserSubmit(c context.Context, req *model.UserSubmitReq) (resp *model.UserSubmitResp) {

	user := &model.User{
		OpenID:       req.OpenID,
		Name:         req.Name,
		StudentID:    req.StudentID,
		Gender:       req.Gender,
		Class:        req.Class,
		Position:     req.Position,
		Mobile:       req.Mobile,
		Mail:         req.Email,
		Group:        req.Group,
		Introduction: req.Introduction,
		Awards:       req.Awards,
		Reason:       req.Reason,
	}
	err := s.db.User.WithContext(c).Create(user)
	if err != nil {
		log.Logger.Errorf(c, "UserSubmit req(%v) openID(%s) err(%v)", req, req.OpenID, err)
		resp = &model.UserSubmitResp{
			CommonResp: &model.CommonResp{
				Code: 101,
				Msg:  "数据库查询错误",
			},
		}
		return
	}
	resp = &model.UserSubmitResp{
		CommonResp: &model.CommonResp{
			Code: 0,
			Msg:  "success",
		},
	}
	return
}

func (s *Service) UserGet(c *gin.Context, openID string) (resp *model.UserGetResp) {

	u := s.db.User
	user, err := s.db.User.WithContext(c).Where(u.OpenID.Eq(openID)).First()
	if err != nil {
		log.Logger.Errorf(c, "UserGet openID(%s) err(%v)", openID, err)
		resp = &model.UserGetResp{
			CommonResp: &model.CommonResp{
				Code: 500,
				Msg:  "服务器错误",
			},
		}
		return
	}
	resp = &model.UserGetResp{
		CommonResp: &model.CommonResp{
			Code: 0,
			Msg:  "success",
		},
		Name:         user.Name,
		StudentID:    user.StudentID,
		Gender:       user.Gender,
		Avatar:       user.Avatar,
		IsELABer:     user.IsELABer,
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

func (s *Service) UploadAvatar(c *gin.Context, openID string, avatarBytes []byte) (resp *model.UploadAvatarResp) {
	// 上传头像到OSS
	fileID := uuid.New().String()
	if !s.UploadFile(avatarBytes, fileID, "picture") {
		log.Logger.Errorf(c, "UploadAvatar openID(%s) fileID(%s) err", openID, fileID)
	}
	// 更新数据库
	u := s.db.User
	_, err := s.db.User.WithContext(c).Where(u.OpenID.Eq(openID)).Update(u.Avatar, fileID)
	if err != nil {
		log.Logger.Errorf(c, "UploadAvatar openID(%s) fileID(%s) err(%v)", openID, fileID, err)
		resp = &model.UploadAvatarResp{
			CommonResp: &model.CommonResp{
				Code: 500,
				Msg:  "服务器错误",
			},
		}
		return
	}
	videoURL := s.Conf.OssConfig.OssPreURL + fileID + ".jpg"
	resp = &model.UploadAvatarResp{
		CommonResp: &model.CommonResp{
			Code: 0,
			Msg:  "success",
		},
		Avatar: videoURL,
	}
	return
}

func (s *Service) UserUpdate(c *gin.Context, req *model.UserUpdateReq) (resp *model.UserUpdateResp) {
	resp = new(model.UserUpdateResp)
	user := &model.User{
		OpenID:       req.OpenID,
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
	_, err := s.db.User.WithContext(c).Where(u.OpenID.Eq(req.OpenID)).Updates(user)
	if err != nil {
		log.Logger.Errorf(c, "UserSubmit req(%v) openID(%s) err(%v)", req, req.OpenID, err)
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
