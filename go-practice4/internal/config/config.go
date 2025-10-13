package config

type Config struct {
    DSN string
}

func New() *Config {
    return &Config{
        DSN: "postgres://user:password@localhost:5430/mydatabase?sslmode=disable",
    }
}
