package array

import (
	"errors"
	"strconv"
)

func Get(structure any, keys ...string) (any, error) {
	if len(keys) == 0 {
		return nil, errors.New("invalid key")
	}

	if structure == nil {
		return nil, errors.New("invalid structure")
	}

	switch structure := structure.(type) {
	case map[string]any:
		if value, ok := structure[keys[0]]; ok {
			if len(keys) == 1 {
				return value, nil
			} else {
				return Get(value, keys[1:]...)
			}
		} else {
			return nil, errors.New("key not found")
		}
	case []any:
		key := keys[0]
		index, err := parseIndex(key)
		if err != nil {
			return nil, err
		}

		if index < 0 || index >= len(structure) {
			return nil, errors.New("index out of range")
		}

		value := structure[index]
		if len(keys) == 1 {
			return value, nil
		} else {
			return Get(value, keys[1:]...)
		}
	default:
		return nil, errors.New("invalid structure")
	}
}

func parseIndex(key string) (int, error) {
	index, err := strconv.Atoi(key)
	if err != nil {
		return -1, err
	}

	return index, nil
}
