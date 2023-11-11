package services

import (
	"api_gateway_task/config"
	"api_gateway_task/genproto/integrations_service"

	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	IntegrationsService() integrations_service.IntegrationsServiceClient
}

type grpcClients struct {
	integrationsService integrations_service.IntegrationsServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {

	// Integrations Service...
	connIntegrationsService, err := grpc.Dial(
		cfg.IntegrationsServiceHost+cfg.IntegrationsGRPCPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(100*1024*1024*1024),
			grpc.MaxCallSendMsgSize(100*1024*1024*1024),
		),
		grpc.WithUnaryInterceptor(
			otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer()),
		),
		grpc.WithStreamInterceptor(
			otgrpc.OpenTracingStreamClientInterceptor(opentracing.GlobalTracer()),
		),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		integrationsService: integrations_service.NewIntegrationsServiceClient(connIntegrationsService),
	}, nil
}

func (g *grpcClients) IntegrationsService() integrations_service.IntegrationsServiceClient {
	return g.integrationsService
}
