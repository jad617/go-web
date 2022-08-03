package configs

import (
	"go-web/env"
)

type Config struct {
	TemplateDir string
}

func FetchVars() *Config {
	return &Config{
		TemplateDir: env.GetEnv("TEMPLATE_DIR", "/var/www/html/"),
	}
}
