package server

import (
	"context"
	"encoding/json"
	"github.com/costa92/krm/pkg/log"
	krtlog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"github.com/superproj/onex/pkg/authn"
	"github.com/superproj/onex/pkg/i18n"
	"golang.org/x/text/language"
)

// ProviderSet defines a wire provider set.
var ProviderSet = wire.NewSet(NewServers, NewGRPCServer, NewHTTPServer, NewMiddlewares)

// NewServers is a wire provider function that creates and returns a slice of transport servers.
func NewServers(hs *http.Server, gs *grpc.Server) []transport.Server {
	return []transport.Server{hs, gs}
}

// NewMiddlewares return middlewares used by grpc and http server both.
func NewMiddlewares(logger krtlog.Logger, a authn.Authenticator, v validate.IValidator) []middleware.Middleware {
	return []middleware.Middleware{
		recovery.Recovery(
			recovery.WithHandler(func(ctx context.Context, rq, err any) error {
				data, _ := json.Marshal(rq)
				log.C(ctx).Errorw(err.(error), "Catching a panic", "rq", string(data))
				return nil
			}),
		),

		i18nmw.Translator(i18n.WithLanguage(language.English), i18n.WithFS(locales.Locales)),
		ratelimit.Server(),
		tracing.Server(),
		selector.Server(jwt.Server(a)).Match(NewWhiteListMatcher()).Build(),
		validate.Validator(v),
		logging.Server(logger),
	}
}
