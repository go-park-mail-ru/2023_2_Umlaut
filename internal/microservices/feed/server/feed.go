package server

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/constants"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/model/core"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/utils"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/microservices/feed/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FeedServer struct {
	proto.UnimplementedFeedServer

	FeedService *service.FeedService
}

func NewFeedServer(feed *service.FeedService) *FeedServer {
	return &FeedServer{FeedService: feed}
}

func (fs *FeedServer) Feed(ctx context.Context, params *proto.FilterParams) (*proto.FeedData, error) {
	feed, err := fs.FeedService.GetNextUser(ctx, core.FilterParams{
		UserId: int(params.UserId),
		MinAge: int(params.MinAge),
		MaxAge: int(params.MaxAge),
		Tags:   params.Tags,
	})
	if errors.Is(err, constants.ErrBannedUser) {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	birthdayProto := timestamppb.New(*feed.User.Birthday)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.FeedData{
		User: &proto.User{
			Id:           int32(feed.User.Id),
			Name:         feed.User.Name,
			UserGender:   utils.ModifyInt(feed.User.UserGender),
			PreferGender: utils.ModifyInt(feed.User.PreferGender),
			Description:  utils.ModifyString(feed.User.Description),
			Age:          utils.ModifyInt(feed.User.Age),
			Looking:      utils.ModifyString(feed.User.Looking),
			ImagePaths:   utils.ModifyArray(feed.User.ImagePaths),
			Education:    utils.ModifyString(feed.User.Education),
			Hobbies:      utils.ModifyString(feed.User.Hobbies),
			Birthday:     birthdayProto,
			Online:       feed.User.Online,
			Tags:         utils.ModifyArray(feed.User.Tags),
		},
		LikeCounter: int32(feed.LikeCounter),
	}, nil
}
