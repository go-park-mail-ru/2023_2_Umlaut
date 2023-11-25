package server

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/admin/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
)

type AdminServer struct {
	proto.UnimplementedAdminServer

	AdminService *service.AdminService
}

func NewAdminServer(feed *service.AdminService) *AdminServer {
	return &AdminServer{AdminService: feed}
}

func (fs *AdminServer) GetStatistic(ctx context.Context, _ *proto.Empty) (*proto.Statistic, error) {
	return &proto.Statistic{}, nil
}
