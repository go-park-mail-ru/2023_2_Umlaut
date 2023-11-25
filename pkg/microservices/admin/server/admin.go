package server

import (
	"context"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/admin/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/golang/protobuf/ptypes"
)

type AdminServer struct {
	proto.UnimplementedAdminServer

	AdminService *service.AdminService
}

func NewAdminServer(feed *service.AdminService) *AdminServer {
	return &AdminServer{AdminService: feed}
}

func (as *AdminServer) CreateRecommendation(context.Context, *proto.Recommendation) (*proto.Empty, error) {

}

func (as *AdminServer) CreateStatistic(ctx context.Context, stat *proto.Statistic) (*proto.Empty, error) {
	rating := int(stat.Rating)
	as.AdminService.CreateStatistic(
		ctx, 
		model.Statistic{
			Id: int(stat.Id),
			UserId: int(stat.UserId),
			Rating: &rating,
			Liked: &stat.Liked,
			NeedFix: &stat.NeedFix,
			CommentFix: &stat.CommentFix,
			Comment: &stat.Comment,
	})
}

func (as *AdminServer) GetAllStatistic(context.Context, *proto.Empty) (*proto.Statistic, error) {

}
