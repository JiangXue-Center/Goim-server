package server

import (
	"Goim-server/app/user/internal/svc"
	"Goim-server/common/pb"
)

type CallbackServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedCallbackServiceServer
}

func NewCallbackServiceServer(svcCtx *svc.ServiceContext) *CallbackServiceServer {
	return &CallbackServiceServer{
		svcCtx: svcCtx,
	}
}
