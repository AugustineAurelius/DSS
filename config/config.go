package config

var DefaultConfig = Config{
	Network: "tcp",
	Address: ":4000",
}

type Config struct {
	Network string
	Address string
}
