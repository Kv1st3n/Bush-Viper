package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {

	mode := flag.Int("mode", 1, "Modes: 1 (DNS), 2 (Single), 3 (Wide/Rust)")
	ip := flag.String("ip", "127.0.0.1", "Target IP Adress")
	port := flag.String("port", "80", "Port (e.g., '80' for mode 2, or '80-90' for mode 3)")

	flag.Parse()

	const (
		ColorGreen = "\033[32m"
		ColorReset = "\033[0m"
	)

	const banner = `
		 ____  _   _ ____  _   _  __     _____ ____  _____ ____  
		| __ )| | | / ___|| | | | \ \   / /_ _|  _ \| ____|  _ \ 
		|  _ \| | | \___ \| |_| |  \ \ / / | || |_) |  _| | |_) |
		| |_) | |_| |___) |  _  |   \ V /  | ||  __/| |___|  _ < 
		|____/ \___/|____/|_| |_|    \_/  |___|_|   |_____|_| \_\
	`

	const underLine = `
		_____________________________________________________________
	`

	fmt.Println(ColorGreen + underLine + ColorReset)
	fmt.Println(ColorGreen + banner + ColorReset)
	fmt.Println(ColorGreen + underLine + ColorReset)
	fmt.Println("V1.0.0 - Starting Viper Engine...")

	switch *mode {
	case 1:

		ipAddress, _ := getHostAddress(*ip)
		fmt.Printf("DNS result for %s: %v\n", *ip, ipAddress)

	case 2:

		result, err := singlePortScan(*ip, *port)

		if err != nil || result.Opened == "" {
			fmt.Printf("[-] %s:%s is CLOSED or Unreachable\n", *ip, *port)
		} else {
			fmt.Printf("[+] %s:%s is OPEN!\n", result.IP, result.Port)

			if result.Banner != "" {
				fmt.Printf("    Banner: %q\n", result.Banner)
			}
		}

	case 3:
		startPort, endPort := "1", "65536"

		if strings.Contains(*port, "-") {
			parts := strings.Split(*port, "-")
			startPort = parts[0]
			endPort = parts[1]
		} else {
			startPort = *port
		}

		fmt.Printf("[*] Starting wide port scan on %s (Ports %s-%s)...\n", *ip, startPort, endPort)

		results, err := widePortScan(*ip, startPort, endPort)
		if err != nil {
			fmt.Printf("[-] Error during wide scan: %v\n", err)
			return
		}

		if len(results) == 0 {
			fmt.Println("[-] No open ports found in the specified range.")
		} else {
			fmt.Printf("[+] Found %d open ports:\n", len(results))

			for _, res := range results {
				if res.Banner != "" {
					fmt.Printf("    - Port %-5s: %-7s | Banner: %q\n", res.Port, res.Opened, res.Banner)
				} else {
					fmt.Printf("    - Port %-5s: %-7s | No banner captured\n", res.Port, res.Opened)
				}
			}
		}

	default:
		fmt.Printf("Unknown mode: %d. Use -help for usage.\n", *mode)
	}

}
