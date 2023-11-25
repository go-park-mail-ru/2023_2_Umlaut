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

func (fs *AdminServer) CreateRecommendation(context.Context, *proto.Recommendation) (*proto.Empty, error) {
	
}

func (fs *AdminServer) CreateStatistic(context.Context, *proto.Statistic) (*proto.Empty, error) {

}

func (fs *AdminServer) GetAllStatistic(context.Context, *proto.Empty) (*proto.Statistic, error) {

}
