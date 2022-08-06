package v1

import (
	"sync"

	"github.com/dyjwl/gin-web-plugin-demo/internal/store"
)

//go:generate mockgen -self_package=github.com/marmotedu/iam/internal/apiserver/service/v1 -destination mock_service.go -package v1 github.com/marmotedu/iam/internal/apiserver/service/v1 Service,UserSrv,SecretSrv,PolicySrv
var (
	defaultSrv *service
	once       sync.Once
)

// Service defines functions used to return resource interface.
type Service interface {
	Users() UserSrv
}

type service struct {
	userSrv UserSrv
}

// NewService returns Service interface.
func NewService(store store.Factory) Service {
	once.Do(func() {
		defaultSrv = &service{
			userSrv: &userService{store: store},
		}
	})
	return defaultSrv
}

func (s *service) Users() UserSrv {
	return s.userSrv
}
