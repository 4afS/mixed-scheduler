package presenter

import (
	"fmt"
	"os"
)

func PrintErr(err error) {
	fmt.Fprintf(os.Stderr, "mxs: %s\n", err.Error())
	os.Exit(1)
}
