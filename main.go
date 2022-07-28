package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/d-sohan/avatargen/painter"
)

func main() {

	count, size := 5, 70

	for i, v := range os.Args {
		switch s := strings.TrimSpace(v); {
		case strings.HasPrefix(s, "--count="):
			c, err := strconv.Atoi(s[len("--count="):])
			if err != nil {
				fmt.Println("avatargen -count=<block count>")
				return
			}
			count = c
		case strings.HasPrefix(s, "--size="):
			sz, err := strconv.Atoi(s[len("--size="):])
			if err != nil {
				fmt.Println("avatargen -size=<block size>")
				return
			}
			size = sz
		default:
			if i != 0 {
				fmt.Printf("avatargen: unrecognized option '%v'\n", s)
				fmt.Println("Try '--count=<block count>' and/or '--size=<block size in pixel>'")
				return
			}
		}
	}

	painter.Paint(count, size)

}
