package server

import (
	"Goim-server/app/user/internal/svc"
	"Goim-server/common/pb"
)

type InfoServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedInfoServiceServer
}

func NewInfoServiceServer(svcCtx *svc.ServiceContext) *InfoServiceServer {
	return &InfoServiceServer{
		svcCtx: svcCtx,
	}
}
