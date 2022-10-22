package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"simpleNetdisk/gateway"
	pb "simpleNetdisk/protoc"
	"simpleNetdisk/util/config"
)

var (
	DefaultIP      = "0.0.0.0"
	DefaultPort    = 80
	DefaultNetwork = "tcp"
)

func init() {
	cfg, err := config.LoadConfig(config.DefaultConfigPath)
	if err != nil {
		panic("load config fail: " + err.Error())
	}
	if cfg == nil || cfg.Service.IP == "" {
		cfg = &config.Config{}
		cfg.Service.IP = DefaultIP
		cfg.Service.Port = DefaultPort
		cfg.Service.Network = DefaultNetwork
	}
	config.SetGlobalConfig(cfg)
}

func newListener(opt ...grpc.ServerOption) net.Listener {
	network := config.GetServiceNetWork()
	address := config.GetServiceAddress()
	listener, err := net.Listen(network, address)
	if err != nil {
		panic("load config fail: " + err.Error())
	}

	return listener
}
func main() {
	s := grpc.NewServer()
	pb.RegisterFileMgtServer(s, gateway.NewFileMgtServiceImpl())

	if err := s.Serve(newListener()); err != nil {
		log.Fatal(err)
	}
}
