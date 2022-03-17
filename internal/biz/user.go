package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
}

type UserRepo interface {
	GetUserInfoByEmail(ctx context.Context, email string) error
	GetUserInfoByPhone(ctx context.Context, phone string) error
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) GetUserInfoByEmail(ctx context.Context, email string) error {
	//conn, err := grpc.Dial(
	//	fmt.Sprintf("%s:///%s", r.Scheme(), SerName),
	//	grpc.WithBalancerName("weight"),
	//	grpc.WithInsecure(),
	//)
	return uc.repo.GetUserInfoByEmail(ctx, email)
}

func (uc *UserUsecase) GetUserInfoByPhone(ctx context.Context, phone string) error {
	return uc.repo.GetUserInfoByPhone(ctx, phone)
}
