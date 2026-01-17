package micro

import (
	pb "github.com/muxi-Infra/muxi-micro/example/transport/grpc/proto"
	"github.com/muxi-Infra/muxi-micro/pkg/transport/grpc"
	"github.com/muxi-Infra/muxi-micro/pkg/transport/grpc/discovery/etcd"
)

// 和proto一起生成的客户端模板的大致形式
func HelloClient() (pb.HelloServiceClient, func(), error) {
	center, err := etcd.NewEtcdDiscovery()
	if err != nil {
		return nil, nil, err
	}

	client, err := grpc.NewGRPCClient(
		grpc.WithServiceDiscovery(center),
		grpc.WithDiscoveryName("test"),
	)
	if err != nil {
		return nil, nil, err
	}

	c := pb.NewHelloServiceClient(client.Conn())
	return c, func() { _ = client.Close() }, nil
}
