// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package usercenter

import (
	"github.com/go-kratos/kratos/v2"

	"github.com/costa92/krm/internal/pkg/bootstrap"
	validation2 "github.com/costa92/krm/internal/pkg/validation"
	"github.com/costa92/krm/internal/usercenter/biz"
	"github.com/costa92/krm/internal/usercenter/server"
	"github.com/costa92/krm/internal/usercenter/service"
	"github.com/costa92/krm/internal/usercenter/validation"
	"github.com/costa92/krm/pkg/options"
)

// Injectors from wire.go:

func wireApp(appInfo bootstrap.AppInfo, config *server.Config, jwtOptions *options.JWTOptions, redisOptions *options.RedisOptions) (*kratos.App, func(), error) {
	logger := bootstrap.NewLogger(appInfo)
	appConfig := bootstrap.AppConfig{
		Info:   appInfo,
		Logger: logger,
	}
	bizBiz := biz.NewBiz()
	userCenterService := service.NewUserCenterService(bizBiz)
	validator, err := validation.New()
	if err != nil {
		return nil, nil, err
	}
	validationValidator := validation2.New(validator)
	v := server.NewMiddlewares(logger, validationValidator)
	httpServer := server.NewHTTPServer(config, userCenterService, v)
	grpcServer := server.NewGRPCServer(config, userCenterService, v)
	v2 := server.NewServers(httpServer, grpcServer)
	app := bootstrap.NewApp(appConfig, v2...)
	return app, func() {
	}, nil
}
