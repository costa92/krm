package service

import (
	"context"
	biz2 "github.com/costa92/krm/internal/usercenter/biz"
	"github.com/costa92/krm/internal/usercenter/biz/user"
	v1 "github.com/costa92/krm/pkg/api/usercenter/v1"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreateUser(t *testing.T) {
	//初始化一个mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	biz := biz2.NewMockIBiz(ctrl)
	userCenterServer := NewUserCenterService(biz)

	ctx := context.Background()
	req := v1.CreateUserRequest{
		Username: "costalong",
		Nickname: "costa",
		Password: "abcd",
		Email:    "costa9293@gmail.com",
		Phone:    "1557****411",
	}
	// NewMockUserBiz
	userBiz := user.NewMockUserBiz(ctrl)
	// 模拟 Create 方法
	userBiz.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&v1.UserReply{
		UserID: "costalong",
	}, nil)
	biz.EXPECT().Users().Return(userBiz).AnyTimes()
	res, err := userCenterServer.CreateUser(ctx, &req)
	assert.ErrorIs(t, err, nil)
	assert.Equal(t, res.UserID, "costalong")
}
