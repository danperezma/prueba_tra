package handler

import(
	"github.com/joho/godotenv"
	"fmt"
)

// Function to load the .env file
func LoadEnv(){
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error al cargar el archivo .env")
	}
}