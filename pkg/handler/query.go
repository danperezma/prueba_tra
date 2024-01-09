// // pkg/handler/query
// // funtion to perfom the query
package handler


// func performSearch(host string, query SearchRequest) (*Hits, error) {
// 	url := host + "/your_search_endpoint" // Reemplaza esto con la verdadera URL de búsqueda

// 	jsonData, err := json.Marshal(query)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		return nil, err
// 	}

// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var result map[string]interface{}
// 	if err := json.Unmarshal(body, &result); err != nil {
// 		return nil, err
// 	}

// 	hitsData, ok := result["hits"].(map[string]interface{})
// 	if !ok {
// 		return nil, fmt.Errorf("No se pudo acceder al campo 'hits'")
// 	}

// 	// Acceder al campo "hits" dentro de "hits"
// 	hitsArray, ok := hitsData["hits"].([]interface{})
// 	if !ok {
// 		return nil, fmt.Errorf("No se pudo acceder al campo 'hits'")
// 	}

// 	var contenidos []string

// 	// Recorrer el arreglo de hits
// 	for _, hit := range hitsArray {
// 		hitMap, ok := hit.(map[string]interface{})
// 		if !ok {
// 			fmt.Println("No se pudo convertir el hit a mapa")
// 			continue
// 		}

// 		// Acceder al campo "_source"
// 		source, ok := hitMap["_source"].(map[string]interface{})
// 		if !ok {
// 			fmt.Println("No se pudo acceder al campo '_source'")
// 			continue
// 		}
		
// 		// Acceder al campo "contenido"
// 		contenido, ok := source["contenido"].(string)
// 		if !ok {
// 			fmt.Println("No se pudo acceder al campo 'contenido'")
// 			continue
// 		}
		
// 		// Hacer algo con el campo "contenido", por ejemplo, imprimirlo
// 		contenidos = append(contenidos, contenido)
// 	}
	
// 	// fmt.Println("Contenidos", contenidos)
// 	return &contenidos, nil
// }

// func main() {
// 	host := "http://localhost:4080" // Reemplaza esto con tu verdadero host
// 	query := SearchRequest{
// 		SearchType: "match",
// 		Query: Query{
// 			Term:      "value",
// 			StartTime: "2023-01-01T14:30:00Z",
// 			EndTime:   "2024-01-07T14:30:00Z",
// 		},
// 		MaxResults: 5,
// 	}

// 	result, err := performSearch(host, query)
// 	if err != nil {
// 		fmt.Println("Error al realizar la búsqueda:", err)
// 		return
// 	}
// 	fmt.Println(string(result))

// 	// Imprime el resultado
// 	// resultJSON, err := json.MarshalIndent(result, "", "   ")
// 	// if err != nil {
// 	// 	fmt.Println("Error al formatear el resultado en JSON:", err)
// 	// 	return
// 	// }

// 	// fmt.Println(string(resultJSON))
// }