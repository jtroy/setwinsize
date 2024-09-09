package main

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

func main() {
	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Winsize: %+v\n", *ws)

	err = unix.IoctlSetWinsize(int(os.Stdout.Fd()), unix.TIOCSWINSZ, ws)
	if err != nil {
		panic(err)
	}
}
