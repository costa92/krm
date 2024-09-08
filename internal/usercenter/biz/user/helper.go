package user

import (
	"github.com/costa92/krm/internal/usercenter/model"
	v1 "github.com/costa92/krm/pkg/api/usercenter/v1"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ModelToReply converts a model.UserM to a v1.UserReply. It copies the data from
// userM to user and sets the CreatedAt and UpdatedAt fields to their respective timestamps.
func ModelToReply(userM *model.UserM) *v1.UserReply {
	var user v1.UserReply
	_ = copier.Copy(&user, userM)
	user.CreatedAt = timestamppb.New(userM.CreatedAt)
	user.UpdatedAt = timestamppb.New(userM.UpdatedAt)
	return &user
}
