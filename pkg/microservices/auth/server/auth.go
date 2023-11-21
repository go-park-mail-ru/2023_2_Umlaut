package server

import (
	"context"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/auth/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	proto.UnimplementedAuthorizationServer

	Authorization *service.AuthService
}

func NewAuthServer(auth *service.AuthService) *AuthServer {
	return &AuthServer{Authorization: auth}
}

func (as *AuthServer) LogOut(ctx context.Context, cookie *proto.Cookie) (*proto.Empty, error) {
	if err := as.Authorization.DeleteCookie(ctx, cookie.Cookie); err != nil {
		return &proto.Empty{}, status.Error(codes.Internal, "invalid cookie deletion")
	}

	return &proto.Empty{}, nil
}

func (as *AuthServer) SignIn(ctx context.Context, input *proto.SignInInput) (*proto.Cookie, error) {
	if input.Mail == "" || input.Password == "" {
		return &proto.Cookie{}, status.Error(codes.InvalidArgument, "missing required fields")
	}

	user, err := as.Authorization.GetUser(ctx, input.Mail, input.Password)
	if err != nil {
		return &proto.Cookie{}, status.Error(codes.Unauthenticated, "invalid mail or password")
	}

	SID, err := as.Authorization.GenerateCookie(ctx, user.Id)
	if err != nil {
		return &proto.Cookie{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.Cookie{Cookie: SID}, nil
}

func (as *AuthServer) SignUp(ctx context.Context, input *proto.SignUpInput) (*proto.UserId, error) {
	if input.Name == "" || input.Mail == "" || input.Password == "" {
		return &proto.UserId{}, status.Error(codes.InvalidArgument, "missing required fields")
	}

	user := model.User{Name: input.Name, Mail: input.Mail, PasswordHash: input.Password}

	id, err := as.Authorization.CreateUser(ctx, user)
	if err != nil {
		return &proto.UserId{}, status.Error(codes.InvalidArgument, err.Error())
	}

	SID, err := as.Authorization.GenerateCookie(ctx, id)
	if err != nil {
		return &proto.UserId{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.UserId{Id: int64(id), Cookie: &proto.Cookie{Cookie: SID}}, nil
}
