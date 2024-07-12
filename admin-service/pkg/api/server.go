package server

import (
	"fmt"
	"healy-admin/pkg/config"
	"net"

	"google.golang.org/grpc"
	pb "healy-admin/pkg/pb/admin"
)

type Server struct {
	server   *grpc.Server
	listener net.Listener
}

func NewGRPCServer(cfg config.Config, server pb.AdminServer) (*Server, error) {

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, err
	}

	newServer := grpc.NewServer()
	pb.RegisterAdminServer(newServer, server)

	return &Server{
		server:   newServer,
		listener: lis,
	}, nil
}

func (c *Server) Start() error {
	fmt.Println("grpc server listening on port :50053")
	return c.server.Serve(c.listener)
}
