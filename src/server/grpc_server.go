package server

import (
	"fmt"
	"net"

	"src/config"
	pb "src/server/proto"

	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	pb.UnimplementedSpotifyCurateServer

	server *grpc.Server
	status health.HealthCheckResponse_ServingStatus
	config *config.Config
	logger *zap.Logger
}

func NewGrpcServer(config *config.Config, logger *zap.Logger) *GrpcServer {
	return &GrpcServer{
		server: grpc.NewServer(),
		config: config,
		logger: logger,
	}
}

func (g *GrpcServer) logAPICallStart(apiEndpoint apiEndpoint) {
	g.logger.Info("calling endpoint", zap.String("endpoint", string(apiEndpoint)))
}

func (g *GrpcServer) logAPICallSuccess(apiEndpoint apiEndpoint) {
	g.logger.Info("successfully executed endpoint", zap.String("endpoint", string(apiEndpoint)))
}

func (g *GrpcServer) logError(apiEndpoint apiEndpoint, err error) {
	g.logger.Error("error encountered during call", zap.String("endpoint", string(apiEndpoint)), zap.String("error", err.Error()))
}

func (g *GrpcServer) Run() error {
	health.RegisterHealthServer(g.server, g)
	pb.RegisterSpotifyCurateServer(g.server, g)

	lis, err := net.Listen(g.config.GRPC.Network, fmt.Sprintf("%s:%d", g.config.GRPC.Host, g.config.GRPC.Port))
	if err != nil {
		return err
	}

	g.logger.Info("starting grpc server", zap.String("network", g.config.GRPC.Network), zap.String("host", g.config.GRPC.Host), zap.Int("port", g.config.GRPC.Port))

	return g.server.Serve(lis)
}

func (g *GrpcServer) Check(ctx context.Context, in *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	g.logger.Info("Received healthcheck request", zap.String("request", in.String()))

	return &health.HealthCheckResponse{Status: g.status}, nil
}

func (g *GrpcServer) Watch(in *health.HealthCheckRequest, _ health.Health_WatchServer) error {
	g.logger.Debug("Received healthwatch request", zap.String("request", in.String()))

	return status.Error(codes.Unimplemented, "unimplemented")
}

func (g *GrpcServer) Stop() {
	g.status = health.HealthCheckResponse_NOT_SERVING

	// GracefulStop blocks incoming RPC calls and blocks until all RPC calls are completed
	g.server.GracefulStop()
}
