package env

import (
	"github.com/stretchr/testify/assert"
	"github.com/thomasgouveia/goutils/env/testutils"
	"testing"
)

func TestWithString(t *testing.T) {
	cases := []testutils.TestCase[string]{
		{
			Name:     "correctly retrieves the value",
			Given:    "FOO",
			Expected: "bar",
			Env: map[string]string{
				"FOO": "bar",
			},
		},
		{
			Name:       "returns ErrUndefinedVariable if the env var is not defined",
			Given:      "FOO",
			ShouldFail: true,
			Error:      ErrUndefinedVariable,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.PopulateEnv(t)

			received, err := WithString(tc.Given)

			if tc.ShouldFail {
				assert.Error(t, err)
				if tc.Error != nil {
					assert.ErrorIs(t, err, tc.Error)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.Expected, received)
			}
		})
	}
}

func TestWithBool(t *testing.T) {
	cases := []testutils.TestCase[bool]{
		{
			Name:     "correctly retrieves the value",
			Given:    "FOO",
			Expected: true,
			Env: map[string]string{
				"FOO": "true",
			},
		},
		{
			Name:       "returns ErrUndefinedVariable if the env var is not defined",
			Given:      "FOO",
			ShouldFail: true,
			Error:      ErrUndefinedVariable,
		},
		{
			Name:       "returns an error if the value is not a boolean",
			Given:      "FOO",
			ShouldFail: true,
			Env: map[string]string{
				"FOO": "notaboolean",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.PopulateEnv(t)

			received, err := WithBool(tc.Given)

			if tc.ShouldFail {
				assert.Error(t, err)
				if tc.Error != nil {
					assert.ErrorIs(t, err, tc.Error)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.Expected, received)
			}
		})
	}
}

func TestWithInt(t *testing.T) {
	cases := []testutils.TestCase[int]{
		{
			Name:     "correctly retrieves the value",
			Given:    "FOO",
			Expected: 10,
			Env: map[string]string{
				"FOO": "10",
			},
		},
		{
			Name:       "returns ErrUndefinedVariable if the env var is not defined",
			Given:      "FOO",
			ShouldFail: true,
			Error:      ErrUndefinedVariable,
		},
		{
			Name:       "returns an error if the value is not an integer",
			Given:      "FOO",
			ShouldFail: true,
			Env: map[string]string{
				"FOO": "notaninteger",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.PopulateEnv(t)

			received, err := WithInt(tc.Given)

			if tc.ShouldFail {
				assert.Error(t, err)
				if tc.Error != nil {
					assert.ErrorIs(t, err, tc.Error)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.Expected, received)
			}
		})
	}
}

func TestWithInt8(t *testing.T) {
	cases := []testutils.TestCase[int8]{
		{
			Name:     "correctly retrieves the value",
			Given:    "FOO",
			Expected: 10,
			Env: map[string]string{
				"FOO": "10",
			},
		},
		{
			Name:       "returns ErrUndefinedVariable if the env var is not defined",
			Given:      "FOO",
			ShouldFail: true,
			Error:      ErrUndefinedVariable,
		},
		{
			Name:       "returns an error if the value is not an integer",
			Given:      "FOO",
			ShouldFail: true,
			Env: map[string]string{
				"FOO": "notaninteger",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.PopulateEnv(t)

			received, err := WithInt8(tc.Given)

			if tc.ShouldFail {
				assert.Error(t, err)
				if tc.Error != nil {
					assert.ErrorIs(t, err, tc.Error)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.Expected, received)
			}
		})
	}
}

func TestWithInt16(t *testing.T) {
	cases := []testutils.TestCase[int16]{
		{
			Name:     "correctly retrieves the value",
			Given:    "FOO",
			Expected: 32767, // maximum value for int16
			Env: map[string]string{
				"FOO": "32767",
			},
		},
		{
			Name:       "returns ErrUndefinedVariable if the env var is not defined",
			Given:      "FOO",
			ShouldFail: true,
			Error:      ErrUndefinedVariable,
		},
		{
			Name:       "returns an error if the value is not an integer",
			Given:      "FOO",
			ShouldFail: true,
			Env: map[string]string{
				"FOO": "notaninteger",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.PopulateEnv(t)

			received, err := WithInt16(tc.Given)

			if tc.ShouldFail {
				assert.Error(t, err)
				if tc.Error != nil {
					assert.ErrorIs(t, err, tc.Error)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.Expected, received)
			}
		})
	}
}

func TestWithInt32(t *testing.T) {
	cases := []testutils.TestCase[int32]{
		{
			Name:     "correctly retrieves the value",
			Given:    "FOO",
			Expected: 2147483647, // maximum value for int32
			Env: map[string]string{
				"FOO": "2147483647",
			},
		},
		{
			Name:       "returns ErrUndefinedVariable if the env var is not defined",
			Given:      "FOO",
			ShouldFail: true,
			Error:      ErrUndefinedVariable,
		},
		{
			Name:       "returns an error if the value is not an integer",
			Given:      "FOO",
			ShouldFail: true,
			Env: map[string]string{
				"FOO": "notaninteger",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.PopulateEnv(t)

			received, err := WithInt32(tc.Given)

			if tc.ShouldFail {
				assert.Error(t, err)
				if tc.Error != nil {
					assert.ErrorIs(t, err, tc.Error)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.Expected, received)
			}
		})
	}
}

func TestWithInt64(t *testing.T) {
	cases := []testutils.TestCase[int64]{
		{
			Name:     "correctly retrieves the value",
			Given:    "FOO",
			Expected: 9223372036854775807, // maximum value for int64
			Env: map[string]string{
				"FOO": "9223372036854775807",
			},
		},
		{
			Name:       "returns ErrUndefinedVariable if the env var is not defined",
			Given:      "FOO",
			ShouldFail: true,
			Error:      ErrUndefinedVariable,
		},
		{
			Name:       "returns an error if the value is not an integer",
			Given:      "FOO",
			ShouldFail: true,
			Env: map[string]string{
				"FOO": "notaninteger",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.PopulateEnv(t)

			received, err := WithInt64(tc.Given)

			if tc.ShouldFail {
				assert.Error(t, err)
				if tc.Error != nil {
					assert.ErrorIs(t, err, tc.Error)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.Expected, received)
			}
		})
	}
}

func TestWithFloat32(t *testing.T) {
	cases := []testutils.TestCase[float32]{
		{
			Name:     "correctly retrieves the value",
			Given:    "FOO",
			Expected: 0.1,
			Env: map[string]string{
				"FOO": "0.1",
			},
		},
		{
			Name:       "returns ErrUndefinedVariable if the env var is not defined",
			Given:      "FOO",
			ShouldFail: true,
			Error:      ErrUndefinedVariable,
		},
		{
			Name:       "returns an error if the value is not a float32",
			Given:      "FOO",
			ShouldFail: true,
			Env: map[string]string{
				"FOO": "notafloat",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.PopulateEnv(t)

			received, err := WithFloat32(tc.Given)

			if tc.ShouldFail {
				assert.Error(t, err)
				if tc.Error != nil {
					assert.ErrorIs(t, err, tc.Error)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.Expected, received)
			}
		})
	}
}

func TestWithFloat64(t *testing.T) {
	cases := []testutils.TestCase[float64]{
		{
			Name:     "correctly retrieves the value",
			Given:    "FOO",
			Expected: 0.1,
			Env: map[string]string{
				"FOO": "0.1",
			},
		},
		{
			Name:       "returns ErrUndefinedVariable if the env var is not defined",
			Given:      "FOO",
			ShouldFail: true,
			Error:      ErrUndefinedVariable,
		},
		{
			Name:       "returns an error if the value is not a float64",
			Given:      "FOO",
			ShouldFail: true,
			Env: map[string]string{
				"FOO": "notafloat",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.PopulateEnv(t)

			received, err := WithFloat64(tc.Given)

			if tc.ShouldFail {
				assert.Error(t, err)
				if tc.Error != nil {
					assert.ErrorIs(t, err, tc.Error)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.Expected, received)
			}
		})
	}
}
