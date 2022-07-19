package config

type Config struct {
	Server   *Server   `json:"server"`
	Telegram *Telegram `json:"telegram"`
	Template string   `json:"template"`
	LogLevel string   `json:"log_level"`
}

type Server struct {
	Address  string `json:"address"`
}

type Telegram struct {
	Token               string           `json:"token"`
	DefaultChat         int64            `json:"defaultChat"`
	MapChats            map[string]int64 `json:"mapChats"`
	DisableNotification bool             `json:"disableNotification"`
}
