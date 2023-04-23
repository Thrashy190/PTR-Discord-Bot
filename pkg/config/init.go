package config

import "log"

func Init() {
	log.Println("Initializing config...")
	loadEnv()
	log.Println("Config initialized")
}
