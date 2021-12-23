package server

type Args struct {
	Port     int
	Refresh  int
	Username string
	Password string
	Server   string
	IPAddr   string
	Insecure bool
}
