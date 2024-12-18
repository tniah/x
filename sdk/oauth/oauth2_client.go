package oauthsdk

import (
	"context"
	pbv1 "github.com/tniah/iam-grpc/oauth2/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ClientManager interface {
	List(context.Context, *pbv1.ListOAuth2ClientsRequest) (*pbv1.ListOAuth2ClientsResponse, error)
	Create(ctx context.Context, req *pbv1.CreateOAuth2ClientRequest) (*pbv1.OAuth2Client, error)
	Get(ctx context.Context, req *pbv1.GetOAuth2ClientRequest) (*pbv1.OAuth2Client, error)
	Update(ctx context.Context, req *pbv1.UpdateOAuth2ClientRequest) (*pbv1.OAuth2Client, error)
	Delete(ctx context.Context, req *pbv1.DeleteOAuth2ClientRequest) (*emptypb.Empty, error)
}

type clientManagerImpl struct {
	c pbv1.OAuth2ClientServiceClient
}

func newOAuth2ClientManager(conn *grpc.ClientConn) ClientManager {
	c := pbv1.NewOAuth2ClientServiceClient(conn)
	return &clientManagerImpl{c: c}
}

func (m *clientManagerImpl) List(ctx context.Context, req *pbv1.ListOAuth2ClientsRequest) (*pbv1.ListOAuth2ClientsResponse, error) {
	return m.c.ListOAuth2Clients(ctx, req)
}

func (m *clientManagerImpl) Create(ctx context.Context, req *pbv1.CreateOAuth2ClientRequest) (*pbv1.OAuth2Client, error) {
	return m.c.CreateOAuth2Client(ctx, req)
}

func (m *clientManagerImpl) Get(ctx context.Context, req *pbv1.GetOAuth2ClientRequest) (*pbv1.OAuth2Client, error) {
	return m.c.GetOAuth2Client(ctx, req)
}

func (m *clientManagerImpl) Update(ctx context.Context, req *pbv1.UpdateOAuth2ClientRequest) (*pbv1.OAuth2Client, error) {
	return m.c.UpdateOAuth2Client(ctx, req)
}

func (m *clientManagerImpl) Delete(ctx context.Context, req *pbv1.DeleteOAuth2ClientRequest) (*emptypb.Empty, error) {
	return m.c.DeleteOAuth2Client(ctx, req)
}
