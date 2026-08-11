package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type backend struct {
	name    string
	weight  int
	current int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	var pool []backend
	totalWeight := 0

	pick := func() string {
		var best *backend
		for i := range pool {
			b := &pool[i]
			b.current += b.weight
			if best == nil || b.current > best.current {
				best = b
			}
		}
		best.current -= totalWeight
		return best.name
	}

	for sc.Scan() {
		fields := strings.Fields(sc.Text())
		if len(fields) == 0 {
			continue
		}

		switch fields[0] {
		case "POOL":
			pool = pool[:0]
			totalWeight = 0
			for _, entry := range fields[1:] {
				name, weightStr, ok := strings.Cut(entry, ":")
				if !ok {
					continue
				}
				w, err := strconv.Atoi(weightStr)
				if err != nil {
					continue
				}
				pool = append(pool, backend{name: name, weight: w})
				totalWeight += w
			}
			fmt.Println("OK")
		case "PICK":
			if len(pool) == 0 {
				fmt.Println("EMPTY")
				continue
			}
			fmt.Println(pick())
		case "PICKN":
			if len(pool) == 0 {
				fmt.Println("EMPTY")
				continue
			}
			n, err := strconv.Atoi(fields[1])
			if err != nil || n < 0 {
				continue
			}
			names := make([]string, n)
			for i := 0; i < n; i++ {
				names[i] = pick()
			}
			fmt.Println(strings.Join(names, ","))
		case "RESET":
			for i := range pool {
				pool[i].current = 0
			}
			fmt.Println("OK")
		}
	}
}
