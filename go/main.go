package main

import (
	"fmt"
	"os"

	"github.com/monu1122/mycd/go/mycd"
)

func main() {
	ac := len(os.Args)
	if ac != 3 {
		fmt.Println("Number of arguments does not match")
		os.Exit(1)
	}
	z := mycd.NewMyCd(os.Args[1], os.Args[2])
	z.PrintDir()

}
