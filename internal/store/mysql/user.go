package mysql

import (
	"context"
	"errors"

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
	return u.db.Create(&user).Error
}
func (u *users) Update(ctx context.Context, user *model.User) error {
	return u.db.Save(user).Error
}
func (u *users) Delete(ctx context.Context, username string) error {
	return u.db.Where("nickname = ?", username).Delete(&model.User{}).Error
}
func (u *users) DeleteCollection(ctx context.Context, usernames []string) error {
	return u.db.Where("nickname in (?)", usernames).Delete(&model.User{}).Error
}
func (u *users) Get(ctx context.Context, username string) (*model.User, error) {
	user := &model.User{}
	err := u.db.Where("nickname = ? and status = 1", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}
func (u *users) List(ctx context.Context, pageSize, pageNo int, name string) (*model.UserList, error) {

	ret := &model.UserList{}

	u.db = u.db.Model(&model.User{})
	if name != "" {
		u.db = u.db.Where(" nickname like ? and status = 1", "%"+name+"%")
	}
	err := u.db.Count(&ret.Total).Error
	if err != nil {
		return ret, err
	}
	if pageNo > 0 {
		offset := pageSize * (pageNo - 1)
		u.db = u.db.Offset(offset)
	}
	if pageSize > 0 {
		u.db = u.db.Limit(pageSize)
	}
	d := u.db.Order("id desc").
		Find(&ret.Items)
	return ret, d.Error
}
