package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/kratos-world/user_service/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) GetUserInfoByEmail(ctx context.Context, email string) error {
	return nil
}

func (r *userRepo) GetUserInfoByPhone(ctx context.Context, phone string) error {
	return nil
}
