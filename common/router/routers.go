package router

import (
	"context"
	"go-web/common/hooks"
	"go-web/internal/loginserver"
	"go-web/internal/testserver"
	"go-web/internal/userserver"
	"go-web/rpc/login"
	"go-web/rpc/test"
	"go-web/rpc/user"
	"net/http"

	"github.com/twitchtv/twirp"
)

func Routers() *http.ServeMux {
	mux := http.NewServeMux()
	hooks := hooks.NewServerHooks()

	lserver := &loginserver.LoginServer{}
	loginHandler := login.
		NewLoginServer(
			lserver,
			twirp.WithServerPathPrefix(""))
	mux.Handle(
		loginHandler.PathPrefix(),
		loginHandler,
	)

	hbserver := &testserver.Server{}
	testHandler := test.
		NewTestServer(
			hbserver,
			twirp.WithServerPathPrefix(""))
	mux.Handle(
		testHandler.PathPrefix(),
		testHandler,
	)

	userserver := &userserver.UserServer{}
	userHandler := user.NewUserServer(
		userserver,
		twirp.WithServerHooks(hooks),
		twirp.WithServerPathPrefix(""),
	)
	mux.Handle(
		userHandler.PathPrefix(),
		WithAuth(userHandler),
	)

	return mux
}

func WithAuth(base http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		auth := r.Header.Get("Authorization")
		ctx = context.WithValue(ctx, hooks.AuthHeader("auth"), auth)
		r = r.WithContext(ctx)
		base.ServeHTTP(w, r)
	})
}
