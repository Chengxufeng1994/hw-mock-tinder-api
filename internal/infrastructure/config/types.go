package config

type SecureString string

// String returns an elided version.  It is safe to call for logging.
func (SecureString) String() string {
	return "[SECRET]"
}

// SecureValue returns the actual value of s as a string.
func (s SecureString) SecureValue() string {
	return string(s)
}

func (s SecureString) MarshalText() ([]byte, error) {
	if string(s) == "" {
		return []byte(""), nil
	}
	return []byte("[SECRET]"), nil
}
