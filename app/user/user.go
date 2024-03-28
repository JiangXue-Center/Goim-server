package main

import (
	"Goim-server/app/user/internal/config"
	accountservice "Goim-server/app/user/internal/server/accountservice"
	callbackservice "Goim-server/app/user/internal/server/callbackservice"
	infoservice "Goim-server/app/user/internal/server/infoservice"
	"Goim-server/app/user/internal/svc"
	"Goim-server/common/pb"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterAccountServiceServer(grpcServer, accountservice.NewAccountServiceServer(ctx))
		pb.RegisterCallbackServiceServer(grpcServer, callbackservice.NewCallbackServiceServer(ctx))
		pb.RegisterInfoServiceServer(grpcServer, infoservice.NewInfoServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	logx.Info("staring rpc server at %s...\n", c.ListenOn)
	s.Start()
}
