//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/sqlite"
)

var Sessions = newSessionsTable("", "sessions", "")

type sessionsTable struct {
	sqlite.Table

	// Columns
	ID            sqlite.ColumnInteger
	EnvironmentID sqlite.ColumnInteger
	UserID        sqlite.ColumnString
	Tags          sqlite.ColumnString
	CreatedAt     sqlite.ColumnTimestamp
	UpdatedAt     sqlite.ColumnTimestamp

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type SessionsTable struct {
	sessionsTable

	EXCLUDED sessionsTable
}

// AS creates new SessionsTable with assigned alias
func (a SessionsTable) AS(alias string) *SessionsTable {
	return newSessionsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new SessionsTable with assigned schema name
func (a SessionsTable) FromSchema(schemaName string) *SessionsTable {
	return newSessionsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new SessionsTable with assigned table prefix
func (a SessionsTable) WithPrefix(prefix string) *SessionsTable {
	return newSessionsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new SessionsTable with assigned table suffix
func (a SessionsTable) WithSuffix(suffix string) *SessionsTable {
	return newSessionsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newSessionsTable(schemaName, tableName, alias string) *SessionsTable {
	return &SessionsTable{
		sessionsTable: newSessionsTableImpl(schemaName, tableName, alias),
		EXCLUDED:      newSessionsTableImpl("", "excluded", ""),
	}
}

func newSessionsTableImpl(schemaName, tableName, alias string) sessionsTable {
	var (
		IDColumn            = sqlite.IntegerColumn("id")
		EnvironmentIDColumn = sqlite.IntegerColumn("environment_id")
		UserIDColumn        = sqlite.StringColumn("user_id")
		TagsColumn          = sqlite.StringColumn("tags")
		CreatedAtColumn     = sqlite.TimestampColumn("created_at")
		UpdatedAtColumn     = sqlite.TimestampColumn("updated_at")
		allColumns          = sqlite.ColumnList{IDColumn, EnvironmentIDColumn, UserIDColumn, TagsColumn, CreatedAtColumn, UpdatedAtColumn}
		mutableColumns      = sqlite.ColumnList{EnvironmentIDColumn, UserIDColumn, TagsColumn, CreatedAtColumn, UpdatedAtColumn}
	)

	return sessionsTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:            IDColumn,
		EnvironmentID: EnvironmentIDColumn,
		UserID:        UserIDColumn,
		Tags:          TagsColumn,
		CreatedAt:     CreatedAtColumn,
		UpdatedAt:     UpdatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
