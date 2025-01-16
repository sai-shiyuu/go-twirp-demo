package loginserver

import (
	"context"
	pb "go-web/rpc/login"
	"go-web/service/login"
)

// Server implements the Login service
type LoginServer struct{}

func (s *LoginServer) Login(ctx context.Context, req *pb.LoginReq) (u *pb.LoginRes, err error) {
	token, err := login.Login(req.Id, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.LoginRes{Token: token}, nil
}

func (s *LoginServer) Register(ctx context.Context, req *pb.RegisterReq) (u *pb.RegisterRes, err error) {
	num, err := login.Register(req.Name, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterRes{Code: num}, nil
}
