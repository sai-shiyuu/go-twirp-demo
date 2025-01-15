package userserver

import (
	"context"
	pb "go-web/rpc/user"
	us "go-web/service/user"
)

// Server implements the User service
type UserServer struct{}

func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserReq) (u *pb.UserRes, err error) {
	user, err := us.GetUser(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.UserRes{Id: user.Id,
		Name:  user.Name,
		Ctime: user.Ctime,
		Mtime: user.Mtime}, nil
}

func (s *UserServer) GetUsers(ctx context.Context, req *pb.GetUsersReq) (u *pb.GetUsersRes, err error) {
	users, err := us.GetUsers(req.Name)
	if err != nil {
		return nil, err
	}
	res := &pb.GetUsersRes{}
	for _, user := range users {
		res.Users = append(res.Users, &pb.UserRes{Id: user.Id,
			Name:  user.Name,
			Ctime: user.Ctime,
			Mtime: user.Mtime})
	}
	return res, nil
}
