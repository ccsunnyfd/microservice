package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/goburrow/cache"
	"github.com/gorilla/handlers"
	courseEdgeV1 "microservice/api/edge/course/interface/v1"
	"microservice/app/edge/course/interface/internal/biz"
	"microservice/app/edge/course/interface/internal/conf"
	"microservice/app/edge/course/interface/internal/service"
	"microservice/app/middleware"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger, auth middleware.Auth, s *service.CourseInterface) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			middleware.Authenticator(auth),
			validate.Validator(),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}
	//if c.Http.Network != "" {
	//	opts = append(opts, http.Network(c.Http.Network))
	//}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	courseEdgeV1.RegisterCourseEdgeInterfaceHTTPServer(srv, s)
	return srv
}

type Authenticator func(token string) error

func (a Authenticator) Authenticate(token string) error {
	return a(token)
}

func NewAuthMiddleware(c *conf.Data_Cache, repo biz.AuthRepo) (auth middleware.Auth, cleanup func()) {
	load := func(k cache.Key) (cache.Value, error) {
		token, _ := k.(string)
		return repo.VerifyToken(context.Background(), token)
	}
	local := cache.NewLoadingCache(load,
		cache.WithMaximumSize(int(c.MaximumSize)),                     // Limit number of entries in the cache.
		cache.WithExpireAfterAccess(c.ExpireAfterAccess.AsDuration()), // Expire entries after 1 minute since last accessed.
		cache.WithRefreshAfterWrite(c.RefreshAfterWrite.AsDuration()), // Expire entries after 2 minutes since last created.
	)
	return Authenticator(func(token string) error {
			_, err := local.Get(token)
			return err
		}), func() {
			if err := local.Close(); err != nil {
				_ = log.DefaultLogger.Log(log.LevelError, "close local cache error")
			}
		}
}
