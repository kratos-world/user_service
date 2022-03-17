package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/kratos-world/user_service/internal/biz"

	pb "github.com/kratos-world/user_service/api/user"
)

type UserService struct {
	pb.UnimplementedUserServer

	uc  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	return &pb.CreateUserReply{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}

func (s *UserService) GetLoginInfo(ctx context.Context, req *pb.GetLoginInfoRequest) (*pb.GetLoginInfoReply, error) {
	s.log.WithContext(ctx).Info("GetLoginInfo params", req.Email, req.Phone)
	uid := ""
	if req.Email == "test@qq.com" {
		uid = "u_eNAjI9cFGZ1EfaoU" // 默认
	}
	if req.Phone == "12345678910" {
		uid = "u_quVaM29CKonWxbes"
	}
	s.log.WithContext(ctx).Info("GetLoginInfo uid", uid)
	res := &pb.GetLoginInfoReply{Uid: uid}
	return res, nil
}
