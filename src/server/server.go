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
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb/kafka"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type Server struct {
	grpcServer         *grpc.Server
	healthServer       *health.Server
	listener           net.Listener
	pool               *pgxpool.Pool
	healthCheckService string
}

func NewServer(port int) (*Server, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.New(context.Background(), constant_postgre.DB_URL)
	if err != nil {
		return nil, err
	}
	q := db.New(pool)
	r := dao_impl.NewReader(pool)
	s := grpc.NewServer()
	kafka.RegisterEventRepositoryServer(s, NewGrpcHandler(q, r))
	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(s, healthServer)
	return &Server{grpcServer: s, listener: lis, pool: pool, healthServer: healthServer, healthCheckService: "grpc.health.v1.Health"}, nil
}

func (s *Server) Start() error {
	log.Printf("gRPC server listening at %s", s.listener.Addr())
	s.healthServer.SetServingStatus(s.healthCheckService, healthpb.HealthCheckResponse_SERVING)
	if err := s.grpcServer.Serve(s.listener); err != nil {
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) {
	if s.healthServer != nil {
		s.healthServer.SetServingStatus(s.healthCheckService, healthpb.HealthCheckResponse_NOT_SERVING)
	}
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

	if s.pool != nil {
		s.pool.Close()
	}
}
