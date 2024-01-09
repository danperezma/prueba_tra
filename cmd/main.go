
package main

import (
    // "fmt"
    // _"os"
    //_"net/http"
    "back_go/pkg/handler"
    // "path/filepath"
    //_"github.com/go-chi/chi"
)


func main() {

    // r := chi.NewRouter()

    // r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    //     fmt.Fprint(w, "¡Hola, mundo!")
    // })

    // http.ListenAndServe(":8080", r)

    handler.Index()
	// fmt.Println("hi world")
}
