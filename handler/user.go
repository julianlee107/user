package handler

import (
	"context"

	"github.com/julianlee107/user/domain/model"

	"github.com/julianlee107/user/domain/service"
	"github.com/julianlee107/user/proto/user/pb/user"
)

type User struct {
	UserDataService service.IUserDataService
}

// Register 注册
func (u *User) Register(ctx context.Context, in *user.UserRegisterRequest, out *user.UserRegisterResponse) error {
	userRegister := &model.User{
		UserName:     in.UserName,
		FirstName:    in.FirstName,
		HashPassword: in.Pwd,
	}
	_, err := u.UserDataService.AddUser(userRegister)
	if err != nil {
		return err
	}
	out.Message = "添加成功"
	return nil
}

func (u *User) Login(ctx context.Context, in *user.UserLoginRequest, out *user.UserLoginResponse) error {
	isOk, err := u.UserDataService.CheckPwd(in.UserName, in.Pwd)
	if err != nil {
		return err
	}
	out.IsSuccess = isOk
	return nil
}

func (u *User) GetUserInfo(ctx context.Context, in *user.UserInfoRequest, out *user.UserInfoResponse) error {
	userInfo, err := u.UserDataService.FindUserByName(in.UserName)
	if err != nil {
		return err
	}
	out = UserForResponse(userInfo)
	return nil
}

func UserForResponse(userInfo *model.User) *user.UserInfoResponse {
	response := &user.UserInfoResponse{
		UserId:    userInfo.ID,
		UserName:  userInfo.UserName,
		FirstName: userInfo.FirstName,
	}
	return response
}
