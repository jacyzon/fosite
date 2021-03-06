package jwt

import (
	"time"
)

type Mapper interface {
	ToMap() map[string]interface{}
	Add(key string, value interface{})
	Get(key string) interface{}
}

func ToString(i interface{}) string {
	if i == nil {
		return ""
	}

	if s, ok := i.(string); ok {
		return s
	}

	return ""
}

func ToTime(i interface{}) time.Time {
	if i == nil {
		return time.Time{}
	}

	if t, ok := i.(int64); ok {
		return time.Unix(t, 0)
	} else if t, ok := i.(float64); ok {
		return time.Unix(int64(t), 0)
	}

	return time.Time{}
}

func Filter(elements map[string]interface{}, keys ...string) map[string]interface{} {
	var keyIdx = make(map[string]bool)
	var result = make(map[string]interface{})

	for _, key := range keys {
		keyIdx[key] = true
	}

	for k, e := range elements {
		if _, ok := keyIdx[k]; !ok {
			result[k] = e
		}
	}

	return result
}

func Copy(elements map[string]interface{}) (result map[string]interface{}) {
	result = make(map[string]interface{}, len(elements))
	for k, v := range elements {
		result[k] = v
	}

	return result
}
