package config

import "github.com/joho/godotenv"

const Secret = "Hatsunemiku123!"

func GetSecretConfig() string {
	godotenv.Load()
	secret := "Hatsunemiku123!"

	return secret
}
