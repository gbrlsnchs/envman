package envman_test

import (
	"fmt"
	"os"

	"github.com/gbrlsnchs/envman"
)

func Example() {
	// mocked environment variable
	err := os.Setenv("MY_ENV_VAR", "1337")

	if err != nil {
		// handle error
	}

	err = envman.SetInt(
		"MY_ENV_VAR",
		"MY_FOO_BAR_VAR",
	)

	if err != nil {
		// handle error
	}

	myVar := envman.Get("MY_ENV_VAR")
	foobar := envman.Get("MY_FOO_BAR_VAR")

	fmt.Println(myVar)
	fmt.Println(foobar)

	// Output: 1337
	// <nil>
}
