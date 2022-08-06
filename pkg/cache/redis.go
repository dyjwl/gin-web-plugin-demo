package cache

import (
	"fmt"
	"time"

	"github.com/dyjwl/gin-web-plugin-demo/configs"
	"github.com/dyjwl/gin-web-plugin-demo/pkg/log"
	redigo "github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
)

var SyncRedis *PoolClient

type PoolClient struct {
	Pool *redigo.Pool
}

func NewClient(conf configs.RedisConf) *PoolClient {
	url := fmt.Sprintf("redis://%s:%s", conf.Host, conf.Port)
	if conf.Password != "" {
		url = fmt.Sprintf("redis://:%s@%s:%s", conf.Password, conf.Host, conf.Port)
	}
	pool := newPool(url, conf.MaxIdle, conf.IdleTimeout)
	return &PoolClient{Pool: pool}
}

type Trans struct {
	conn redigo.Conn
}

func (t *Trans) Send(cmd string, args ...interface{}) {
	t.conn.Send(cmd, args...)
}

func (t *Trans) Exec() (reply interface{}, err error) {
	defer t.conn.Close()
	return t.conn.Do("EXEC")
}

func (self *PoolClient) Do(cmd string, args ...interface{}) (reply interface{}, err error) {
	conn := self.Pool.Get()
	defer conn.Close()
	return conn.Do(cmd, args...)
}

func (self *PoolClient) GetConn() redigo.Conn {
	return self.Pool.Get()
}

func (self *PoolClient) BeginTrans() *Trans {
	conn := self.Pool.Get()
	conn.Send("MULTI")
	return &Trans{conn}
}

func newPool(url string, maxIdle, idleTimeout int) *redigo.Pool {
	return &redigo.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: time.Duration(idleTimeout) * time.Second,
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.DialURL(url)
			if err != nil {
				log.Info("redis dial failure...", zap.Error(err))
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				log.Info("redis test on borrow failure...", zap.Error(err))
			}
			return err
		},
	}
}
