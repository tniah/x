package server

import (
	"fmt"
	middware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type ShutdownHook func()
type BootstrapHook func()

type GrpcServer interface {
	Start(addr string, port uint16, bootstrapHook BootstrapHook) error
	Shutdown(shutdownHook ShutdownHook)
	RegisterService(svc ServiceServer)
	GetListener() net.Listener
}

type GrpcServerBuilder struct {
	options           []grpc.ServerOption
	enabledReflection bool
}

type ServiceServer interface {
	RegisterWithServer(*grpc.Server)
}

type grpcServer struct {
	server   *grpc.Server
	listener net.Listener
}

func (sb *GrpcServerBuilder) AddOption(opts ...grpc.ServerOption) {
	sb.options = append(sb.options, opts...)
}

func (sb *GrpcServerBuilder) EnableReflection(enabled bool) {
	sb.enabledReflection = enabled
}

func (sb *GrpcServerBuilder) SetUnaryInterceptors(interceptors []grpc.UnaryServerInterceptor) {
	chain := grpc.UnaryInterceptor(middware.ChainUnaryServer(interceptors...))
	sb.AddOption(chain)
}

func (sb *GrpcServerBuilder) SetStreamInterceptors(interceptors []grpc.StreamServerInterceptor) {
	chain := grpc.StreamInterceptor(middware.ChainStreamServer(interceptors...))
	sb.AddOption(chain)
}

func (sb *GrpcServerBuilder) Build() GrpcServer {
	srv := grpc.NewServer(sb.options...)

	if sb.enabledReflection {
		reflection.Register(srv)
	}
	return &grpcServer{server: srv}
}

func (s *grpcServer) Start(addr string, port uint16, bootstrapHook BootstrapHook) error {
	var err error
	s.listener, err = net.Listen("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		return fmt.Errorf("grpcServer - Serve - net.Listen: %w", err)
	}

	if bootstrapHook != nil {
		bootstrapHook()
	}

	err = s.server.Serve(s.listener)
	if err != nil {
		return fmt.Errorf("grpcServer - Start - s.server.Serve: %w", err)
	}

	return nil
}

func (s *grpcServer) Shutdown(shutdownHook ShutdownHook) {
	if s.server != nil {
		s.server.GracefulStop()
	}

	if s.listener != nil {
		s.listener.Close()
	}

	if shutdownHook != nil {
		shutdownHook()
	}
}

func (s *grpcServer) RegisterService(svc ServiceServer) {
	svc.RegisterWithServer(s.server)
}

func (s *grpcServer) GetListener() net.Listener {
	return s.listener
}
