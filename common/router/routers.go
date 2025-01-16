package router

import (
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
		twirp.WithServerPathPrefix(""),
	)
	mux.Handle(
		userHandler.PathPrefix(),
		userHandler,
	)

	return mux
}
