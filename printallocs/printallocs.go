package printallocs

import (
	"fmt"
	"runtime"
)

// Prints allocated memory
// to prevent memory leak while working with slices we can copy part that we need or clear the rest of the data we don't need
// e.g.
// clear(data[2:])
// return data[:2]

func printAllocs() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/1024/1024)
}
