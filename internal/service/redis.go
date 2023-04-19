package service

import (
	log "ELAB-registration-system-Backend/logger"
	"context"
	"github.com/google/uuid"
	"strconv"
	"time"
)

// Interview lock duration
const (
	interviewLockDuration = 10 * time.Second
	retryInterval         = 100 * time.Millisecond
	maxRetries            = 10
)

func (s *Service) Lock(interviewId int) (bool, func() bool) {

	ctx := context.Background()
	lockName := "interview_lock_" + strconv.Itoa(interviewId)
	identifier := uuid.New().String()

	result, err := s.Rdb.SetNX(ctx, lockName, identifier, interviewLockDuration).Result()
	if err != nil {
		log.Logger.Errorf(ctx, "Lock s.Rdb.SetNX err(%v)", err)
		return false, nil
	}
	if result == false {
		return false, nil
	}
	unlock := func() bool {
		ctx := context.Background()
		lockName := "interview_lock_" + strconv.Itoa(interviewId)

		if s.Rdb.Get(ctx, lockName).Val() == identifier {
			s.Rdb.Del(ctx, lockName)
			return true
		}
		return false
	}
	return true, unlock
}

func (s *Service) GetLock(interviewId int) (bool, func() bool) {

	ok, unlock := s.Lock(interviewId)
	if ok {
		return true, unlock
	}

	retries := 0
	for retries < maxRetries {
		time.Sleep(retryInterval)
		ok, unlock := s.Lock(interviewId)
		if ok {
			return true, unlock
		}
		retries++
	}

	return false, nil
}
