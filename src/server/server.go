package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	constant_postgre "github.com/tihaya-anon/tx_sys-event-event_repository/src/constant/postgre"
	dao_impl "github.com/tihaya-anon/tx_sys-event-event_repository/src/dao/impl"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/db"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/kafka_bridge"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb/kafka"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
	listener   net.Listener
}

func NewServer(port int, c *kafka_bridge.APIClient) (*Server, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.New(context.Background(), constant_postgre.DB_URL)
	if err != nil {
		return nil, err
	}
	defer pool.Close()
	q := db.New(pool)
	r := dao_impl.NewReader(pool)
	s := grpc.NewServer()
	kafka.RegisterEventRepositoryServer(s, NewGrpcHandler(q, r, c))
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
