package config

import "os"

type AuthConfig struct {
	JwtSecret string
	JwtExpiry string
}

func LoadAuthConfig() AuthConfig {
	return AuthConfig{
		JwtSecret: os.Getenv("JWT_SECRET"),
		JwtExpiry: os.Getenv("JWT_EXPIRED"),
	}
}
