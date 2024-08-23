//go:build wireinject
// +build wireinject

package usercenter

import (
	"github.com/costa92/krm/internal/pkg/bootstrap"
	"github.com/costa92/krm/internal/pkg/validation"
	"github.com/costa92/krm/internal/usercenter/biz"
	"github.com/costa92/krm/internal/usercenter/server"
	"github.com/costa92/krm/internal/usercenter/service"
	customvalidation "github.com/costa92/krm/internal/usercenter/validation"
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

//go:generate go run github.com/google/wire/cmd/wire

func wireApp(
	bootstrap.AppInfo,
	*server.Config,
) (*kratos.App, func(), error) {
	wire.Build(
		bootstrap.ProviderSet,
		server.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		validation.ProviderSet,
		customvalidation.ProviderSet,
	)

	return nil, nil, nil
}
