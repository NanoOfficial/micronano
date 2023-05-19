//
//
// @filename: config/config.go
// @author: Krisna Pranav
// @license COPYRIGHT 2023 Krisna Pranav, NanoBlocksDevelopers
//
//

package config

import (
	"github.com/NanoOfficial/micronano/common"
	"github.com/NanoOfficial/micronano/common/types"
	log "github.com/NanoOfficial/micronano/logger"
)

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
	HeartbeatInterval uint                 `yaml:"heartbeatInterval"`
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
	LevelDB map[string]DatabaseLevelDB `yaml:"leveldb"`
	SQLite  map[string]DatabaseSQLite  `yaml:"sqlite"`
	CouchDB DatabaseCouchDB            `yaml:"couch_db"`
}

type DatabaseLevelDB struct {
	DbPath          string `yaml:"path"`
	WriteBufferSize int    `yaml:"writeBufferSize"`
}

type DatabaseSQLite struct {
	DbPath             string `yaml:"path"`
	Options            string `yaml:"options"`
	MaxOpenConnections int    `yaml:"maxOpenConnections"`
}

type DatabaseCouchDB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Prefix   string `yaml:"prefix"`
}

func NewConfig() Config {
	return Config{}
}

func (c *Config) GetConfig() *Config {
	return c
}

func (c *Config) LoadFromFile(filePath string) error {
	log.New("load config from file: "+filePath, log.TypeDebug)

	file, err := common.OpenFile(filePath, common.YML)
	if err != nil {
		return err
	}

	defer file.Close()

	errUnmarshal := file.Parse(&c)
	if errUnmarshal != nil {
		return errUnmarshal
	}

	return nil
}
