package multiservices

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

// GRPCMod ...
type GRPCMod struct {
	Port   string
	Server *grpc.Server
}

// Start ...
func (s *GRPCMod) Start() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", s.Port))
	if err != nil {
		return err
	}

	return s.Server.Serve(l)
}

// Close ...
func (s *GRPCMod) Close() error {
	s.Server.GracefulStop()
	return nil
}
