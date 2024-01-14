// pk/handler/indexer
package handler

import (
	"fmt"
	"os"
	"path/filepath"
	"back_go/pkg/zincsearch"
	"encoding/json"
	"net/http"
	"log"
	"io/ioutil"
	"bytes"
)

var Counter int = 0
var Emails []map[string]interface{}
var index string = "enronJELM"

// Construct the request and perform the petition
func Indexer(){ 
	LoadEnv()
	url := os.Getenv("ZINC_HOST") + ":" + os.Getenv("ZINC_PORT") + "/api/_bulkv2"
	request := zincsearch.CreateDocumentsRequest{
		Index: index,
		Records: Emails,
	}

	// Convertir el objeto a formato JSON para impresión
	jsonData, err := json.MarshalIndent(request, "", "   ")
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error reading request. ", err)
	}
	// Set headers
	req.SetBasicAuth(os.Getenv("ZINC_ADMIN_USER"), os.Getenv("ZINC_ADMIN_PASSWORD"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error al realizar la solicitud HTTP:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Código de estado inesperado: %d", resp.StatusCode)
	}

}

func Get_files(path string, f os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if !f.IsDir() {
		content_file, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println(err)
		}

		data := map[string]interface{}{
			"email_content": string(content_file),
		}
		Emails = append(Emails, data)

		Counter += 1
	}

	return nil
}

func Index(){
	err := filepath.Walk(os.Getenv("FILES_DIR"), Get_files)
	if err != nil {
		fmt.Println("Error al caminar por el directorio: %v\n", err) // Manejar el error de manera adecuada
	}
	fmt.Println(Counter)

	Indexer()
}
