package data

import "errors"


type JSON struct{
	KeyToValues map[string]*JSON
}

// NewJSON creates a new JSON object
func NewJSON(keyToValues map[string]*JSON) *JSON {
	return &JSON{KeyToValues: keyToValues}
}

//Get retrieves the JSON value associated with the given key
func (j *JSON) Get(key string) (*JSON, error) {
	if val, exists := j.KeyToValues[key]; exists {
		return val, nil
	}
	return nil, errors.New("key not found")
}

// GetAllKeys return a ;ist of all keys in the JSON object
func (j *JSON) GetAllKeys() []string {
	keys := make([]string, 0, len(j.KeyToValues))
	for key := range j.KeyToValues {
		keys = append(keys, key)
	}

	return keys
}

// IsLeaf checks if the JSON object is a leaf (contains only one key-value pair with a nil value).
func (j *JSON) IsLeaf() bool {
	if len(j.KeyToValues) == 1 {
		for _, val := range j.KeyToValues {
			return val == nil
		}
	}
	return false
}