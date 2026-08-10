package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	var pool []string
	next := 0

	for sc.Scan() {
		fields := strings.Fields(sc.Text())
		if len(fields) == 0 {
			continue
		}

		switch fields[0] {
		case "POOL":
			pool = append([]string(nil), fields[1:]...)
			next = 0
			fmt.Println("OK")
		case "PICK":
			if len(pool) == 0 {
				fmt.Println("EMPTY")
				continue
			}

			fmt.Println(pool[next])
			next = (next + 1) % len(pool)
		case "RESET":
			next = 0
			fmt.Println("OK")
		}
	}
}
