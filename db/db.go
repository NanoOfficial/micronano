//
//
// @filename: db/db.go
// @author: Krisna Pranav
// @license COPYRIGHT 2023 Krisna Pranav, NanoBlocksDevelopers
//
//

package db

import (
	"sync"

	"github.com/NanoOfficial/micronano/config"
	"github.com/ethereum/go-ethereum/ethdb/leveldb"
	"github.com/syndtr/goleveldb/leveldb"
)

var lock = new(sync.Mutex)
var database *Database

type Database struct {
	config  *config.Database
	leveldb map[string]*leveldb.DB
	sqlite  map[string]*SQLite
}

func NewDatabase(config *config.Database) *Database {
	if database == nil {
		lock.Lock()
		defer lock.Unlock()

		database = &Database{
			config: config,
			leveldb: make(map[string]*leveldb.DB),
			sqlite: make(map[string]*SQLite)
		}
		
		var err error

		if len(config.LevelDB) > 0 {
			for dbName, dbConfig := range config.LevelDB {

			}


		}
	}

	return database
}


func getDatabase() *Database {
	return database
}

func (db *Database) runMigrations() error {
	
}