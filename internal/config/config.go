package config

type Config struct {
	DatabaseURL string `envconfig:"database_url" required:"true"`
	ServiceURL string `envconfig:"service_url" required:"true"`
	ServiceType string `envconfig:"service_type" required:"true"`
}