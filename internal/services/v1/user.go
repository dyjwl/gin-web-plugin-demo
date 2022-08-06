package v1

import (
	"context"
	"regexp"

	"github.com/dyjwl/gin-web-plugin-demo/internal/store"
	"github.com/dyjwl/gin-web-plugin-demo/internal/store/model"
	"github.com/dyjwl/gin-web-plugin-demo/pkg/log"
	"go.uber.org/zap"
)

// UserSrv defines functions used to handle user request.
type UserSrv interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, username string) error
	DeleteCollection(ctx context.Context, usernames []string) error
	Get(ctx context.Context, username string) (*model.User, error)
	List(ctx context.Context) (*model.UserList, error)
	ChangePassword(ctx context.Context, user *model.User) error
}

type userService struct {
	store store.Factory
}

var _ UserSrv = (*userService)(nil)

// List returns user list in the storage. This function has a good performance.
func (u *userService) List(ctx context.Context) (*model.UserList, error) {
	users, err := u.store.Users().List(ctx)
	if err != nil {
		log.Error("list users from storage failed: %s", zap.Error(err))
		return nil, err
	}
	return users, nil
}

func (u *userService) Create(ctx context.Context, user *model.User) error {
	if err := u.store.Users().Create(ctx, user); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'idx_name'", err.Error()); match {
			return err
		}

		return err
	}

	return nil
}

func (u *userService) DeleteCollection(ctx context.Context, usernames []string) error {
	if err := u.store.Users().DeleteCollection(ctx, usernames); err != nil {
		return err
	}

	return nil
}

func (u *userService) Delete(ctx context.Context, username string) error {
	if err := u.store.Users().Delete(ctx, username); err != nil {
		return err
	}

	return nil
}

func (u *userService) Get(ctx context.Context, username string) (*model.User, error) {
	user, err := u.store.Users().Get(ctx, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) Update(ctx context.Context, user *model.User) error {
	if err := u.store.Users().Update(ctx, user); err != nil {
		return err
	}

	return nil
}

func (u *userService) ChangePassword(ctx context.Context, user *model.User) error {
	// Save changed fields.
	if err := u.store.Users().Update(ctx, user); err != nil {
		return err
	}

	return nil
}
