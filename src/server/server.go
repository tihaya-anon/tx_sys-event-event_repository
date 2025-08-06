package server

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/constant"
	dao_impl "github.com/tihaya-anon/tx_sys-event-event_repository/src/dao/impl"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/db"
	"github.com/tihaya-anon/tx_sys-event-event_repository/src/pb/kafka"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
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
		log.Error().Err(err).Msg("failed to listen")
		return nil, err
	}
	pool, err := pgxpool.New(context.Background(), constant.DB_URL)
	if err != nil {
		log.Error().Err(err).Msg("failed to create pool")
		return nil, err
	}
	q := db.New(pool)
	r := dao_impl.NewReader(pool)
	s := grpc.NewServer()
	handler := NewGrpcHandler(q, r)
	if handler == nil {
		log.Error().Err(err).Msg("failed to create handler")
		return nil, errors.New("failed to create handler")
	}
	kafka.RegisterEventRepositoryServer(s, handler)
	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(s, healthServer)
	return &Server{grpcServer: s, listener: lis, pool: pool, healthServer: healthServer, healthCheckService: "grpc.health.v1.Health"}, nil
}

func (s *Server) Start() error {
	log.Info().Msgf("gRPC server listening at %s", s.listener.Addr())
	s.healthServer.SetServingStatus(s.healthCheckService, healthpb.HealthCheckResponse_SERVING)
	reflection.Register(s.grpcServer)
	services := s.grpcServer.GetServiceInfo()
	if len(services) == 0 {
		log.Error().Msg("no gRPC services registered")
		return errors.New("server error")
	} else {
		for name := range services {
			log.Info().Msgf("Registered service: %s", name)
		}
	}
	if err := s.grpcServer.Serve(s.listener); err != nil {
		s.healthServer.SetServingStatus(s.healthCheckService, healthpb.HealthCheckResponse_NOT_SERVING)
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
		log.Info().Msg("gRPC server shutdown gracefully")
	case <-ctx.Done():
		log.Info().Msg("gRPC server shutdown timeout, forcing stop")
		s.grpcServer.Stop()
	}

	if s.pool != nil {
		s.pool.Close()
	}
}
