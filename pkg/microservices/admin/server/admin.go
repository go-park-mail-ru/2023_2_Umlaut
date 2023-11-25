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
	feedbackStat, err := as.AdminService.GetFeedbackStatistics(ctx)
	if err != nil {
		return &proto.FeedbackStatistic{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.FeedbackStatistic{
		AvgRating:   feedbackStat.AvgRating,
		RatingCount: feedbackStat.RatingCount,
		LikedMap:    getProtoLikedMap(feedbackStat.LikedMap),
		NeedFixMap:  getProtoNeedFixMap(feedbackStat.NeedFixMap),
		Comments:    feedbackStat.Comments,
	}, nil
}

func (as *AdminServer) GetRecommendationStatistic(ctx context.Context, _ *proto.Empty) (*proto.RecommendationStatistic, error) {

}

func getProtoLikedMap(likedMap map[string]int32) []*proto.LikedMap {
	result := []*proto.LikedMap{}
	for key, value := range likedMap {
		result = append(result, &proto.LikedMap{Liked: key, Count: value})
	}
	return result
}

func getProtoNeedFixMap(needFixMap map[string]model.NeedFixObject) []*proto.NeedFixMap {
	result := []*proto.NeedFixMap{}
	for key, value := range needFixMap {
		result = append(
			result,
			&proto.NeedFixMap{
				NeedFix:       key,
				NeedFixObject: &proto.NeedFixObject{Count: value.Count, CommentFix: value.CommentFix}})
	}
	return result
}
