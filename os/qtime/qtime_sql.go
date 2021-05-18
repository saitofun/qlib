package qtime

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
	"time"

	"git.querycap.com/ss/lib/database"
)

var (
	_ sql.Scanner   = (*Time)(nil)
	_ driver.Valuer = (*Time)(nil)
)

func (t *Time) DataType(dialect string) string {
	switch strings.ToLower(dialect) {
	case "sqlite", "sqlite3":
		return "integer"
	default:
		return "bigint"
	}
}

func (t *Time) SQLType(dialect database.DialectName) string {
	return t.DataType(string(dialect))
}

func (t *Time) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		n, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			return fmt.Errorf("sql.Scan() strfmt.Time from: %#v failed: %s",
				v, err.Error())
		}
		t.Time = time.Unix(n, 0)
	case int64:
		if v < 0 {
			t.Time = Zero.Time
		} else {
			t.Time = time.Unix(v, 0)
		}
	case nil:
		t.Time = Zero.Time
	default:
		return fmt.Errorf("cannot sql.Scan() strfmt.Time from: %#v", v)
	}
	return nil
}

func (t Time) Value() (driver.Value, error) {
	v := t.Unix()
	if v < 0 {
		v = 0
	}
	return v, nil
}
