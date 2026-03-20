package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var modeChoice int
	var host string

	var ip string
	var port string

	num := 0

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

	for num < 1 {
		fmt.Println("Modes of Operations")
		fmt.Println("1. DNS - Get Host Address:")
		fmt.Println("2. Port Scan - Single")
		fmt.Println("3. Port Scan - Wide")
		fmt.Println("4. Quit")
		fmt.Print("Select mode: ")
		fmt.Scan(&modeChoice)

		fmt.Println(modeChoice)

		switch modeChoice {
		case 1:
			fmt.Print("Enter host: ")

			host, _ = reader.ReadString('\n')
			host = strings.TrimSpace(host)

			if host != "" {
				ips, _ := getHostAddress(host)
				fmt.Println(ips)
			} else {
				fmt.Println("Error: Invalid string, can not be empty")
			}

		case 2:
			fmt.Println("2")

			fmt.Print("Enter IP-Address: ")
			fmt.Scan(&ip)

			fmt.Print("Enter specific port: ")
			fmt.Scan(&port)

			fmt.Printf("Scanning %s:%s...\n", ip, port) // Feedback for the user
			res, err := singlePortScan(ip, port)

			if err != nil {
				fmt.Printf("[-] Port %s is CLOSED (Error: %v)\n", port, err)
			} else {
				fmt.Printf("[+] Success! %s:%s is %v\n", res.IP, res.Port, res.Opened)
			}

		case 3:
			fmt.Println("3")
		case 4:
			num += 1
			fmt.Println("Quiting Viper...")
		}

	}

}
