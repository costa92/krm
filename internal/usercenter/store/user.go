package store

import (
	"context"
	"github.com/costa92/krm/internal/usercenter/model"
	"gorm.io/gorm"
)

// UsersStore defines methods used to interact with the user store.
type UsersStore interface {
	// Create adds a new user record to the database.
	Create(ctx context.Context, user *model.UserM) error
}

type userStore struct {
	ds *datastore
}

// newUserStore returns a new instance of userStore with the provided datastore.
func newUserStore(ds *datastore) *userStore {
	return &userStore{
		ds: ds,
	}
}

// db is an alias for d.ds.Core(ctx context.Context).
func (d *userStore) db(ctx context.Context) *gorm.DB {
	return d.ds.Core(ctx)
}

// Create adds a new user record to the database.
func (d *userStore) Create(ctx context.Context, user *model.UserM) error {
	return d.db(ctx).Create(user).Error
}
