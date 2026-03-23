package main

type Service struct {
	Name     string
	Port     int
	Protocol string
	Probe    []byte
}

var ServiceDB = []Service{

	{
		Name:     "FTP",
		Port:     21,
		Protocol: "tcp",
	},

	{
		Name:     "SSH",
		Port:     22,
		Protocol: "tcp",
	},

	{
		Name:     "SMTP",
		Port:     25,
		Protocol: "tcp",
	},

	{
		Name:     "DNS",
		Port:     53,
		Protocol: "tcp",
	},

	{
		Name:     "HTTP",
		Port:     80,
		Protocol: "tcp",
		Probe:    []byte("GET / HTTP/1.1\r\nHost: %s\r\n\r\n"),
	},

	{
		Name:     "HTTP",
		Port:     8080,
		Protocol: "tcp",
		Probe:    []byte("GET / HTTP/1.1\r\nHost: %s\r\n\r\n"),
	},

	{
		Name:     "HTTP",
		Port:     8888,
		Protocol: "tcp",
		Probe:    []byte("GET / HTTP/1.1\r\nHost: %s\r\n\r\n"),
	},

	{
		Name:     "Redis",
		Port:     6379,
		Protocol: "tcp",
		Probe:    []byte("INFO\r\nQUIT\r\n"),
	},

	{
		Name:     "MySQL",
		Port:     3306,
		Protocol: "tcp",
	},

	{
		Name:     "PostgreSQL",
		Port:     5432,
		Protocol: "tcp",
		Probe:    []byte("\x00\x00\x00\x08\x04\xd2\x16\x2f"),
	},

	{
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
