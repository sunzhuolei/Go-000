package service

import (
	"context"
	"errors"
	"github.com/google/wire"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	api_user_v1 "task/Go-000/Week04/api/user/v1"
	"task/Go-000/Week04/internal/dao"
)

type UserService struct {
	dao dao.Dao
}
var Provider = wire.NewSet(NewUserService, dao.Provider)

func NewUserService(d dao.Dao) *UserService {
	return &UserService{dao: d}
}

/**
获取用户信息
 */
func(this *UserService)GetUser(ctx context.Context, request *api_user_v1.UserRequest) (*api_user_v1.UserResponse, error){
	user, err := this.dao.GetUser(ctx, int(request.Uid))
	if err != nil {
		if errors.Is(err, dao.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "Object Not Found")
		}
		return nil, status.Errorf(codes.Internal, "Error:%v", err)
	}
	return &api_user_v1.UserResponse{Name: user.Name, Age: int64(user.Age)}, nil
}
