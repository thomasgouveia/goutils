package testutils

import "testing"

type TestCase[T any] struct {
	Name       string
	Env        map[string]string
	Given      string
	Expected   T
	ShouldFail bool
	Error      error
}

// PopulateEnv sets the test case environment variables for the given test.
func (tc *TestCase[T]) PopulateEnv(t *testing.T) {
	for k, v := range tc.Env {
		t.Setenv(k, v)
	}
}
