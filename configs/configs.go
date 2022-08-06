package configs

var Config *config

type config struct {
	Log    Log
	Server Server
	Redis  RedisConf
}

type Log struct {
	Filename  string
	LogLevel  string
	LogMode   string
	MaxSize   int
	MaxAge    int //days
	Compress  bool
	WithColor bool
	ShowLine  bool
}

type Server struct {
	Port int
}

type RedisConf struct {
	Tag         string
	Host        string
	Port        string
	Password    string
	MaxIdle     int
	IdleTimeout int
}
