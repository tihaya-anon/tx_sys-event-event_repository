package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb/kafka"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
	listener   net.Listener
}

func NewServer(port int) (*Server, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}
	s := grpc.NewServer()
	// TODO mock db required
	kafka.RegisterEventRepositoryServer(s, newGrpcHandler(nil, nil))
	return &Server{grpcServer: s, listener: lis}, nil
}

func (s *Server) Start() {
	log.Printf("gRPC server listening at %s", s.listener.Addr())
	if err := s.grpcServer.Serve(s.listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *Server) Shutdown(ctx context.Context) {
	done := make(chan struct{})
	go func() {
		s.grpcServer.GracefulStop()
		close(done)
	}()

	select {
	case <-done:
		log.Println("gRPC server shutdown gracefully")
	case <-ctx.Done():
		log.Println("gRPC server shutdown timeout, forcing stop")
		s.grpcServer.Stop()
	}
}
