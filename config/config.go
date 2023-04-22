package config

import "github.com/joho/godotenv"

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}