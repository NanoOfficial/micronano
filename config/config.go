//
//
// @filename: common/config.go
// @author: Krisna Pranav
// @license COPYRIGHT 2023 Krisna Pranav, NanoBlocksDevelopers
//
//

package config

import "github.com/NanoOfficial/micronano/common/types"

type ApplicationConfig interface {
	LoadFromFile(filePath string) error
	GetConfig() *Config
}

type Config struct {
	ApplicationConfig `yaml:"-"`
	Debug             bool        `yaml:"debug"`
	Transporter       Transporter `yaml:"transporter"`
	Database          Database    `yaml:"database"`
}

type Transporter struct {
	Timeout           uint                 `yaml:"timeout"`
	HeartbeatInterval uint                 `yaml:"heartbeatinterval"`
	DeliveryMethod    types.DeliveryMethod `yaml:"deliveryMethod"`
	Redis             *RedisTransporter    `yaml:"redis"`
}

type RedisTransporter struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Db       uint   `yaml:"db"`
}

type Database struct {
	LevelDB map[string]DatabaseLevelDB
	SQlite  map[string]DatabaseSQLite
	CouchDB DatabaseCouchDB
}

type DatabaseLevelDB struct {
}

type DatabaseSQLite struct {

}

type DatabaseCouchDB {
	
}

func NewConfig() Config {
	return Config{}
}
