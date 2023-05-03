package model

import (
	"elab-backend/internal/model/application"
	"elab-backend/internal/model/space"
	"elab-backend/internal/model/space/content"
	"elab-backend/internal/model/space/member"
	"elab-backend/internal/model/user"
	"golang.org/x/sync/errgroup"
)

func Init() error {
	var eg errgroup.Group
	eg.Go(application.Init)
	eg.Go(content.Init)
	eg.Go(space.Init)
	eg.Go(member.Init)
	eg.Go(user.Init)
	return eg.Wait()
}
