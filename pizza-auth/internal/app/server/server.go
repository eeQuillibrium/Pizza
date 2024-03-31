package server

import (
	"context"

	nikita_auth1 "github.com/eeQuillibrium/protos/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// service (in services/auth)
type Auth interface {
	Login(
		ctx context.Context,
		phone string,
		pass string,
	) (string, error)
	Register(
		ctx context.Context,
		phone string,
		pass string,
	) (int64, error)
	IsAdmin(
		ctx context.Context,
		UserId int64,
	) (bool, error)
}

// handler
type serverAPI struct {
	nikita_auth1.UnimplementedAuthServer
	auth Auth
}

func Register(gRPC *grpc.Server, auth Auth) {
	nikita_auth1.RegisterAuthServer(gRPC, &serverAPI{auth: auth})
}

func (s *serverAPI) Login(
	ctx context.Context,
	req *nikita_auth1.LoginRequest,
) (*nikita_auth1.LoginResponse, error) {

	if err := validateLogin(req.GetAppId(), req.GetPhone(), req.GetPass()); err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	jwt, err := s.auth.Login(ctx, req.GetPhone(), req.GetPass())
	if err != nil {
		return nil, err
	}

	resp := &nikita_auth1.LoginResponse{Token: jwt}

	return resp, nil
}
func (s *serverAPI) Register(
	ctx context.Context,
	req *nikita_auth1.RegRequest,
) (*nikita_auth1.RegResponse, error) {

	if err := validateRegister(req.GetPhone(), req.GetPass()); err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	userId, err := s.auth.Register(ctx, req.GetPhone(), req.GetPass())
	if err != nil {
		return nil, err
	}

	return &nikita_auth1.RegResponse{UserId: userId}, nil
}
func (s *serverAPI) IsAdmin(
	ctx context.Context,
	req *nikita_auth1.IsAdminRequest,
) (*nikita_auth1.IsAdminResponse, error) {

	if err := validateIsAdmin(req.GetUserId()); err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	if _, err := s.auth.IsAdmin(ctx, req.GetUserId()); err != nil {
		return nil, err
	}

	return nil, nil
}

const (
	emptyInt    = 0
	emptyString = ""
)

func validateLogin(appId int32, login string, pass string) error {
	if appId == emptyInt {
		return status.Error(codes.InvalidArgument, "wrong appId")
	}
	if login == emptyString {
		return status.Error(codes.InvalidArgument, "wrong login")
	}
	if pass == emptyString {
		return status.Error(codes.InvalidArgument, "wrong pass")
	}
	return nil
}
func validateRegister(login string, pass string) error {
	if login == emptyString {
		return status.Error(codes.InvalidArgument, "wrong login")
	}
	if pass == emptyString {
		return status.Error(codes.InvalidArgument, "wrong pass")
	}
	return nil
}
func validateIsAdmin(userId int64) error {
	if userId == emptyInt {
		return status.Error(codes.InvalidArgument, "wrong userId")
	}
	return nil
}
