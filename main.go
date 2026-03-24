package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: bush-viper <mode> <ip> <port/prefix>")
		fmt.Println("Modes: 1 (DNS), 2 (Single), 3 (Wide/Rust)")
		return
	}

	modeChoice := os.Args[1]

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

	switch modeChoice {
	case "1":

		if len(os.Args) < 3 {
			fmt.Println("Error: Mode 1 requires a hostname. Example: ./bush-viper 1 google.com")
			return
		}
		host := os.Args[2]
		ipAddress, _ := getHostAddress(host)
		fmt.Printf("DNS result for %s: %v\n", host, ipAddress)

	case "2":

		if len(os.Args) < 4 {
			fmt.Println("Error: Mode 2 requires IP and Port. Example: ./bush-viper 2 127.0.0.1 80")
		}

		ip := os.Args[2]
		port := os.Args[3]

		result, err := singlePortScan(ip, port)
		if err != nil || result.Opened == "" {
			fmt.Printf("[-] %s:%s is CLOSED or Unreachable\n", ip, port)
		} else {
			fmt.Printf("[+] %s:%s is OPEN!\n", result.IP, result.Port)

			if result.Banner != "" {
				fmt.Printf("    Banner: %q\n", result.Banner)
			} else {
				fmt.Println("    Banner: <No response or no probe matches>")
			}
		}

	case "3":

		if len(os.Args) < 3 {
			fmt.Println("Error: Mode 3 requires IP. Example: ./bush-viper 3 127.0.0.1")
			return
		}

		ip := os.Args[2]
		fmt.Printf("[*] Starting wide port scan on %s (Ports 1-65536)...\n", ip)

		results, err := widePortScan(ip)
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

	case "4":
		fmt.Println("Quiting Viper...")
	}

}
