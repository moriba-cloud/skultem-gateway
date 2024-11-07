package app

import (
	"context"
	"fmt"
	"github.com/moriba-build/ose/ddd"
	"github.com/moriba-cloud/skultem-gateway/domain/auth"
	"github.com/moriba-cloud/skultem-gateway/domain/core"
	"github.com/moriba-cloud/skultem-gateway/domain/user"
	"go.uber.org/zap"
)

type (
	aAuth struct {
		repo   user.Repo
		logger *zap.Logger
	}
	argsAuth struct {
		Repo   user.Repo
		Logger *zap.Logger
	}
)

func (a aAuth) Login(ctx context.Context, email string, password string) (*ddd.Response[auth.Domain], error) {
	res, err := a.repo.FindByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("no account found for this email: %s", email)
	}
	record := res.Record()

	if match := core.CheckPassword(record.Password().Hash, password); !match {
		return nil, fmt.Errorf("try again this password is incorrect")
	}

	token, err := auth.New(record.ID())

	return ddd.NewResponse(ddd.ResponseArgs[auth.Domain]{
		Record: token,
	}), nil
}

func (a aAuth) Access(ctx context.Context) (*ddd.Response[auth.Domain], error) {
	refresh := auth.ActiveRefreshToken(ctx)
	record, err := auth.Existing(refresh)
	if err != nil {
		return nil, err
	}

	return ddd.NewResponse(ddd.ResponseArgs[auth.Domain]{
		Record: record,
	}), nil
}

func (a aAuth) Me(ctx context.Context) (*ddd.Response[auth.User], error) {
	payload := auth.ActiveUser(ctx, "user")
	return ddd.NewResponse(ddd.ResponseArgs[auth.User]{
		Record: payload,
	}), nil
}

func NewAuth(args argsAuth) auth.App {
	return &aAuth{
		repo:   args.Repo,
		logger: args.Logger,
	}
}
