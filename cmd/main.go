
package main

import (
    "fmt"
    // "os"
    "net/http"
    "back_go/pkg/handler"
    "io/ioutil"
    // "strings"
	"encoding/json"
    //"path/filepath"
    "github.com/go-chi/chi"
	"github.com/rs/cors"
)


func main() {

    // r := chi.NewRouter()

    // r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    //     fmt.Fprint(w, "¡Hola, mundo!")
    // })

    // http.ListenAndServe(":8080", r)

    r := chi.NewRouter()
	
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	// Agregar CORS a las rutas de chi
	r.Use(corsMiddleware.Handler)
	// Configurar tus rutas aquí
	r.Post("/api/query", HandlerFunc)
	

	// Escuchar en el puerto 8080 (o el puerto que desees)
	port := 8080
	fmt.Printf("Servidor escuchando en el puerto %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
    // handler.SearchDocuments("bike")

    // handler.sea()
	// fmt.Println("hi world")
}

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	// Lógica para manejar la solicitud POST
	fmt.Println("Se recibió una solicitud POST en", r.URL.Path)

	// Verificar el método de solicitud
	if r.Method != http.MethodPost {
		http.Error(w, "Método de solicitud no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Leer el cuerpo de la solicitud
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusInternalServerError)
		return
	}

	// Decodificar el cuerpo JSON en un mapa genérico
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Error al decodificar el JSON del cuerpo", http.StatusBadRequest)
		return
	}

	// Acceder al campo searchTerm
	searchTerm, ok := data["searchTerm"].(string)
	if !ok {
		http.Error(w, "Campo 'searchTerm' no encontrado o no es una cadena", http.StatusBadRequest)
		return
	}

	res, err := handler.SearchDocuments(searchTerm)
	if err != nil {
		fmt.Println("Error al procesar la solicitud de consulta", err)
		http.Error(w, "Error al procesar la solicitud de consulta", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
		return
	}

	// Escribir la respuesta JSON en la respuesta HTTP
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
