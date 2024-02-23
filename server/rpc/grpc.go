package rpc

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"sokwva/acfun/dougaInfo/conf"
	"sokwva/acfun/dougaInfo/fetch"
	rpcproto "sokwva/acfun/dougaInfo/server/protoLib"
	"sokwva/acfun/dougaInfo/server/protoLib/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	rpcproto.UnimplementedGetServer
}

func (me *server) GetDougaInfo(ctx context.Context, in *rpcproto.Acid) (*rpcproto.DougaInfo, error) {
	fmt.Println(in.Acid)
	x, err := fetch.GetVideoInfo(in.Acid)
	var converter generated.ConverterImpl
	convtRes := converter.Convert(*x)
	if err != nil {
		return &rpcproto.DougaInfo{}, nil
	}
	return &convtRes, nil
}

func Start() {
	lis, err := net.Listen("tcp", conf.Conf.Server.Address+":"+conf.Conf.Server.RPCPort)
	if err != nil {
		slog.Error("failed to listen: ", err)
	}

	s := grpc.NewServer()
	rpcproto.RegisterGetServer(s, &server{})
	reflection.Register(s)

	// 启动grpc服务
	if err := s.Serve(lis); err != nil {
		slog.Error("failed to serve: %v", err)
	}
}
