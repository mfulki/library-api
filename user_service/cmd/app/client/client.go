package client

import (
	"user-service/config/env"
	"user-service/internal/apperror"
	"user-service/pkg/llog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	BookService     *grpc.ClientConn
	AuthorService   *grpc.ClientConn
	CategoryService *grpc.ClientConn
}

func NewGRPCClient(gRPCCfg *env.GRPCConfig) (*GRPCClient, error) {
	bookServiceConn, err := grpc.NewClient(gRPCCfg.BookServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		llog.Fatal(err)
		return nil, apperror.Wrap(err)
	}

	authorServiceConn, err := grpc.NewClient(gRPCCfg.AuthorServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		llog.Fatal(err)
		return nil, apperror.Wrap(err)
	}

	categoryServiceConn, err := grpc.NewClient(gRPCCfg.CategoryServicePort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		llog.Fatal(err)
		return nil, apperror.Wrap(err)
	}
	return &GRPCClient{
		BookService:     bookServiceConn,
		AuthorService:   authorServiceConn,
		CategoryService: categoryServiceConn,
	}, nil
}
