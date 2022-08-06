package pg

import (
	"context"

	"github.com/dyjwl/gin-web-plugin-demo/internal/store/model"
	"gorm.io/gorm"
)

type users struct {
	db *gorm.DB
}

func newUsers(ds *datastore) *users {
	return &users{ds.db}
}

func (u *users) Create(ctx context.Context, user *model.User) error {
	return nil
}
func (u *users) Update(ctx context.Context, user *model.User) error {
	return nil
}
func (u *users) Delete(ctx context.Context, username string) error {
	return nil
}
func (u *users) DeleteCollection(ctx context.Context, usernames []string) error {
	return nil
}
func (u *users) Get(ctx context.Context, username string) (*model.User, error) {
	return nil, nil
}
func (u *users) List(ctx context.Context, pageSize, pageNo int, name string) (*model.UserList, error) {
	return nil, nil
}
