package config

type HTTPConfig struct {
	Port string
}

type AMQPConfig struct {
	DSN string
}

type ClickHouseConfig struct {
	DSN string
}

type PostgresConfig struct {
	DSN string
}

type Config struct {
	AmqpConfig       AMQPConfig
	HttpConfig       HTTPConfig
	ClickHouseConfig ClickHouseConfig
	PostgresConfig   PostgresConfig
}

type ConfigLoader interface {
	Load() Config
}
