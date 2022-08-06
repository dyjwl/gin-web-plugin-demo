package store

import (
	"context"

	"github.com/dyjwl/gin-web-plugin-demo/internal/store/model"
)

// UserStore defines the user storage interface.
type UserStore interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, username string) error
	DeleteCollection(ctx context.Context, usernames []string) error
	Get(ctx context.Context, username string) (*model.User, error)
	List(ctx context.Context) (*model.UserList, error)
}
