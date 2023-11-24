package server

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
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

func (fs *FeedServer) Feed(ctx context.Context, params *proto.FilterParams) (*proto.User, error) {
	nextUser, err := fs.FeedService.GetNextUser(ctx, model.FilterParams{
		UserId: int(params.UserId),
		MinAge: int(params.MinAge),
		MaxAge: int(params.MaxAge),
		Tags:   params.Tags,
	})
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
		UserGender:   modifyInt(nextUser.UserGender),
		PreferGender: modifyInt(nextUser.PreferGender),
		Description:  modifyString(nextUser.Description),
		Age:          modifyInt(nextUser.Age),
		Looking:      modifyString(nextUser.Looking),
		ImagePaths:   modifyArray(nextUser.ImagePaths),
		Education:    modifyString(nextUser.Education),
		Hobbies:      modifyString(nextUser.Hobbies),
		Birthday:     birthdayProto,
		Online:       nextUser.Online,
		Tags:         modifyArray(nextUser.Tags),
	}, nil
}

func modifyString(data *string) string {
	if data == nil {
		return ""
	}
	return *data
}

func modifyInt(data *int) int32 {
	if data == nil {
		return 0
	}
	return int32(*data)
}

func modifyArray(data *[]string) []string {
	if data == nil {
		return []string{}
	}
	var result []string
	for _, path := range *data {
		result = append(result, path)
	}
	return result
}
