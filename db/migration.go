//
//
// @filename: db/migration.go
// @author: Krisna Pranav
// @license COPYRIGHT 2023 Krisna Pranav, NanoBlocksDevelopers
//
//

package db

const CreateMigrationTable string = `
		CREATE TABLE IF NOT EXISTS migration (
			name TEXT NOT NULL PRIMARY KEY,
			created_at INTEGER NOT NULL
		);
`

type SQLMigration struct {
	name      string
	upQuery   string
	downQuery string
}

func NewSQLMigration(name string, upQuery string, downQuery string) SQLMigration {
	return SQLMigration{
		name:      name,
		upQuery:   upQuery,
		downQuery: downQuery,
	}
}
