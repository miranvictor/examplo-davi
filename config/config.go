package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Config(key string) string {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Erro: não foi possivel ler o arquivo .env")
	}
	return os.Getenv(key)
}
