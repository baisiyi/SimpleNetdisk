package main

import (
	"google.golang.org/grpc"
	"simpleNetdisk/gateway"
	pb "simpleNetdisk/protoc"
	"simpleNetdisk/util/config"
)

func newServerWithConfig(opt ...grpc.ServerOption) *grpc.Server {
	cfg, err := config.LoadConfig(config.DefaultConfigPath)
	if err != nil {
		panic("load config fail: " + err.Error())
	}
	return grpc.NewServer()
}
func main() {
	s := newServerWithConfig()
	pb.RegisterFileMgtServer(s, gateway.NewFileMgtServiceImpl())
}
