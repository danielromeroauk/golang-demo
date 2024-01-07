package exectime

import (
	"fmt"
	"time"
)

func Log(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}
