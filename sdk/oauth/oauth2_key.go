package oauthsdk

import (
	"context"
	pbv1 "github.com/tniah/iam-grpc/oauth2/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type KeyManager interface {
	List(ctx context.Context, req *pbv1.ListOAuth2KeysRequest) (*pbv1.ListOAuth2KeysResponse, error)
	Generate(ctx context.Context, req *pbv1.CreateOAuth2KeyRequest) (*pbv1.OAuth2Key, error)
	Get(ctx context.Context, req *pbv1.GetOAuth2KeyRequest) (*pbv1.OAuth2Key, error)
	Update(ctx context.Context, req *pbv1.UpdateOAuth2KeyRequest) (*pbv1.OAuth2Key, error)
	Delete(ctx context.Context, req *pbv1.DeleteOAuth2KeyRequest) (*emptypb.Empty, error)
}

type keyManagerImpl struct {
	c pbv1.OAuth2KeyServiceClient
}

func newOAuth2KeyManager(conn *grpc.ClientConn) KeyManager {
	c := pbv1.NewOAuth2KeyServiceClient(conn)
	return &keyManagerImpl{c: c}
}

func (m *keyManagerImpl) List(ctx context.Context, req *pbv1.ListOAuth2KeysRequest) (*pbv1.ListOAuth2KeysResponse, error) {
	return m.c.ListOAuth2Keys(ctx, req)
}

func (m *keyManagerImpl) Generate(ctx context.Context, req *pbv1.CreateOAuth2KeyRequest) (*pbv1.OAuth2Key, error) {
	return m.c.CreateOAuth2Key(ctx, req)
}

func (m *keyManagerImpl) Get(ctx context.Context, req *pbv1.GetOAuth2KeyRequest) (*pbv1.OAuth2Key, error) {
	return m.c.GetOAuth2Key(ctx, req)
}

func (m *keyManagerImpl) Update(ctx context.Context, req *pbv1.UpdateOAuth2KeyRequest) (*pbv1.OAuth2Key, error) {
	return m.c.UpdateOAuth2Key(ctx, req)
}

func (m *keyManagerImpl) Delete(ctx context.Context, req *pbv1.DeleteOAuth2KeyRequest) (*emptypb.Empty, error) {
	return m.c.DeleteOAuth2Key(ctx, req)
}
