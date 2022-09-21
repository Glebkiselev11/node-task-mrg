package main

import (
	"fmt"

	ps "github.com/keybase/go-ps"
)

func main() {
	ps, err := ps.Processes()

	if err != nil {
		fmt.Print("Error in get processes %w", err)
		return
	}

	for pp := range ps {
		fmt.Printf("%d, %s\n", ps[pp].Pid(), ps[pp].Executable())
	}
}
