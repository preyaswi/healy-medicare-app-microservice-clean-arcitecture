package server

import (
	"doctor-service/pkg/config"
	"doctor-service/pkg/pb"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	server  *grpc.Server
	listner net.Listener
}

func NewGRPCServer(cfg config.Config, server pb.DoctorServer) (*Server, error) {
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, err
	}
	newServer := grpc.NewServer()
	pb.RegisterDoctorServer(newServer, server)
	return &Server{
		server:  newServer,
		listner: lis,
	}, nil
}
func (c *Server) Start() error {
	fmt.Println("grpc server listening on port :50052")
	return c.server.Serve(c.listner)
}
