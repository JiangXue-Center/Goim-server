package server

import (
	"Goim-server/app/user/internal/svc"
	"Goim-server/common/pb"
)

type AccountServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedAccountServiceServer
}

func NewAccountServiceServer(svcCtx *svc.ServiceContext) *AccountServiceServer {
	return &AccountServiceServer{
		svcCtx: svcCtx,
	}
}
