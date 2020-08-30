package dig

import (
	"fmt"
)

// Dig extracts the nested value specified by the keys from v
func Dig(v interface{}, keys ...interface{}) (interface{}, error) {
	n := len(keys)
	for i, key := range keys {
		stringKey, ok := key.(string)
		if ok {
			raw, ok := v.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("%v isn't a map", v)
			}
			v, ok = raw[stringKey]
			if !ok {
				return nil, fmt.Errorf("key %v not found in %v", stringKey, v)
			}
			if i == n-1 {
				return v, nil
			}
			continue
		}
		intKey, ok := key.(int)
		if ok {
			raw, ok := v.([]interface{})
			if !ok {
				return nil, fmt.Errorf("%v isn't a slice", v)
			}
			if intKey < 0 || intKey >= len(raw) {
				return nil, fmt.Errorf("index out of range [%v]: %v", intKey, raw)
			}
			v = raw[intKey]
			if i == n-1 {
				return v, nil
			}
			continue
		}
		return nil, fmt.Errorf("unsupported key type: %v", key)
	}
	return nil, fmt.Errorf("no key is given")
}
