package env

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var (
	ErrUndefinedVariable = errors.New("environment variable is undefined")
	ErrEmptyVariable     = errors.New("environment variable is empty")
)

// WithString retrieves the value of an environment variable identified by the key.
//
// ErrUndefinedVariable is returned if the environment variable lookup fails.
func WithString(key string) (string, error) {
	val, err := lookup(key)
	if err != nil {
		return "", err
	}
	return val, nil
}

// WithBool retrieves the value of an environment variable identified by the key
// and tries to parse it as a boolean. An error is returned if the boolean parsing fails.
//
// ErrUndefinedVariable is returned if the environment variable lookup fails.
func WithBool(key string) (bool, error) {
	val, err := lookup(key)
	if err != nil {
		return false, err
	}

	parsed, err := strconv.ParseBool(val)
	if err != nil {
		return false, fmt.Errorf("could not parse boolean: %w", err)
	}

	return parsed, nil
}

// WithInt retrieves the value of an environment variable identified by the key
// and tries to parse it as an integer. An error is returned if the integer parsing fails.
//
// ErrUndefinedVariable is returned if the environment variable lookup fails.
func WithInt(key string) (int, error) {
	val, err := lookup(key)
	if err != nil {
		return -1, err
	}

	parsed, err := parseInt(val, 0, 0)
	if err != nil {
		return -1, err
	}

	return int(parsed), nil
}

// WithInt8 retrieves the value of an environment variable identified by the key
// and tries to parse it as an integer. An error is returned if the integer parsing fails.
//
// ErrUndefinedVariable is returned if the environment variable lookup fails.
func WithInt8(key string) (int8, error) {
	val, err := lookup(key)
	if err != nil {
		return -1, err
	}

	parsed, err := parseInt(val, 0, 8)
	if err != nil {
		return -1, err
	}

	return int8(parsed), nil
}

// WithInt16 retrieves the value of an environment variable identified by the key
// and tries to parse it as an integer. An error is returned if the integer parsing fails.
//
// ErrUndefinedVariable is returned if the environment variable lookup fails.
func WithInt16(key string) (int16, error) {
	val, err := lookup(key)
	if err != nil {
		return -1, err
	}

	parsed, err := parseInt(val, 0, 16)
	if err != nil {
		return -1, err
	}

	return int16(parsed), nil
}

// WithInt32 retrieves the value of an environment variable identified by the key
// and tries to parse it as an integer. An error is returned if the integer parsing fails.
//
// ErrUndefinedVariable is returned if the environment variable lookup fails.
func WithInt32(key string) (int32, error) {
	val, err := lookup(key)
	if err != nil {
		return -1, err
	}

	parsed, err := parseInt(val, 0, 32)
	if err != nil {
		return -1, err
	}

	return int32(parsed), nil
}

// WithInt64 retrieves the value of an environment variable identified by the key
// and tries to parse it as an integer. An error is returned if the integer parsing fails.
//
// ErrUndefinedVariable is returned if the environment variable lookup fails.
func WithInt64(key string) (int64, error) {
	val, err := lookup(key)
	if err != nil {
		return -1, err
	}

	parsed, err := parseInt(val, 0, 64)
	if err != nil {
		return -1, err
	}

	return int64(parsed), nil
}

// WithFloat32 retrieves the value of an environment variable identified by the key
// and tries to parse it as a float32. An error is returned if the float32 parsing fails.
//
// ErrUndefinedVariable is returned if the environment variable lookup fails.
func WithFloat32(key string) (float32, error) {
	val, err := lookup(key)
	if err != nil {
		return -1, err
	}

	parsed, err := parseFloat(val, 32)
	if err != nil {
		return -1, err
	}

	return float32(parsed), nil
}

// WithFloat64 retrieves the value of an environment variable identified by the key
// and tries to parse it as a float64. An error is returned if the float64 parsing fails.
//
// ErrUndefinedVariable is returned if the environment variable lookup fails.
func WithFloat64(key string) (float64, error) {
	val, err := lookup(key)
	if err != nil {
		return -1, err
	}

	parsed, err := parseFloat(val, 64)
	if err != nil {
		return -1, err
	}

	return parsed, nil
}

// lookup is a helper function to check the content of an environment variable.
func lookup(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return "", ErrUndefinedVariable
	}

	if val == "" {
		return "", ErrEmptyVariable
	}

	return val, nil
}
