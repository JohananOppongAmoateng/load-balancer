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

func find(pool []backend, name string) int {
	for i := range pool {
		if pool[i].name == name {
			return i
		}
	}
	return -1
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
			if len(fields) < 3 {
				fmt.Println("ERR")
				continue
			}
			x, y := find(pool, fields[1]), find(pool, fields[2])
			if x < 0 || y < 0 {
				fmt.Println("ERR")
				continue
			}
			best := x
			if pool[y].active < pool[x].active ||
				(pool[y].active == pool[x].active && pool[y].name < pool[x].name) {
				best = y
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
