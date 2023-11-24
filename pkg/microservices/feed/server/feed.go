package server

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/feed/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
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

func (fs *FeedServer) Feed(ctx context.Context, userId *proto.UserIdFeed) (*proto.User, error) {
	//return &proto.User{Id: 11}, nil
	nextUser, err := fs.FeedService.GetNextUser(ctx, int(userId.Id))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
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
		UserGender:   int32(*nextUser.UserGender),
		PreferGender: int32(*nextUser.PreferGender),
		Description:  *nextUser.Description,
		Age:          int32(*nextUser.Age),
		Looking:      *nextUser.Looking,
		//ImagePaths:   *nextUser.ImagePaths,
		Education: *nextUser.Education,
		Hobbies:   *nextUser.Hobbies,
		Birthday:  birthdayProto,
		Online:    nextUser.Online,
		//Tags:         *nextUser.Tags,
	}, nil
}
