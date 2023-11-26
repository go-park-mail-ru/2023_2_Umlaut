package server

import (
	"context"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/feed/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/go-park-mail-ru/2023_2_Umlaut/utils"
	"github.com/golang/protobuf/ptypes"
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

func (fs *FeedServer) Feed(ctx context.Context, params *proto.FilterParams) (*proto.User, error) {
	nextUser, err := fs.FeedService.GetNextUser(ctx, model.FilterParams{
		UserId: int(params.UserId),
		MinAge: int(params.MinAge),
		MaxAge: int(params.MaxAge),
		Tags:   params.Tags,
	})
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	birthdayProto, err := ptypes.TimestampProto(*nextUser.Birthday)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.User{
		Id:           int32(nextUser.Id),
		Name:         nextUser.Name,
		Mail:         nextUser.Mail,
		PasswordHash: nextUser.PasswordHash,
		Salt:         nextUser.Salt,
		UserGender:   utils.ModifyInt(nextUser.UserGender),
		PreferGender: utils.ModifyInt(nextUser.PreferGender),
		Description:  utils.ModifyString(nextUser.Description),
		Age:          utils.ModifyInt(nextUser.Age),
		Looking:      utils.ModifyString(nextUser.Looking),
		ImagePaths:   utils.ModifyArray(nextUser.ImagePaths),
		Education:    utils.ModifyString(nextUser.Education),
		Hobbies:      utils.ModifyString(nextUser.Hobbies),
		Birthday:     birthdayProto,
		Online:       nextUser.Online,
		Tags:         utils.ModifyArray(nextUser.Tags),
	}, nil
}
