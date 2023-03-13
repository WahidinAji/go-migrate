package migrate

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type Column struct {
	Type   string
	Field  string
	Length int `default:"255"`
}

func (def *Column) DefaultLength() int {
	def.Length = 255
	return def.Length
}

func Migrate(ctx context.Context, db *sql.DB, table string, column []Column) (res *Response) {
	conn, err := db.Conn(ctx)
	if err != nil {
		res = &Response{status: false, message: err}
		return
	}
	defer func() {
		if errClose := conn.Close(); errClose != nil {
			res = &Response{status: false, message: errClose}
			return
		}
	}()

	return
}

type Migration struct {
	db *sql.DB
}
type Table struct {
	Name   string `json:"name"`
	Fields []Field
}
type Field struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Nullable  bool   `json:"nullable"`
	Exists    bool   `json:"exists"`
	IsDefault bool   `json:"is_default"`
}

func (db Migration) createTable(ctx context.Context, table Table) (err error) {
	var query string
	var fields []string
	for _, field := range table.Fields {
		fields = append(fields, field.Name+" "+field.Type)
	}
	query = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", table.Name, strings.Join(fields, ","))
	_, err = db.db.ExecContext(ctx, query)
	if err != nil {
		err = fmt.Errorf("failed to create table %s: %v", table.Name, err)
		return
	}
	return
}

func migrateVersionOne()
