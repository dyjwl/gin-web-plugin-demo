package configs

var Config *config

type config struct {
	Log    Log
	Server Server
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
