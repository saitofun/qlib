package mysql

import (
	"database/sql"

	"github.com/saitofun/qlib/database"
)

type Driver struct {
}

func (d *Driver) Open(dsn string) (*sql.DB, error) { return nil, nil }

func (d *Driver) Tables() ([]string, error) { return nil, nil }

func (d *Driver) TableColumns() (map[string][]*database.Field, error) { return nil, nil }

func (d *Driver) TableIndexes() (map[string][]*database.Index, error) { return nil, nil }

func (d *Driver) Options() (map[string]string, error) { return nil, nil }
