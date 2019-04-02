package main

import "fmt"

const (
	// KB = "1000" + [itoa*3]byte.String() むりだった...
	KB int64 = 1000
	MB       = KB * KB
	GB       = MB * KB
	TB       = GB * KB
	PB       = TB * KB
	EB       = PB * KB
	ZB       = EB * KB
	YB       = ZB * KB
)

func main() {
	fmt.Println(Kb, MB, YB)
}
