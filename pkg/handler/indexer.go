package handler

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"back_go/pkg/zincsearch"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"bytes"
	"flag"
	"runtime/pprof"
	"runtime"
	"time"
	"sync"
)

var inserted int = 0
var emailbig int = 0
var rejected int = 0
var Emails []map[string]interface{}
var index string = "indexerEnron"
var max_size int64 = 1800
var wg sync.WaitGroup

// Construct the request and perform the petition of batch of files
func Indexer() {
	url := os.Getenv("ZINC_HOST") + ":" + os.Getenv("ZINC_PORT") + "/api/_bulkv2"

	request := zincsearch.CreateDocumentsRequest{
		Index:   index,
		Records: Emails,
	}

	jsonData, err := json.MarshalIndent(request, "", "   ")
	if err != nil {
		log.Println("Error converting to JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error reading request:", err)
		return
	}
	// Set headers
	req.SetBasicAuth(os.Getenv("ZINC_ADMIN_USER"), os.Getenv("ZINC_ADMIN_PASSWORD"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error performing HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Unexpected status code: %d", resp.StatusCode)
	}

	fmt.Println("Index insert", inserted, "Files")
	fmt.Println("Index reject by size", emailbig, "Files")
	fmt.Println("Index reject by error", rejected, "Files")
}

func addMail(path string, resultCh chan<- error) {
	defer wg.Done()
	content_file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("Error reading the file ", path, err)
		resultCh <- err
		return
	}

	data := map[string]interface{}{
		"email_content": string(content_file),
	}

	Emails = append(Emails, data)
	inserted++
	return
}

func Get_files(path string, f os.FileInfo, err error, resultCh chan<- error) {
	if err != nil {
		log.Println("Error while loading the file ", path, err)
		resultCh <- err
		return
	}

	if !f.IsDir() {
		if f.Size() > max_size {
			emailbig++
			return
		}

		wg.Add(1)
		go addMail(path, resultCh)
	}
}

// flags for profiling
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func Index() {
	flag.Parse()

	//cpu profile
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Println("Could not create CPU profile:", err)
		}
		defer f.Close() 
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Println("Could not start CPU profile:", err)
		}
		defer pprof.StopCPUProfile()
	}

	LoadEnv()
	ti := time.Now()
	resultCh := make(chan error, 1)

	err := filepath.Walk(os.Getenv("FILES_DIR"), func(path string, f os.FileInfo, err error) error {
		Get_files(path, f, err, resultCh)
		return nil
	})
	if err != nil {
		log.Println("Error walking through the directory: %v\n", err)
	}

	wg.Wait()
	close(resultCh)

	for err := range resultCh {
		if err != nil {
			rejected++
		}
	}

	fmt.Println("Duration: ", time.Since(ti))

	Indexer()

	//mem profile
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Println("Could not create memory profile:", err)
		}
		defer f.Close() 
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Println("Could not write memory profile:", err)
		}
	}
}
