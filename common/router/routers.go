package router

import (
	"go-web/internal/haberdasherserver"
	"go-web/internal/userserver"
	"go-web/rpc/haberdasher"
	"go-web/rpc/user"
	"net/http"

	"github.com/twitchtv/twirp"
)

func Routers() *http.ServeMux {
	mux := http.NewServeMux()

	hbserver := &haberdasherserver.Server{} // implements Haberdasher interface
	haberdasherHandler := haberdasher.
		NewHaberdasherServer(
			hbserver,
			twirp.WithServerPathPrefix(""))
	mux.Handle(
		haberdasherHandler.PathPrefix(),
		haberdasherHandler,
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
