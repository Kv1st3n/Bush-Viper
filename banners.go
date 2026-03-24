package main

type Service struct {
	Name     string
	Port     int
	Protocol string
	Probe    []byte
}

var ServiceDB = map[int]Service{
	21: {
		Name:     "FTP",
		Port:     21,
		Protocol: "tcp",
	},
	22: {
		Name:     "SSH",
		Port:     22,
		Protocol: "tcp",
	},
	25: {
		Name:     "SMTP",
		Port:     25,
		Protocol: "tcp",
	},
	53: {
		Name:     "DNS",
		Port:     53,
		Protocol: "tcp",
	},
	80: {
		Name:     "HTTP",
		Port:     80,
		Protocol: "tcp",
		Probe:    []byte("GET / HTTP/1.1\r\nHost: localhost\r\n\r\n"),
	},
	3306: {
		Name:     "MySQL",
		Port:     3306,
		Protocol: "tcp",
	},
	5432: {
		Name:     "PostgreSQL",
		Port:     5432,
		Protocol: "tcp",
		Probe:    []byte("\x00\x00\x00\x08\x04\xd2\x16\x2f"),
	},
	6379: {
		Name:     "Redis",
		Port:     6379,
		Protocol: "tcp",
		Probe:    []byte("INFO\r\nQUIT\r\n"),
	},
	8080: {
		Name:     "HTTP",
		Port:     8080,
		Protocol: "tcp",
		Probe:    []byte("GET / HTTP/1.1\r\nHost: localhost\r\n\r\n"),
	},
	8888: {
		Name:     "HTTP",
		Port:     8888,
		Protocol: "tcp",
		Probe:    []byte("GET / HTTP/1.1\r\nHost: localhost\r\n\r\n"),
	},
	27017: {
		Name:     "MongoDB",
		Port:     27017,
		Protocol: "tcp",
		Probe: []byte{
			0x13, 0x00, 0x00, 0x00,
			0x10,
			0x69, 0x73, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x00,
			0x01, 0x00, 0x00, 0x00,
			0x00,
		},
	},
}
