package qtime

import (
	"errors"
	"time"
)

var (
	UTC      = time.UTC
	CST      = time.FixedZone("CST", 60*60*8)
	Zero     = Time{time.Time{}}
	UnixZero = Time{time.Unix(0, 0)}
)

const (
	LogFormatNano = "2006-01-02 15:04:05.000"
	FormatNano    = time.RFC3339Nano
	LogFormat     = "2006-01-02 15:04:05"
	Format        = time.RFC3339
)

var Formats = []string{
	"2006/01/02 15:04:05.000",
	LogFormatNano,
	FormatNano,
	LogFormat,
	Format,
}

type Time struct {
	time.Time
}

func NewTime() *Time { return &Time{time.Now()} }

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
	t.Round(0)
	if t.IsZero() {
		return []byte(`""`), nil
	}
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
	if string(data) == "null" || string(data) == `""` {
		*t = Zero
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

func (t Time) MarshalText() ([]byte, error) {
	t.Round(0)
	return []byte(t.String()), nil
}

func (t *Time) UnmarshalText(data []byte) (err error) {
	str := string(data)
	if len(str) == 0 || str == "0" {
		return nil
	}
	*t, err = Parse(str)
	return
}

func (t *Time) IsZero() bool {
	if t.Time.IsZero() {
		return true
	}
	unix := t.Unix()
	return unix == 0 || unix == UnixZero.Unix()
}

func Now() Time {
	return Time{time.Now()}
}

func NowLocal() Time {
	return Time{time.Now().Local()}
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

func Parse(s string) (t Time, e error) {
	t.Time, e = time.Parse(Format, s)
	return
}

func ParseWithLayout(val, layout string) (t Time, e error) {
	t.Time, e = time.ParseInLocation(layout, val, CST)
	if e != nil {
		t = UnixZero
	}
	return
}
