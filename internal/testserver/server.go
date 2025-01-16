package testserver

import (
	"context"

	utils "go-web/common/utils"
	pb "go-web/rpc/test"
)

// Server implements the testserver service
type Server struct{}

func (s *Server) SHA256(ctx context.Context, req *pb.SHA256Req) (res *pb.SHA256Res, err error) {
	return &pb.SHA256Res{Sha256: utils.GenerateSHA256Hash(req.Password)}, nil
}
