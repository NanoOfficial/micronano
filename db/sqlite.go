//
//
// @filename: db/sqlite.go
// @author: Krisna Pranav
// @license COPYRIGHT 2023 Krisna Pranav, NanoBlocksDevelopers
//
//

package db

import (
	"database/sql"
	"time"
)

type SQLite struct {
	Db         *sql.DB
	migrations []SQLMigration
}

func (s *SQLite) Migrations() []SQLMigration {
	return s.migrations
}

func (s *SQLite) AddMigration(migration ...SQLMigration) {
	s.migrations = append(s.migrations, migration...)
}

func (s *SQLite) RunMigrations() error {
	if len(s.migrations) < 1 {
		return nil
	}

	for _, migration := range s.migrations {
		row := s.Db.QueryRow(FindMigrationEntry, migration.name)
		if row.Err() != nil {
			return row.Err()
		}

		var name string
		_ = row.Scan(&name)
		if name != "" {
			continue
		}

		if _, err := s.Db.Exec(migration.upQuery); err != nil {
			return err
		}

		_, err := s.Db.Exec(NewMigrationEntry, migration.name, time.Now().Unix())
		if err != nil {
			return err
		}

	}

	return nil
}

func (s *SQLite) RevertMigrations() error {
	if len(s.migrations) < 1 {
		return nil
	}

	for i := len(s.migrations) - 1; i >= 0; i-- {
		if _, err := s.Db.Exec(s.migrations[i].downQuery); err != nil {
			return err
		}

		if _, err := s.Db.Exec(DeleteMigrationEntry, s.migrations[i].name); err != nil {
			return err
		}
	}

	return nil
}
