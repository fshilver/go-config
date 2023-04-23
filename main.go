package main

import (
	"fmt"
)

func main() {
	parseFlags()

	fmt.Printf("Host: %s\n", cfg.Host)
	fmt.Printf("Port: %d\n", cfg.Port)
	// Access more configuration values as needed
}
