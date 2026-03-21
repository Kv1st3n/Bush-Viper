# Bush-Viper
Bush-Viper is a simple CLI-based port scanner built in Golang. It was developed (and is still being developed) to explore networking. Currently, it provides tools for DNS lookup and port scanning (both single-port or wide-range)

[IMPORTANT]

This is an early build, with sequential scanning. Expect updates

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
2 = singl-port scan
3 = wide-range scan
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
go run . <mode> <hostname> <port>

# Example
go run . 2 127.0.0.1 80

```

### Using wide-range scan
``` 
# Command
go run . <mode> <hostname>

# Example
go run . 3 127.0.0.1

```

[DISCLAIMER]
This tool is for educational and authorized testing purposes only. Do not use it on targets that you do not own or have explicit permission to scan.