package env

// WithDefaultString retrieves the value of an environment variable identified by the key.
// If the environment variable is not set, it returns the fallback default string provided.
//
// This method is a convenience wrapper around WithString to silent the error and return a fallback value.
func WithDefaultString(key string, fallback string) string {
	val, err := WithString(key)
	if err != nil {
		return fallback
	}
	return val
}

// WithDefaultBool retrieves the value of an environment variable identified by the key.
// If the environment variable is not set, it returns the fallback default bool provided.
//
// This method is a convenience wrapper around WithBool to silent the error and return a fallback value.
func WithDefaultBool(key string, fallback bool) bool {
	val, err := WithBool(key)
	if err != nil {
		return fallback
	}
	return val
}

// WithDefaultInt retrieves the value of an environment variable identified by the key.
// If the environment variable is not set, it returns the fallback default int provided.
//
// This method is a convenience wrapper around WithInt to silent the error and return a fallback value.
func WithDefaultInt(key string, fallback int) int {
	val, err := WithInt(key)
	if err != nil {
		return fallback
	}
	return val
}

// WithDefaultInt8 retrieves the value of an environment variable identified by the key.
// If the environment variable is not set, it returns the fallback default int8 provided.
//
// This method is a convenience wrapper around WithInt8 to silent the error and return a fallback value.
func WithDefaultInt8(key string, fallback int8) int8 {
	val, err := WithInt8(key)
	if err != nil {
		return fallback
	}
	return val
}

// WithDefaultInt16 retrieves the value of an environment variable identified by the key.
// If the environment variable is not set, it returns the fallback default int16 provided.
//
// This method is a convenience wrapper around WithInt16 to silent the error and return a fallback value.
func WithDefaultInt16(key string, fallback int16) int16 {
	val, err := WithInt16(key)
	if err != nil {
		return fallback
	}
	return val
}

// WithDefaultInt32 retrieves the value of an environment variable identified by the key.
// If the environment variable is not set, it returns the fallback default int32 provided.
//
// This method is a convenience wrapper around WithInt32 to silent the error and return a fallback value.
func WithDefaultInt32(key string, fallback int32) int32 {
	val, err := WithInt32(key)
	if err != nil {
		return fallback
	}
	return val
}

// WithDefaultInt64 retrieves the value of an environment variable identified by the key.
// If the environment variable is not set, it returns the fallback default int64 provided.
//
// This method is a convenience wrapper around WithInt8 to silent the error and return a fallback value.
func WithDefaultInt64(key string, fallback int64) int64 {
	val, err := WithInt64(key)
	if err != nil {
		return fallback
	}
	return val
}

// WithDefaultFloat32 retrieves the value of an environment variable identified by the key.
// If the environment variable is not set, it returns the fallback default int provided.
//
// This method is a convenience wrapper around WithFloat32 to silent the error and return a fallback value.
func WithDefaultFloat32(key string, fallback float32) float32 {
	val, err := WithFloat32(key)
	if err != nil {
		return fallback
	}
	return val
}

// WithDefaultFloat64 retrieves the value of an environment variable identified by the key.
// If the environment variable is not set, it returns the fallback default int provided.
//
// This method is a convenience wrapper around WithFloat64 to silent the error and return a fallback value.
func WithDefaultFloat64(key string, fallback float64) float64 {
	val, err := WithFloat64(key)
	if err != nil {
		return fallback
	}
	return val
}
