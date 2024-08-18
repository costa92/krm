package usercenter

import (
	"github.com/costa92/krm/internal/pkg/bootstrap"
	"github.com/costa92/krm/internal/usercenter/server"
	"github.com/costa92/krm/internal/usercenter/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

func wireApp(bootstrap.AppInfo, *server.Config) (*kratos.App, func(), error) {
	wire.Build(
		bootstrap.ProviderSet,
		server.ProviderSet,
		service.ProviderSet,
	)
	return nil, nil, nil
}
