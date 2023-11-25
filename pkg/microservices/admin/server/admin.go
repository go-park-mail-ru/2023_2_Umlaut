package server

import (
	"context"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/admin/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AdminServer struct {
	proto.UnimplementedAdminServer

	AdminService *service.AdminService
}

func NewAdminServer(feed *service.AdminService) *AdminServer {
	return &AdminServer{AdminService: feed}
}

func (as *AdminServer) CreateRecommendation(ctx context.Context, rec *proto.Recommendation) (*proto.Empty, error) {
	recommend := int(rec.Recommend)
	_, err := as.AdminService.CreateRecommendation(
		ctx,
		model.Recommendation{
			Id:        int(rec.Id),
			UserId:    int(rec.UserId),
			Recommend: &recommend,
			Show:      rec.Show,
		})

	if err != nil {
		return &proto.Empty{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.Empty{}, nil
}

func (as *AdminServer) CreateFeedFeedback(ctx context.Context, rec *proto.Recommendation) (*proto.Empty, error) {
	recommend := int(rec.Recommend)
	_, err := as.AdminService.CreateFeedFeedback(
		ctx,
		model.Recommendation{
			Id:        int(rec.Id),
			UserId:    int(rec.UserId),
			Recommend: &recommend,
			Show:      rec.Show,
		})

	if err != nil {
		return &proto.Empty{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.Empty{}, nil
}

func (as *AdminServer) CreateFeedback(ctx context.Context, stat *proto.Feedback) (*proto.Empty, error) {
	rating := int(stat.Rating)
	_, err := as.AdminService.CreateFeedback(
		ctx,
		model.Feedback{
			Id:         int(stat.Id),
			UserId:     int(stat.UserId),
			Rating:     &rating,
			Liked:      &stat.Liked,
			NeedFix:    &stat.NeedFix,
			CommentFix: &stat.CommentFix,
			Comment:    &stat.Comment,
			Show:       stat.Show,
		})

	if err != nil {
		return &proto.Empty{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.Empty{}, nil
}

func (as *AdminServer) GetFeedbackStatistic(ctx context.Context, _ *proto.Empty) (*proto.FeedbackStatistic, error) {

}

func (as *AdminServer) GetRecommendationStatistic(ctx context.Context, _ *proto.Empty) (*proto.RecommendationStatistic, error) {

}
