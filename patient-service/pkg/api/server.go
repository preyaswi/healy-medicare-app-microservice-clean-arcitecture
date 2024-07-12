package server

import (
	"fmt"
	"net"
	"patient-service/pkg/config"
	"patient-service/pkg/pb"

	"google.golang.org/grpc"
)

type Server struct {
	server  *grpc.Server
	listner net.Listener
}

func NewGRPCServer(cfg config.Config, server pb.PatientServer) (*Server, error) {
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, err
	}
	newServer := grpc.NewServer()
	pb.RegisterPatientServer(newServer, server)
	return &Server{
		server:  newServer,
		listner: lis,
	}, nil
}
func (c *Server) Start() error {
	fmt.Println("grpc server listening on port :50051")
	return c.server.Serve(c.listner)
}
