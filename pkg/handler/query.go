// // pkg/handler/query
package handler

import (
	"fmt"
	"os"
	"time"
	"back_go/pkg/zincsearch"
	"encoding/json"
	"net/http"
	// "log"
	// "io/ioutil"
	"bytes"
	"strings"
)

var max_results = 100
var search_type = "match"

type Email struct {
    From    string `json:"from"`
    To      string `json:"to"`
    Subject string `json:"subject"`
    Content string `json:"content"`
    FullFile    string `json:"full_file"`
}

func parseEmail(emailText string) (*Email, error) {
    email := &Email{}

    lines := strings.Split(emailText, "\n")
    var bodyStarted bool

    for _, line := range lines {
        if strings.HasPrefix(line, "From:") && email.From == "" {
            email.From = strings.TrimSpace(strings.TrimPrefix(line, "From:"))
        } else if strings.HasPrefix(line, "To:") && email.To == "" {
            email.To = strings.TrimSpace(strings.TrimPrefix(line, "To:"))
        } else if strings.HasPrefix(line, "Subject:") && email.Subject == "" {
            email.Subject = strings.TrimSpace(strings.TrimPrefix(line, "Subject:"))
        } else if line == "" && !bodyStarted {
            bodyStarted = true
        } else if bodyStarted {
            email.Content += line + "\n"
        }
		email.FullFile += line + "\n"
    }

    return email, nil
}

// Construct the request and perform the petition
func SearchDocuments(query string) ([]Email, error) {
	fmt.Println(query)
	LoadEnv()
	url := os.Getenv("ZINC_HOST") + ":" + os.Getenv("ZINC_PORT") + "/api/" + index + "/_search"

	now := time.Now()

	// Implementa la función para filtrar por tiempo --------------
	startTime := now.AddDate(0, -7, -7).Format("2006-01-02T15:04:05Z")
	endTime := now.Format("2006-01-02T15:04:05Z")

	request := zincsearch.SearchDocumentsRequest{
		SearchType: search_type,
		MaxResults: max_results,
		Query: zincsearch.SearchDocumentsRequestQuery{
			Term:      query,
			StartTime: startTime,
			EndTime:   endTime,
		},
	}

	jsonData, err := json.MarshalIndent(request, "", "   ")
	if err != nil {
		return nil, fmt.Errorf("Error al convertir a JSON: %v", err)
	}

	fmt.Println(string(jsonData))
	
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("Error al leer la solicitud: %v", err)
	}
	
	req.SetBasicAuth(os.Getenv("ZINC_ADMIN_USER"), os.Getenv("ZINC_ADMIN_PASSWORD"))
	req.Header.Set("Content-Type", "application/json")
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error al realizar la solicitud HTTP: %v", err)
	}
	
	defer resp.Body.Close()
	
	var searchResponse zincsearch.SearchDocumentsResponse
	fmt.Println(searchResponse)
	err = json.NewDecoder(resp.Body).Decode(&searchResponse)
	if err != nil {
		return nil, fmt.Errorf("Error al decodificar la respuesta JSON: %v", err)
	}

	var result []Email

	for _, value := range searchResponse.Hits.Hits {
		parsedEmail, err := parseEmail(fmt.Sprint(value.Source["contenido"]))
		if err != nil {
			// Manejar el error de parseo de alguna manera si es necesario
			fmt.Printf("Error al analizar el correo: %v\n", err)
			continue // Omitir este correo y continuar con el siguiente
		}
		result = append(result, *parsedEmail) // Agregar el Email al slice result
	}

	return result, nil
}