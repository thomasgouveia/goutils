package env_test

import (
	"fmt"
	"github.com/thomasgouveia/goutils/env"
	"os"
)

func Example() {
	// Variables should not be declared here.
	// This is only for demonstration purposes.
	os.Setenv("FOO_STRING", "hello")
	os.Setenv("FOO_INT", "10")

	// To retrieve the values, you can use env.WithString and env.WithInt functions.
	// These functions will return an error if the environment variable is not set
	// or if the parsing of the value into the wanted type fails.
	strVal, _ := env.WithString("FOO_STRING")
	intVal, _ := env.WithInt("FOO_INT")

	fmt.Printf("FOO_STRING: %s\n", strVal) // Output: FOO_STRING: hello
	fmt.Printf("FOO_INT: %d\n", intVal)    // Output: FOO_INT: 10

	// In case you want to automatically fall back to a default value if the environment variable
	// is not set or if the parsing of the value into the wanted type fails, you can use the env.WithDefault* functions.
	fallbackVal := env.WithDefaultString("FOO_UNSET", "myfallbackvalue")

	fmt.Printf("FOO_UNSET: %s\n", fallbackVal) // Output: FOO_UNSET: hello
}
