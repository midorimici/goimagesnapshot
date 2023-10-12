package printer

import "fmt"

const (
	ylw = "\x1b[33m"
	end = "\x1b[0m"
)

func Yellow(s string) {
	fmt.Printf("%s%s%s\n", ylw, s, end)
}

func Yellowf(format string, a ...any) {
	Yellow(fmt.Sprintf(format, a...))
}
