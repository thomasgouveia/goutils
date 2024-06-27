package env

import (
	"github.com/stretchr/testify/assert"
	"github.com/thomasgouveia/goutils/env/testutils"
	"testing"
)

func TestWithDefaultString(t *testing.T) {
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
			Name:     "returns the fallback value if the env var is not defined",
			Given:    "FOO",
			Expected: "fallback",
		},
		{
			Name:     "returns the fallback value if the env var is empty",
			Given:    "FOO",
			Expected: "bar",
			Env: map[string]string{
				"FOO": "",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.PopulateEnv(t)
			assert.Equal(t, tc.Expected, WithDefaultString(tc.Given, tc.Expected))
		})
	}
}

func TestWithDefaultBool(t *testing.T) {
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
			Name:     "returns the fallback value if the env var is not defined",
			Given:    "FOO",
			Expected: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.PopulateEnv(t)
			assert.Equal(t, tc.Expected, WithDefaultBool(tc.Given, tc.Expected))
		})
	}
}

func TestWithDefaultInt(t *testing.T) {
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
			Name:     "returns the fallback value if the env var is not defined",
			Given:    "FOO",
			Expected: 10,
		},
		{
			Name:     "returns the fallback value if the parsing fails",
			Given:    "FOO",
			Expected: 10,
			Env: map[string]string{
				"FOO": "notaninteger",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.PopulateEnv(t)
			assert.Equal(t, tc.Expected, WithDefaultInt(tc.Given, tc.Expected))
		})
	}
}

func TestWithDefaultInt8(t *testing.T) {
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
			Name:     "returns the fallback value if the env var is not defined",
			Given:    "FOO",
			Expected: 10,
		},
		{
			Name:     "returns the fallback value if the parsing fails",
			Given:    "FOO",
			Expected: 10,
			Env: map[string]string{
				"FOO": "notaninteger",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.PopulateEnv(t)
			assert.Equal(t, tc.Expected, WithDefaultInt8(tc.Given, tc.Expected))
		})
	}
}

func TestWithDefaultInt16(t *testing.T) {
	cases := []testutils.TestCase[int16]{
		{
			Name:     "correctly retrieves the value",
			Given:    "FOO",
			Expected: 32767,
			Env: map[string]string{
				"FOO": "32767",
			},
		},
		{
			Name:     "returns the fallback value if the env var is not defined",
			Given:    "FOO",
			Expected: 32767,
		},
		{
			Name:     "returns the fallback value if the parsing fails",
			Given:    "FOO",
			Expected: 32767,
			Env: map[string]string{
				"FOO": "32768",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.PopulateEnv(t)
			assert.Equal(t, tc.Expected, WithDefaultInt16(tc.Given, tc.Expected))
		})
	}
}

func TestWithDefaultInt32(t *testing.T) {
	cases := []testutils.TestCase[int32]{
		{
			Name:     "correctly retrieves the value",
			Given:    "FOO",
			Expected: 2147483647,
			Env: map[string]string{
				"FOO": "2147483647",
			},
		},
		{
			Name:     "returns the fallback value if the env var is not defined",
			Given:    "FOO",
			Expected: 2147483647,
		},
		{
			Name:     "returns the fallback value if the parsing fails",
			Given:    "FOO",
			Expected: 2147483647,
			Env: map[string]string{
				"FOO": "2147483648",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.PopulateEnv(t)
			assert.Equal(t, tc.Expected, WithDefaultInt32(tc.Given, tc.Expected))
		})
	}
}

func TestWithDefaultInt64(t *testing.T) {
	cases := []testutils.TestCase[int64]{
		{
			Name:     "correctly retrieves the value",
			Given:    "FOO",
			Expected: 9223372036854775807,
			Env: map[string]string{
				"FOO": "9223372036854775807",
			},
		},
		{
			Name:     "returns the fallback value if the env var is not defined",
			Given:    "FOO",
			Expected: 9223372036854775807,
		},
		{
			Name:     "returns the fallback value if the parsing fails",
			Given:    "FOO",
			Expected: 9223372036854775807,
			Env: map[string]string{
				"FOO": "9223372036854775808",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.PopulateEnv(t)
			assert.Equal(t, tc.Expected, WithDefaultInt64(tc.Given, tc.Expected))
		})
	}
}

func TestWithDefaultFloat32(t *testing.T) {
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
			Name:     "returns the fallback value if the env var is not defined",
			Given:    "FOO",
			Expected: 2.5,
		},
		{
			Name:     "returns the fallback value if the parsing fails",
			Given:    "FOO",
			Expected: 2.5,
			Env: map[string]string{
				"FOO": "notafloat",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.PopulateEnv(t)
			assert.Equal(t, tc.Expected, WithDefaultFloat32(tc.Given, tc.Expected))
		})
	}
}

func TestWithDefaultFloat64(t *testing.T) {
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
			Name:     "returns the fallback value if the env var is not defined",
			Given:    "FOO",
			Expected: 2.5,
		},
		{
			Name:     "returns the fallback value if the parsing fails",
			Given:    "FOO",
			Expected: 2.5,
			Env: map[string]string{
				"FOO": "notafloat",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.PopulateEnv(t)
			assert.Equal(t, tc.Expected, WithDefaultFloat64(tc.Given, tc.Expected))
		})
	}
}
