package server

import (
	pb "github.com/muxi-Infra/muxi-micro/example/transport/grpc/proto"
	mgrpc "github.com/muxi-Infra/muxi-micro/pkg/transport/grpc"
	"github.com/muxi-Infra/muxi-micro/pkg/transport/grpc/registry/etcd"
	"google.golang.org/grpc"
)

// 和proto一起生成的服务端模板的大致形式
func HelloServer(srv pb.HelloServiceServer) (*mgrpc.GRPCServer, error) {
	center, err := etcd.NewEtcdRegistry()
	if err != nil {
		return nil, err
	}

	server := mgrpc.NewGRPCServer(
		mgrpc.WithName("test"),
		mgrpc.WithRegistrationCenter(center),
	)

	server.ProtoRegister(func(s *grpc.Server) {
		pb.RegisterHelloServiceServer(s, srv)
	})

	return server, nil
}
