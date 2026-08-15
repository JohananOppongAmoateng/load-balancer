package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type backend struct {
	name   string
	active int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	var pool []backend

	for sc.Scan() {
		fields := strings.Fields(sc.Text())
		if len(fields) == 0 {
			continue
		}

		switch fields[0] {
		case "POOL":
			pool = pool[:0]
			for _, name := range fields[1:] {
				pool = append(pool, backend{name: name})
			}
			fmt.Println("OK")
		case "PICK":
			if len(pool) == 0 {
				fmt.Println("EMPTY")
				continue
			}
			best := 0
			for i := 1; i < len(pool); i++ {
				if pool[i].active < pool[best].active {
					best = i
				}
			}
			pool[best].active++
			fmt.Println(pool[best].name)
		case "DONE":
			if len(fields) < 2 {
				continue
			}
			for i := range pool {
				if pool[i].name == fields[1] {
					if pool[i].active > 0 {
						pool[i].active--
					}
					break
				}
			}
			fmt.Println("OK")
		case "STATUS":
			type entry struct {
				name   string
				active int
			}
			entries := make([]entry, len(pool))
			for i, b := range pool {
				entries[i] = entry{name: b.name, active: b.active}
			}
			sort.Slice(entries, func(i, j int) bool {
				return entries[i].name < entries[j].name
			})
			for _, e := range entries {
				fmt.Printf("%s:%d\n", e.name, e.active)
			}
		}
	}
}
