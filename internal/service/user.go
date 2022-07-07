package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/kratos-world/user_service/api/user"
	"github.com/kratos-world/user_service/internal/biz"
	"google.golang.org/protobuf/types/known/structpb"
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
	fmt.Println("req", req.GetUserId())

	return &pb.GetUserReply{
		Name: "test",
		Age:  18,
		//WorkJob: "",
	}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}

func (s *UserService) GetLoginInfo(ctx context.Context, req *pb.GetLoginInfoRequest) (*pb.GetLoginInfoReply, error) {
	//s.log.WithContext(ctx).Info("GetLoginInfo params", req.Email, req.Phone)
	fmt.Println("req_login", req.GetUserId())

	uid := ""
	if req.Email == "test@qq.com" {
		uid = "u_eNAjI9cFGZ1EfaoU" // 默认
	}
	if req.Phone == "12345678910" {
		uid = "u_quVaM29CKonWxbes"
	}
	s.log.WithContext(ctx).Info("GetLoginInfo uid", uid)

	items := []interface{}{"aa", "bb", "cc", "dd"}
	l, err := structpb.NewList(items)
	if err != nil {
		return nil, err
	}

	var answerList []*structpb.Value
	answerList = append(answerList, structpb.NewListValue(l))

	res := &pb.GetLoginInfoReply{Uid: uid, UToken: "u_token", WorkJob: "ddddd", AnswerList: answerList}
	return res, nil
}
