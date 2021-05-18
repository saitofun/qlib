// +build !jsoniter

package qjson

import "encoding/json"

var (
	Marshal       = json.Marshal
	MarshalIndent = json.MarshalIndent
	Unmarshal     = json.Unmarshal
	NewDecoder    = json.NewDecoder
	NewEncoder    = json.NewEncoder
)

// func Marshal(v interface{}) ([]byte, error) {
// 	dat, err := json.Marshal(v)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if _, ok := v.(stringMemberMarker); ok {
// 		return append(append([]byte{'"'}, dat...), '"'), nil
// 	}
// 	return dat, nil
// }
//
// func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
// 	dat, err := json.MarshalIndent(v, prefix, indent)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if _, ok := v.(stringMemberMarker); ok {
// 		return append(append([]byte{'"'}, dat...), '"'), nil
// 	}
// 	return dat, nil
// }
//
// func Unmarshal(data []byte, v interface{}) error {
// 	if _, ok := v.(stringMemberMarker); ok {
// 		if data[0] == '"' && data[len(data)-1] == '"' {
// 			data = data[1 : len(data)-1]
// 		}
// 	}
// 	return json.Unmarshal(data, v)
// }
//
