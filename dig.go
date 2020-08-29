package dig

import (
	"fmt"
)

// Dig extracts the nested value specified by the keys from m
func Dig(m interface{}, keys ...interface{}) (interface{}, error) {
	n := len(keys)
	for i, key := range keys {
		stringKey, ok := key.(string)
		if ok {
			raw, ok := m.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("%v isn't a map", m)
			}
			m, ok = raw[stringKey]
			if !ok {
				return nil, fmt.Errorf("key %v not found in %v", stringKey, m)
			}
			if i == n-1 {
				return m, nil
			}
			continue
		}
		intKey, ok := key.(int)
		if ok {
			raw, ok := m.([]interface{})
			if !ok {
				return nil, fmt.Errorf("%v isn't a slice", m)
			}
			m = raw[intKey]
			if i == n-1 {
				return m, nil
			}
			continue
		}
		return nil, fmt.Errorf("unsupported key type: %v", key)
	}
	return nil, fmt.Errorf("something is wrong")
}
