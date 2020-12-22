package main

import (
	"flag"
	"os"
	"fmt"

	"git.sr.ht/~otl/wildfly"
)

func main() {
	username := flag.String("u", "wildfly", "username")
	password := flag.String("p", "wildfly", "password")
	host := flag.String("h", "localhost", "host")
	warning := flag.Int("w", 50, "warning")
	critical := flag.Int("c", 75, "critical")
	flag.Parse()

	client := wildfly.NewClient(*host, *username, *password)
	mem, err := client.StatMemory()
	if err != nil {
		fmt.Println("fetch memory usage: ", err)
		os.Exit(3)
	}
	pctused := mem.Max / mem.Used
	fmt.Printf("%d%% of heap memory used (%d of %d bytes)", pctused, mem.Used, mem.Max)
	fmt.Printf(" | heap_used_percent=%d, heap_used_bytes=%d\n", pctused, mem.Used)

	if pctused > *critical {
		os.Exit(2)
	} else if pctused > *warning {
		os.Exit(1)
	}
}
