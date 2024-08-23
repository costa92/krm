//go:build wireinject
// +build wireinject

package usercenter

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"

	"github.com/costa92/krm/internal/pkg/bootstrap"
	"github.com/costa92/krm/internal/pkg/validation"
	"github.com/costa92/krm/internal/usercenter/auth"
	"github.com/costa92/krm/internal/usercenter/biz"
	"github.com/costa92/krm/internal/usercenter/server"
	"github.com/costa92/krm/internal/usercenter/service"
	"github.com/costa92/krm/internal/usercenter/store"
	customvalidation "github.com/costa92/krm/internal/usercenter/validation"
	"github.com/costa92/krm/pkg/db"
	genericoptions "github.com/costa92/krm/pkg/options"
)

//go:generate go run github.com/google/wire/cmd/wire

func wireApp(
	bootstrap.AppInfo,
	*server.Config,
	*db.MySQLOptions,
	*genericoptions.JWTOptions,
	*genericoptions.RedisOptions,
	*genericoptions.KafkaOptions,
) (*kratos.App, func(), error) {
	wire.Build(
		bootstrap.ProviderSet,
		server.ProviderSet,
		store.ProviderSet,
		db.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		auth.ProviderSet,
		store.SetterProviderSet,
		NewAuthenticator,
		validation.ProviderSet,
		customvalidation.ProviderSet,
	)

	return nil, nil, nil
}
