# Bush-Viper
Bush-Viper is a simple CLI-based port scanner built in Golang. It was developed (and is still being developed) to explore networking. Currently, it provides tools for DNS lookup and port scanning (both single-port or wide-range)

>[!IMPORTANT]
> This is an early build. 
>Expect updates for improvement in functionality and performance.

## Installing

Prerequisites
- You must have Golang installed

``` 
# Clone the Git
git clone https://github.com/Kv1st3n/Bush-Viper.git

# Change directory
cd bush-viper

```

## Using Bush-Viper
To run bush-viper use 'go run .' once in the directory, there are three available modes: DNS, single-port, and wide-range scans

```
1 = DNS lookup
2 = single-port scan
3 = wide-range scan
4 = Manual (W.I.P.)
```

### Using DNS lookup
``` 
# Command
go run . <mode> <hostname>

# Example
go run . 1 example.com

```

### Using single-port scan
``` 
# Command
go run . <mode> <IP> <port>

# Example
go run . 2 127.0.0.1 80

```

**Mode 3 defaults to ports 1-65535 if only the IP-address is provided but you can optionally specify a custom range**
### Using wide-range scan
``` 
# Basic Command (Default)
go run . <mode> <IP>

# Example
go run . 3 127.0.0.1

# Custom Range Command (Scans only ports 80 through 443)
go run . 3 127.0.0.1 80 443

```

>[!WARNING]
>This tool is for educational and authorized testing purposes only. Do not use it on targets that you do not own or have explicit permission to scan.