package config

import (
	"github.com/syklinux/golib/log"
	"github.com/syklinux/golib/mysql"
)

// Conf 配置项
type Conf struct {
	Port        string          `json:"port"`
	Log         log.LogConfig   `json:"log"`
	Mysql       mysql.MySQLConf `json:"mysql"`
	HTTPConf    HTTPConf        `json:"httConf"`
	RedisConf   RedisConf       `json:"redisConf"`
	TokenConfig TokenConfig     `json:"tokenConfig"`
	EmailConfig EmailConfig     `json:"emailConfig"`
}

// HTTPConf HTTPConfS
type HTTPConf struct {
	Addr           string   `json:"addr"`
	ReadTimeout    int      `json:"read_timeout"`
	WriteTimeout   int      `json:"write_timeout"`
	MaxHeaderBytes int      `json:"max_header_bytes"`
	AllowIP        []string `json:"allow_ip"`
	Cors           []string `json:"cors"`
}

// RedisConf RedisConfS
type RedisConf struct {
	Addr     string `json:"Addr"`
	Password string `json:"Password"`
	Db       int    `json:"Db"`
}

// TokenConfig TokenConfig
type TokenConfig struct {
	ExpiresAt int64  `json:"ExpiresAt"`
	SecureKey string `json:"SecureKey"`
	PadKey    string `json:"PadKey"`
}

// EmailConfig EmailConfig
type EmailConfig struct {
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
