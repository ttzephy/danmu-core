package server

import (
	"danmu-core/generated/api"
	"danmu-core/internal/service"
	"danmu-core/logger"
	"danmu-core/setting"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type RPCServer struct {
	server *grpc.Server
}

func NewRPCServer() *RPCServer {
	// 创建 gRPC 服务器选项
	opts := []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: time.Duration(setting.RpcSetting.ConnectTimeout) * time.Second,
			MaxConnectionAge:  time.Duration(setting.RpcSetting.ConnectTimeout) * time.Second,
			Time:              time.Duration(setting.RpcSetting.ConnectTimeout) * time.Second,
			Timeout:           time.Duration(setting.RpcSetting.ConnectTimeout) * time.Second,
		}),
		grpc.MaxConcurrentStreams(setting.RpcSetting.MaxOpenConns),
	}

	// 创建 gRPC 服务器
	server := grpc.NewServer(opts...)

	return &RPCServer{
		server: server,
	}
}

func (s *RPCServer) Start() error {
	// 注册服务
	api.RegisterLiveServiceServer(s.server, service.NewLiveServer())

	// 创建监听器
	addr := fmt.Sprintf("%s:%s", setting.RpcSetting.Host, setting.RpcSetting.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	logger.Info().Str("addr", addr).Msg("starting gRPC server")

	// 启动服务器
	if err := s.server.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}

func (s *RPCServer) Stop() {
	if s.server != nil {
		logger.Info().Msg("stopping gRPC server")
		s.server.GracefulStop()
	}
}
