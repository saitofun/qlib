package qbuilder

import (
	"reflect"
	"strconv"
	"strings"
	"sync"
)

var (
	ColumnTag        = "db"
	ColumnTagSep     = ","
	onceSetColumnTag = &sync.Once{}
)

func SetColumnTag(tagName, tagSep string, ) {
	if ColumnTag != "" {
		onceSetColumnTag.Do(func() {
			ColumnTag = tagName
			ColumnTagSep = tagSep
		})
	}
}

type StructTags map[string][]string

func ParseTags(tag reflect.StructTag, sep string) (tags StructTags) {
	for tag != "" {
		// skip leading spaces
		i := 0
		for i < len(tag) && tag[i] == ' ' {
			i++
		}
		tag = tag[i:]
		if tag == "" {
			break
		}

		i = 0
		for i < len(tag) && tag[i] > ' ' && tag[i] != ':' && tag[i] != '"' && tag[i] != 0x7f {
			i++
		}
		if i == 0 || i+1 >= len(tag) || tag[i] != ':' || tag[i+1] != '"' {
			break
		}
		key := string(tag[:i])
		tag = tag[i+1:]

		// Scan quoted string to find value.
		i = 1
		for i < len(tag) && tag[i] != '"' {
			if tag[i] == '\\' {
				i++
			}
			i++
		}
		if i >= len(tag) {
			break
		}
		quoted := string(tag[:i+1])
		tag = tag[i+1:]

		value, err := strconv.Unquote(quoted)
		if err != nil {
			break
		}
		if tags == nil {
			tags = make(StructTags)
		}
		if sep != "" {
			tags[key] = append(tags[key], strings.Split(value, ColumnTagSep)...)
		} else {
			tags[key] = append(tags[key], value)
		}
	}
	return
}

func (s StructTags) ColumnName() string {
	return s.GetName(ColumnTag)
}

func (s StructTags) Get(key string) []string {
	return s[key]
}

func (s StructTags) GetName(key string) string {
	if len(s[key]) > 0 {
		return s[key][0]
	}
	return ""
}
