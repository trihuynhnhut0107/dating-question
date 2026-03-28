package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)
func LoadConfig() {
    _, filename,_,ok := runtime.Caller(0)
    if !ok {
		log.Println("Could not determine config file location, falling back to CWD")
		godotenv.Load()
		return
	}

    projectRoot:=filepath.Dir(filepath.Dir(filename))
    envPath:=filepath.Join(projectRoot,".env")
    if err := godotenv.Load(envPath); err != nil {
		log.Printf("No .env file found at %s, relying on system environment variables", envPath)
	} else {
		log.Printf("Loaded config from %s", envPath)
	}    
}
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}