# Bush-Viper
Bush-Viper is a simple CLI-based port scanner built in Golang. It was developed (and is still being developed) to explore networking. Currently, it provides tools for DNS lookup and port scanning (both single-port or wide-range)

>[!IMPORTANT]
> This is an early build. 
>Expect updates for improvement in functionality and performance.

## Installing

### Prerequisites
- You must have Golang installed

```zsh
# Clone the Git
git clone https://github.com/Kv1st3n/Bush-Viper.git

# Change directory
cd bush-viper

```

## Using Bush-Viper
Bush-Viper uses flags for configuration. This means you use specific keys like `-mode` or `ip` to pass your arguments. To see all available flags run:


```
# Help
go run . -help

# Available modes
1 = DNS lookup
2 = single-port scan
3 = wide-range scan
```

### Using DNS lookup
Pass the target hostname to the `-ip` flag.
```zsh 
# Command
go run . -mode 1 -ip <hostname>

# Example
go run . -mode 1 -ip example.com

```

### Using single-port scan
Specify the target IP and a single port.
```zsh
# Command
go run . -mode 2 -ip <IP> -port <port>

# Example
go run . -mode 2 -ip 127.0.0.1 -port 80

```

### Using wide-range scan
**Mode 3 defaults to ports 1-65535 if only the IP-address is provided but you can optionally specify a custom range using a hyphen**
```zsh
# Basic Command (Default)
go run . -mode 3 -ip <IP>

# Example
go run . -mode 3 -ip 127.0.0.1

# Custom Range Command (Scans only ports 80 through 443)
go run . -mode 3 -ip 127.0.0.1 -port 80-443

```

>[!WARNING]
>This tool is for educational and authorized testing purposes only. Do not use it on targets that you do not own or have explicit permission to scan.