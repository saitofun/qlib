// +build jsoniter

package qjson

import (
	"github.com/json-iterator/go"
)

var (
	json          = jsoniter.ConfigCompatibleWithStandardLibrary
	Marshal       = jsoniter.Marshal
	MarshalIndent = jsoniter.MarshalIndent
	NewDecoder    = jsoniter.NewDecoder
	NewEncoder    = jsoniter.NewEncoder
	Unmarshal     = jsoniter.Unmarshal
)
