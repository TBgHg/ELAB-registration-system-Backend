package service

import (
	"context"
	"elab-backend/internal/dao"
	"elab-backend/internal/model"
	log "elab-backend/logger"
	"errors"
	"gorm.io/gorm"
)

func (s *Service) InterviewSelect(c context.Context, req *model.InterviewSelectReq) (resp *model.InterviewSelectResp) {

	u := s.db.User
	user, err := s.db.User.WithContext(c).Where(u.OpenID.Eq(req.OpenID)).First()
	if err != nil {
		log.Logger.Errorf(c, "InterviewSelect req(%v) openID(%s) err(%v)", req, req.OpenID, err)
		resp = &model.InterviewSelectResp{
			CommonResp: &model.CommonResp{
				Code: 101,
				Msg:  "数据库查询错误",
			},
		}
		return
	}
	userID := user.ID

	// 取消面试场次单独进行处理
	if req.InterviewSessionID == 0 {
		// 1.将application表中的user_id、interview_session_id对应的行的state置为0
		// 2.将interview表中的interview_session_id对应的行的applied_num减一
		result, unlock := s.GetLock(int(req.OldInterviewSessionID))
		if result {
			defer unlock()
		} else {
			log.Logger.Errorf(c, "InterviewSelect s.GetLock failed req(%v) openID(%s) err(%v)", req, req.OpenID, errors.New("获取锁失败"))
			return &model.InterviewSelectResp{
				CommonResp: &model.CommonResp{
					Code: 102,
					Msg:  "获取锁失败",
				},
			}
		}
		err = s.db.Transaction(func(tx *dao.Query) (err error) {
			a := tx.Application
			_, err = tx.Application.WithContext(c).Where(a.UserID.Eq(userID)).Update(a.State, 0)
			if err != nil {
				return err
			}
			i := tx.InterviewSession
			_, err = tx.InterviewSession.WithContext(c).Where(i.ID.Eq(req.OldInterviewSessionID)).Update(i.AppliedNum, gorm.Expr("applied_num-?", 1))
			if err != nil {
				return err
			}
			return
		})
		if err != nil {
			log.Logger.Errorf(c, "InterviewSelect req(%v) openID(%s) err(%v)", req, req.OpenID, err)
			return &model.InterviewSelectResp{
				CommonResp: &model.CommonResp{
					Code: 101,
					Msg:  "数据库查询错误",
				},
			}
		}
		resp = &model.InterviewSelectResp{
			CommonResp: &model.CommonResp{
				Code: 0,
				Msg:  "success",
			},
		}
		return
	}

	// 选择或更新面试场次
	result, unlock := s.GetLock(int(req.OldInterviewSessionID))
	if result {
		defer unlock()
	}
	result2, unlock2 := s.GetLock(int(req.InterviewSessionID))
	if result2 {
		defer unlock2()
	}
	if !result || !result2 {
		log.Logger.Errorf(c, "InterviewSelect s.GetLock failed req(%v) openID(%s) err(%v)", req, req.OpenID, errors.New("获取锁失败"))
		return &model.InterviewSelectResp{
			CommonResp: &model.CommonResp{
				Code: 102,
				Msg:  "获取锁失败",
			},
		}
	}

	err = s.db.Transaction(func(tx *dao.Query) (err error) {
		a := tx.Application
		application, err2 := tx.Application.WithContext(c).Where(a.UserID.Eq(userID)).First()
		if err2 != nil {
			if errors.Is(err2, gorm.ErrRecordNotFound) {
				// application表中不存在该数据
				application = &model.Application{
					InterviewID: req.InterviewSessionID,
					State:       1,
					UserID:      userID,
				}
				if err2 = tx.Application.WithContext(c).Create(application); err2 != nil {
					return err2
				}
			} else {
				return err2
			}
		}
		application.InterviewID = req.InterviewSessionID
		application.State = 1

		if _, err = tx.Application.WithContext(c).Updates(application); err != nil {
			return err
		}

		// 这里需要判断能不能选择这个新场次
		i := tx.InterviewSession
		first, err := tx.InterviewSession.WithContext(c).Where(i.ID.Eq(req.InterviewSessionID)).First()
		if err != nil {
			return err
		}
		if first.AppliedNum >= first.Capacity {
			return errors.New("该场次已满")
		}

		_, err = tx.InterviewSession.WithContext(c).Where(i.ID.Eq(req.InterviewSessionID)).Update(i.AppliedNum, gorm.Expr("applied_num+?", 1))
		if err != nil {
			return err
		}

		// 将旧场次的人数-1
		_, err = tx.InterviewSession.WithContext(c).Where(i.ID.Eq(req.OldInterviewSessionID)).Update(i.AppliedNum, gorm.Expr("applied_num-?", 1))
		if err != nil {
			return err
		}

		return
	})

	if err != nil {
		if err.Error() == "该场次已满" {
			return &model.InterviewSelectResp{
				CommonResp: &model.CommonResp{
					Code: 103,
					Msg:  "该场次已满",
				},
			}
		}
		log.Logger.Errorf(c, "InterviewSelect req(%v) openID(%s) err(%v)", req, req.OpenID, err)
		return &model.InterviewSelectResp{
			CommonResp: &model.CommonResp{
				Code: 101,
				Msg:  "数据库查询错误",
			},
		}
	}
	resp = &model.InterviewSelectResp{
		CommonResp: &model.CommonResp{
			Code: 0,
			Msg:  "success",
		},
	}
	return
}

func (s *Service) InterviewGet(c context.Context, openID string) (resp *model.InterviewGetResp) {

	u := s.db.User
	user, err := s.db.User.WithContext(c).Where(u.OpenID.Eq(openID)).First()
	if err != nil {
		log.Logger.Errorf(c, "InterviewGet openID(%s) err(%v)", openID, err)
		resp = &model.InterviewGetResp{
			CommonResp: &model.CommonResp{
				Code: 101,
				Msg:  "数据库查询错误",
			},
		}
		return
	}
	userID := user.ID

	a := s.db.Application
	application, err := s.db.Application.WithContext(c).Where(a.UserID.Eq(userID)).First()
	if err != nil || application.State == 0 {
		if errors.Is(err, gorm.ErrRecordNotFound) || application.State == 0 {
			resp = &model.InterviewGetResp{
				CommonResp: &model.CommonResp{
					Code: 0,
					Msg:  "success",
				},
				InterviewID: 0,
			}
			return
		}
		log.Logger.Errorf(c, "InterviewGet openID(%s) err(%v)", openID, err)
		resp = &model.InterviewGetResp{
			CommonResp: &model.CommonResp{
				Code: 101,
				Msg:  "数据库查询错误",
			},
		}
		return
	}

	interviewID := application.InterviewID
	i := s.db.InterviewSession
	interview, err := s.db.InterviewSession.WithContext(c).Where(i.ID.Eq(interviewID)).First()
	if err != nil {
		log.Logger.Errorf(c, "InterviewGet openID(%s) err(%v)", openID, err)
		resp = &model.InterviewGetResp{
			CommonResp: &model.CommonResp{
				Code: 101,
				Msg:  "数据库查询错误",
			},
		}
		return
	}
	resp = &model.InterviewGetResp{
		CommonResp: &model.CommonResp{
			Code: 0,
			Msg:  "success",
		},
		InterviewID: interview.ID,
		StartTime:   interview.StartTime,
		EndTime:     interview.EndTime,
		Location:    interview.Location,
		Capacity:    interview.Capacity,
		AppliedNum:  interview.AppliedNum,
	}
	return
}
