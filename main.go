package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	failThreshold = 3
	okThreshold   = 2
)

type backend struct {
	name  string
	up    bool
	fails int
	oks   int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	var pool []*backend

	for sc.Scan() {
		fields := strings.Fields(sc.Text())
		if len(fields) == 0 {
			continue
		}

		switch fields[0] {
		case "POOL":
			pool = pool[:0]
			for _, name := range fields[1:] {
				pool = append(pool, &backend{name: name, up: true})
			}
			fmt.Println("OK")
		case "REPORT":
			if len(fields) < 3 {
				fmt.Println("ERR")
				continue
			}
			var b *backend
			for _, p := range pool {
				if p.name == fields[1] {
					b = p
					break
				}
			}
			if b == nil || (fields[2] != "OK" && fields[2] != "FAIL") {
				fmt.Println("ERR")
				continue
			}
			if fields[2] == "OK" {
				b.oks++
				b.fails = 0
				if b.oks >= okThreshold {
					b.up = true
				}
			} else {
				b.fails++
				b.oks = 0
				if b.fails >= failThreshold {
					b.up = false
				}
			}
			fmt.Println("OK")
		case "STATUS":
			for _, b := range pool {
				state := "DOWN"
				if b.up {
					state = "UP"
				}
				fmt.Printf("%s %s\n", b.name, state)
			}
		case "HEALTHY":
			var names []string
			for _, b := range pool {
				if b.up {
					names = append(names, b.name)
				}
			}
			if len(names) == 0 {
				fmt.Println("none")
			} else {
				fmt.Println(strings.Join(names, ","))
			}
		}
	}
}
