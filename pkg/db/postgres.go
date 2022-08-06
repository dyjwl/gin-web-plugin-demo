package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Options defines optsions for database.
type Options struct {
	Host                  string
	Username              string
	Password              string
	Database              string
	Schema                string
	Dialect               string
	Port                  int
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
	Logger                logger.Interface
}

// GormPgSqlByConfig 初始化 Postgresql 数据库 通过参数
func New(opts *Options) (*gorm.DB, error) {
	switch opts.Dialect {
	case "postgres":
		dsn := "host=" + opts.Host + " user=" + opts.Username + " password=" +
			opts.Password + " dbname=" + opts.Database + fmt.Sprintf(" port=%d", opts.Port) +
			" search_path=" + opts.Schema
		pgsqlConfig := postgres.Config{
			DSN:                  dsn, // DSN data source name
			PreferSimpleProtocol: false,
		}
		db, err := gorm.Open(postgres.New(pgsqlConfig), &gorm.Config{
			Logger: opts.Logger,
		})
		if err != nil {
			panic(err)
		}
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)
		sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)
		sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)
		return db, nil
	case "mysql":
		dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
			opts.Username,
			opts.Password,
			opts.Host,
			opts.Database,
			true,
			"Local")

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: opts.Logger,
		})
		if err != nil {
			return nil, err
		}

		sqlDB, err := db.DB()
		if err != nil {
			return nil, err
		}

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)

		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)

		return db, nil
	case "sqlite":
		// github.com/mattn/go-sqlite3
		db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
			Logger: opts.Logger,
		})
		return db, err
	default:
		fmt.Println("valid database config dialect in this platform is:")
		fmt.Println("[mysql|postgres|sqlite]")
		fmt.Println(fmt.Sprintf("your provide a invalid dialect %s",
			opts.Dialect))
	}
	return nil, nil
}
