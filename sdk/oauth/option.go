package oauthsdk

import "google.golang.org/grpc"

type Option func(m *manager)

func WithUnaryInterceptors(interceptors ...grpc.UnaryClientInterceptor) Option {
	return func(m *manager) {
		m.unaryInterceptors = interceptors
	}
}
