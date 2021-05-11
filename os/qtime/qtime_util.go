package qtime

import (
	"time"

	_ "github.com/uniplaces/carbon"
)

func (t Time) Today() Time {
	n := time.Now()
	return Time{
		Time: time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, time.Local),
	}
}

func (t Time) Month() Time {
	n := time.Now()
	return Time{
		Time: time.Date(n.Year(), n.Month(), 0, 0, 0, 0, 0, time.Local),
	}
}

func (t Time) CurrentWeek() Time    { return Time{} }
func (t Time) CurrentQuarter() Time { return Time{} }
func (t Time) CurrentYear() Time    { return Time{} }
func (t Time) IsLeapYear() bool     { return true }
