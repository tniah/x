package oauthsdk

import (
	"github.com/tniah/x/sdk"
	"google.golang.org/grpc"
)

type Manager interface {
	OAuthClient() ClientManager
	Channel() *grpc.ClientConn
	Clean() error
}

type manager struct {
	clientManager     ClientManager
	channel           *grpc.ClientConn
	unaryInterceptors []grpc.UnaryClientInterceptor
}

func New(targetHost string, opts ...Option) (Manager, error) {
	m := &manager{}
	for _, opt := range opts {
		opt(m)
	}

	sb := sdk.NewGrpcClientBuilder()
	sb.WithInsecure()
	sb.WithUnaryInterceptors(m.unaryInterceptors...)

	channel, err := sb.GetConn(targetHost)
	if err != nil {
		return nil, err
	}

	m.channel = channel
	m.clientManager = newOAuth2ClientManager(m.channel)
	return m, nil
}

func (m *manager) OAuthClient() ClientManager {
	return m.clientManager
}

func (m *manager) Channel() *grpc.ClientConn {
	return m.channel
}

func (m *manager) Clean() error {
	return m.channel.Close()
}
