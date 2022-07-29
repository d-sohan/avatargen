package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/d-sohan/avatargen/painter"
)

func help() {
	fmt.Println(
		"usage: avatargen [options]\n",
		"Options:\n",
		"--count=<block count at least 2>\t number of blocks, default 5\n",
		"--size=<block size at least 2>\t\t size of each block in pixels, default 70\n",
		"--fgcolor=<hex value #000000 to #ffffff>\t color of each block, default #64C8C8\n",
		"--bgcolor=<hex value #000000 to #ffffff>\t color of background, default #ffffff\n",
		"--output=<path of output image>\t default 'avatar.png' in the current directory",
	)
}

func main() {

	count, size := 5, 70
	fgcolor, bgcolor, output := "64c8c8", "ffffff", "avatar.png"

	for i, v := range os.Args {
		switch s := strings.TrimSpace(v); {
		case strings.HasPrefix(s, "--count="):
			c, err := strconv.Atoi(s[len("--count="):])
			if err != nil || c <= 1 {
				fmt.Println("avatargen: --count=<block count at least 2>")
				return
			}
			count = c
		case strings.HasPrefix(s, "--size="):
			sz, err := strconv.Atoi(s[len("--size="):])
			if err != nil || sz <= 1 {
				fmt.Println("avatargen: --size=<block size at least 2>")
				return
			}
			size = sz
		case strings.HasPrefix(s, "--fgcolor="):
			hc := s[len("--fgcolor="):]
			if len(hc) == 7 && hc[0] == '#' {
				hc = hc[1:]
			}
			if len(hc) == 6 && hc[0] != '#' {
				fgcolor = strings.ToLower(hc)
			} else {
				fmt.Println("avatargen: --fgcolor=<hex value 000000 to ffffff>")
				return
			}
		case strings.HasPrefix(s, "--bgcolor="):
			hc := s[len("--bgcolor="):]
			if len(hc) == 7 && hc[0] == '#' {
				hc = hc[1:]
			}
			if len(hc) == 6 && hc[0] != '#' {
				bgcolor = strings.ToLower(hc)
			} else {
				fmt.Println("avatargen: --bgcolor=<hex value 000000 to ffffff>")
				return
			}
		case strings.HasPrefix(s, "--output="):
			output = s[len("--output="):]
		case strings.HasPrefix(s, "--help"):
			help()
		default:
			if i != 0 {
				fmt.Printf("avatargen: unrecognized option '%v'\n", s)
				fmt.Println("Try '--help' for help")
				return
			}
		}
	}

	if err := painter.Paint(count, size, fgcolor, bgcolor, output); err != nil {
		fmt.Println("avatargen:", err)
	}

}
