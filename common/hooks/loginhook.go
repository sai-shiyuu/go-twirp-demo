package hooks

import (
	"context"
	"fmt"
	"strings"

	"go-web/common/log"

	redis "go-web/common/cache"

	"github.com/sirupsen/logrus"
	"github.com/twitchtv/twirp"
)

var logger *logrus.Logger

func init() {
	logger = log.GetLogger()
}

type Auth struct {
	Id    int64
	Token string
}

func NewServerHooks() *twirp.ServerHooks {
	return &twirp.ServerHooks{
		RequestReceived: func(ctx context.Context) (context.Context, error) {
			logger.Infof("Request received.%v", ctx)

			authHeader := ctx.Value("auth").(string)
			if authHeader == "" {
				return ctx, twirp.NewError(twirp.Unauthenticated, "no authorization header")
			}
			parts := strings.Split(authHeader, " ")
			if len(parts) < 2 {
				return ctx, twirp.NewError(twirp.Unauthenticated, "invalid authorization header format")
			}
			key := fmt.Sprintf("login:token:%s", parts[0])
			token := redis.GetRedisClient().Get(ctx, key)
			if parts[1] != token.Val() {
				return ctx, twirp.NewError(twirp.Unauthenticated, "invalid token")
			}
			return ctx, nil
		},

		Error: func(ctx context.Context, err twirp.Error) context.Context {
			logger.Errorf("Error: %v", err)
			return ctx
		},
	}
}
