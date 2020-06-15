package pkg

import (
	"fmt"
	"os"
)

func Check(err error, args ...interface{}) {
	if err == nil {
		return
	}

	if len(args) == 0 {
		Fatalf("An unexpected error occurred: %+v", err)
	}

	if len(args) == 1 {
		Fatalf("%s", args[0])
	}

	Fatalf(fmt.Sprintf("%s", args[0]), args[1:]...)
}

func MustGetEnv(key string) (v string) {
	v = os.Getenv(key)
	if len(v) == 0 {
		Fatalf(`Environment variable "%s" must be set.`, key)
	}
	return
}

func Fatalf(message string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, message+"\n", args...)
	os.Exit(1)
}
