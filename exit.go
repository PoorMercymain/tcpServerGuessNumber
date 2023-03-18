package exit

import (
	"fmt"
	"os"
)

func EmergencyExit(exitCode int, v ...any) {
	fmt.Fprintln(os.Stderr, v...)
	os.Exit(exitCode)
}
