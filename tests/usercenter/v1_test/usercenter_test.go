package v1_test

import (
	"context"
	"github.com/costa92/krm/internal/usercenter/service"
	v1 "github.com/costa92/krm/pkg/api/usercenter/v1"
	"github.com/stretchr/testify/assert"
	"testing"
)

var usercenter *service.UserCenterService

func initUserCenter() {
	//
	//biz := biz2.NewBiz()
	//// init usercenter
	//usercenter := service.NewUserCenterService()
}

func TestCreatUser(t *testing.T) {
	ctx := context.Background()
	req := v1.CreateUserRequest{
		Username: "costalong",
		Nickname: "costa",
		Password: "abcd",
		Email:    "costa9293@gmail.com",
		Phone:    "1557****411",
	}
	user, err := usercenter.CreateUser(ctx, &req)
	assert.NoError(t, err)
	assert.Equal(t, user.UserID, "costalong")
}
