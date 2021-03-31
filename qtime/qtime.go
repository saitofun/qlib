package qtime

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	UTC      = time.UTC
	CST      = time.FixedZone("CST", 60*60*8)
	Zero     = Time{time.Time{}}
	UnixZero = Time{time.Unix(0, 0)}
)

const (
	LogFormat     = "2006-01-02 15:04:05"
	LogFormatNano = "2006-01-02 15:04:05.000"
	Format        = time.RFC3339
	FormatNano    = time.RFC3339Nano
	GANsFormat    = "2006/01/02 15:04:05.000"
)

var Formats = []string{
	GANsFormat,
	LogFormat,
	LogFormatNano,
	Format,
	FormatNano,
}

type Time struct {
	time.Time
}

var (
	_ sql.Scanner   = (*Time)(nil)
	_ driver.Valuer = (*Time)(nil)
)

func (t *Time) DataType(engine string) string {
	switch strings.ToLower(engine) {
	case "sqlite", "sqlite3":
		return "integer"
	default:
		return "bigint"
	}
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

func (t Time) String() string {
	if t.IsZero() {
		return ""
	}
	return t.Time.In(CST).Format(Format)
}

func (t Time) LogFormat() string {
	if t.IsZero() {
		return ""
	}
	return t.Time.In(CST).Format(LogFormatNano)
}

func (t Time) MarshalJSON() ([]byte, error) {
	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalJSON: year outside of range")
	}

	b := make([]byte, 0, len(Format)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, Format)
	b = append(b, '"')
	return b, nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error
	for _, v := range Formats {
		t.Time, err = time.ParseInLocation(`"`+v+`"`, string(data), CST)
		if err == nil {
			break
		}
	}
	return err
}

func Now() Time {
	return Time{time.Now()}
}

func NowSecond() int64 {
	return Now().Unix()
}

func NowMillionSecond() int64 {
	return Now().UnixNano() / 1e6
}

func NowMicroSecond() int64 {
	return Now().UnixNano() / 1e3
}

func NowNanoSecond() int64 {
	return Now().UnixNano()
}
