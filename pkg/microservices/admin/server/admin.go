package server

import (
	"context"
	"errors"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/admin/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AdminServer struct {
	proto.UnimplementedAdminServer

	AdminService     *service.AdminService
	ComplaintService *service.ComplaintService
}

func NewAdminServer(admin *service.AdminService, complaint *service.ComplaintService) *AdminServer {
	return &AdminServer{
		AdminService:     admin,
		ComplaintService: complaint,
	}
}

func (as *AdminServer) AcceptComplaint(ctx context.Context, complaint *proto.Complaint) (*proto.AdminEmpty, error) {
	err := as.ComplaintService.AcceptComplaint(ctx, int(complaint.Id))
	if err != nil {
		return &proto.AdminEmpty{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.AdminEmpty{}, nil
}

func (as *AdminServer) DeleteComplaint(ctx context.Context, complaint *proto.Complaint) (*proto.AdminEmpty, error) {
	err := as.ComplaintService.DeleteComplaint(ctx, int(complaint.Id))
	if err != nil {
		return &proto.AdminEmpty{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.AdminEmpty{}, nil
}

func (as *AdminServer) GetNextComplaint(ctx context.Context, _ *proto.AdminEmpty) (*proto.Complaint, error) {
	complaint, err := as.ComplaintService.GetNextComplaint(ctx)
	if errors.Is(err, static.ErrNoData) {
		return &proto.Complaint{}, status.Error(codes.NotFound, "complaints ended")
	}
	if err != nil {
		return &proto.Complaint{}, status.Error(codes.Internal, err.Error())
	}
	createdAt, err := ptypes.TimestampProto(*complaint.CreatedAt)
	if err != nil {
		return &proto.Complaint{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.Complaint{
		Id:              int32(complaint.Id),
		ReporterUserId:  int32(complaint.ReporterUserId),
		ReportedUserId:  int32(complaint.ReportedUserId),
		ComplaintTypeId: int32(complaint.ComplaintTypeId),
		ComplaintText:   *complaint.ComplaintText,
		CreatedAt:       createdAt,
	}, nil
}

func (as *AdminServer) CreateRecommendation(ctx context.Context, rec *proto.Recommendation) (*proto.AdminEmpty, error) {
	rating := int(rec.Rating)
	_, err := as.AdminService.CreateRecommendation(
		ctx,
		model.Recommendation{
			Id:     int(rec.Id),
			UserId: int(rec.UserId),
			Rating: &rating,
		})

	if err != nil {
		return &proto.AdminEmpty{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.AdminEmpty{}, nil
}

func (as *AdminServer) CreateFeedFeedback(ctx context.Context, rec *proto.Recommendation) (*proto.AdminEmpty, error) {
	rating := int(rec.Rating)
	_, err := as.AdminService.CreateFeedFeedback(
		ctx,
		model.Recommendation{
			Id:     int(rec.Id),
			UserId: int(rec.UserId),
			Rating: &rating,
		})

	if err != nil {
		return &proto.AdminEmpty{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.AdminEmpty{}, nil
}

func (as *AdminServer) CreateFeedback(ctx context.Context, stat *proto.Feedback) (*proto.AdminEmpty, error) {
	rating := int(stat.Rating)
	_, err := as.AdminService.CreateFeedback(
		ctx,
		model.Feedback{
			Id:      int(stat.Id),
			UserId:  int(stat.UserId),
			Rating:  &rating,
			Liked:   &stat.Liked,
			NeedFix: &stat.NeedFix,
			Comment: &stat.Comment,
		})

	if err != nil {
		return &proto.AdminEmpty{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.AdminEmpty{}, nil
}

func (as *AdminServer) GetFeedbackStatistic(ctx context.Context, _ *proto.AdminEmpty) (*proto.FeedbackStatistic, error) {
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

func (as *AdminServer) GetRecommendationStatistic(ctx context.Context, _ *proto.AdminEmpty) (*proto.RecommendationStatistic, error) {
	recommendationsStat, err := as.AdminService.GetRecommendationsStatistics(ctx)
	if err != nil {
		return &proto.RecommendationStatistic{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.RecommendationStatistic{
		RecommendCount: recommendationsStat.RecommendCount,
		AvgRecommend:   recommendationsStat.AvgRecommend,
		NPS:            recommendationsStat.NPS,
	}, nil
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
