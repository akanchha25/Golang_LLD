package data

type KeyValuePair struct {
	Key string
	Val string
}

// NewKeyValuePair is a constructor function that creates a new KeyValuePair
func NewKeyValuePair(key string, val string) *KeyValuePair {
    return &KeyValuePair{Key: key, Val: val}
}

// GetKey returns the key of the KeyValuePair
func(kp *KeyValuePair) GetKey() string {
	return kp.Key
}

// GetVal returns the value of the KeyValuePair
func(kp *KeyValuePair) GetVal() string {
	return kp.Val
}