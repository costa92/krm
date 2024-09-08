package store

import (
	"context"
	"github.com/costa92/krm/internal/usercenter/model"
	"gorm.io/gorm"
)

// SecretStore defines methods used to interact with the secret store.
type SecretStore interface {
	Create(ctx context.Context, secret *model.SecretM) error
}

type secretStore struct {
	ds *datastore
}

func newSecretStore(ds *datastore) *secretStore {
	return &secretStore{ds}
}

// db is an alias for accessing the Core method of the datastore using the provided context.
func (d *secretStore) db(ctx context.Context) *gorm.DB {
	return d.ds.Core(ctx)
}

// Create adds a new secret record in the datastore.
func (d *secretStore) Create(ctx context.Context, secret *model.SecretM) error {
	return d.db(ctx).Create(&secret).Error
}
