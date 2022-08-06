package cache

import (
	"fmt"
	"time"

	"github.com/go-redsync/redsync"
)

var (
	RedSync *redsync.Redsync
)

type RedisDisLock struct {
	*redsync.Mutex
}

func NewRedisDisLock(key string, timeout time.Duration, retries int) *RedisDisLock {
	return &RedisDisLock{
		Mutex: RedSync.NewMutex("dislock:"+key,
			redsync.SetExpiry(timeout),
			redsync.SetRetryDelay(50*time.Millisecond),
			redsync.SetTries(retries)),
	}
}

func (l *RedisDisLock) Lock() error {
	return l.Mutex.Lock()
}

func (l *RedisDisLock) Unlock() bool {
	ubool, err := l.Mutex.Unlock()
	if err != nil {
		fmt.Println("reids dis lock unlock err,", err.Error())
	}
	return ubool
}

// initialize redis connection pool for redis distributed lock to use
func InitRedSync(c *PoolClient) {
	RedSync = redsync.New([]redsync.Pool{
		redsync.Pool(c.Pool),
	})
}
