package qconv

import "strconv"

// Float convert any to float64
func Float(v interface{}) float64 {
	if v == nil {
		return 0
	}
	if v, ok := v.(float64); ok {
		return v
	}

	switch v := v.(type) {
	case float32:
		return float64(v)
	case []byte:

	default:
		if v, ok := v.(interface{ Float() float64 }); ok {
			return v.Float()
		}
		if v, err := strconv.ParseFloat(String(v), 64); err == nil {
			return v
		}
	}
	return 0
}
