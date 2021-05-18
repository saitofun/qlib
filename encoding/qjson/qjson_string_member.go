package qjson

// var Stringer = struct {
// 	Marshal   func(interface{}) ([]byte, error)
// 	Unmarshal func([]byte, interface{}) error
// }{
// 	Marshal: func(v interface{}) ([]byte, error) {
// 		raw, err := Marshal(v)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return Marshal(encoding.BytesToStr(raw))
// 	},
// 	Unmarshal: func(data []byte, v interface{}) error {
// 		var raw string
// 		var err = Unmarshal(data, &raw)
// 		if err != nil {
// 			return err
// 		}
// 		return Unmarshal(encoding.StrToBytes(raw), v)
// 	},
// }

type stringMemberMarker interface {
	stringMemberMarker()
	Stepped() bool
	Step()
}

type StringMemberMarker struct {
	stepped int
}

func (s *StringMemberMarker) stringMemberMarker() {}
func (s *StringMemberMarker) Step()               { s.stepped++ }
func (s *StringMemberMarker) Stepped() bool       { return s.stepped > 0 }
