package user

import (
	"context"
	"github.com/costa92/krm/internal/usercenter/model"
	v1 "github.com/costa92/krm/pkg/api/usercenter/v1"
	"github.com/jinzhu/copier"
)

func (b *userBiz) Create(ctx context.Context, req *v1.CreateUserRequest) (*v1.UserReply, error) {
	var userM model.UserM
	// Copy the fields from the request to the model
	_ = copier.Copy(&userM, req)
	err := b.ds.TX(ctx, func(ctx context.Context) error {
		if err := b.ds.Users().Create(ctx, &userM); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return nil, nil
}
