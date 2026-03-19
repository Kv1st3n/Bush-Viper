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

	fmt.Println("Modes of Operations")
	fmt.Println("1. DNS - Get Host Address:")
	fmt.Println("2. Port Scan - Single")
	fmt.Println("3. Port Scan - Wide")
	fmt.Print("Select mode: ")
	fmt.Scan(&modeChoice)

	fmt.Println(modeChoice)

	switch modeChoice {
	case 1:
		fmt.Print("Enter host: ")

		host, _ = reader.ReadString('\n')
		host = strings.TrimSpace(host)
		ips, _ := getHostAddress(host)
		fmt.Println(ips)

	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	}
}
