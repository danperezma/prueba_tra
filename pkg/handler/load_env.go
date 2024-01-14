package handler

import(
	"github.com/joho/godotenv"
	"log"
	"bytes"
	"fmt"
	"encoding/json"
)

func LoadEnv(){
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}
}