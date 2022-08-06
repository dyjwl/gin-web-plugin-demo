package mysql

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/dyjwl/gin-web-plugin-demo/configs"
	"github.com/dyjwl/gin-web-plugin-demo/internal/store"
	"github.com/dyjwl/gin-web-plugin-demo/internal/store/dbloger"
	"github.com/dyjwl/gin-web-plugin-demo/internal/store/migrate"
	"github.com/dyjwl/gin-web-plugin-demo/pkg/db"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type datastore struct {
	db *gorm.DB
}

func (ds *datastore) Users() store.UserStore {
	return newUsers(ds)
}
func (ds *datastore) Close() error {
	db, err := ds.db.DB()
	if err != nil {
		return errors.New("get gorm db instance failed")
	}

	return db.Close()
}

var (
	mysqlFactory store.Factory
	once         sync.Once
)

// GetMysqlFactoryOr create mysql factory with the given config.
func GetMysqlFactoryOr(opts *configs.Database) (store.Factory, error) {
	if opts == nil && mysqlFactory == nil {
		return nil, fmt.Errorf("failed to get mysql store fatory")
	}

	var err error
	var dbIns *gorm.DB
	once.Do(func() {
		_default := logger.New(dbloger.NewWriter(
			log.New(os.Stdout, "\r\n", log.LstdFlags)),
			logger.Config{
				SlowThreshold: 200 * time.Millisecond,
				LogLevel:      logger.Warn,
				Colorful:      true,
			})

		options := &db.Options{
			Host:     opts.Host,
			Port:     opts.Port,
			Username: opts.User,
			Schema:   opts.Schema,
			Password: opts.Password,
			Database: opts.Database,
			Logger:   _default,
			Dialect:  opts.Dialect,
		}
		switch opts.LogMode {
		case "silent", "Silent":
			options.Logger = _default.LogMode(logger.Silent)
		case "error", "Error":
			options.Logger = _default.LogMode(logger.Error)
		case "warn", "Warn":
			options.Logger = _default.LogMode(logger.Warn)
		case "info", "Info":
			options.Logger = _default.LogMode(logger.Info)
		default:
			options.Logger = _default.LogMode(logger.Info)
		}
		dbIns, err = db.New(options)

		// uncomment the following line if you need auto migration the given models
		// not suggested in production environment.
		migrate.MigrateDatabase(dbIns)

		mysqlFactory = &datastore{dbIns}
	})

	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store fatory, mysqlFactory: %+v, error: %w", mysqlFactory, err)
	}

	return mysqlFactory, nil
}
