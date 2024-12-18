package sdk

import (
	"fmt"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

type GrpcClientBuilder interface {
	WithInsecure()
	WithUnaryInterceptors(interceptors ...grpc.UnaryClientInterceptor)
	WithStreamInterceptors(interceptors ...grpc.StreamClientInterceptor)
	WithKeepAliveParams(params keepalive.ClientParameters)
	GetConn(add string) (*grpc.ClientConn, error)
}

type grpcClientBuilderImpl struct {
	options []grpc.DialOption
}

func NewGrpcClientBuilder(opts ...grpc.DialOption) GrpcClientBuilder {
	return &grpcClientBuilderImpl{
		options: opts,
	}
}

func (sb *grpcClientBuilderImpl) WithInsecure() {
	sb.options = append(sb.options, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func (sb *grpcClientBuilderImpl) WithUnaryInterceptors(interceptors ...grpc.UnaryClientInterceptor) {
	sb.options = append(sb.options, grpc.WithUnaryInterceptor(middleware.ChainUnaryClient(interceptors...)))
}

func (sb *grpcClientBuilderImpl) WithStreamInterceptors(interceptors ...grpc.StreamClientInterceptor) {
	sb.options = append(sb.options, grpc.WithStreamInterceptor(middleware.ChainStreamClient(interceptors...)))
}

func (sb *grpcClientBuilderImpl) WithKeepAliveParams(params keepalive.ClientParameters) {
	keepAlive := grpc.WithKeepaliveParams(params)
	sb.options = append(sb.options, keepAlive)
}

func (sb *grpcClientBuilderImpl) GetConn(addr string) (*grpc.ClientConn, error) {
	if addr == "" {
		return nil, fmt.Errorf("grpcClient - GetConn - Missing required parameter: address is empty")
	}

	conn, err := grpc.NewClient(addr, sb.options...)
	if err != nil {
		return nil, fmt.Errorf("grpcClient - GetConn - NewClient - %v", err)
	}

	return conn, nil
}
