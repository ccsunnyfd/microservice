package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

type Auth interface {
	Authenticate(token string) error
}

// Authenticator is a server middleware that auth from user edge service.
func Authenticator(auth Auth) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			header, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, errors.BadRequest("AUTHENTICATOR", "get request context error")
			}
			token := header.RequestHeader().Get("token")
			if err := auth.Authenticate(token); err != nil {
				return nil, errors.BadRequest("AUTHENTICATOR", err.Error())
			}
			return handler(ctx, req)
		}
	}
}


