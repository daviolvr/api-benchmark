package benchmark

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type Result struct {
	StatusCode int     `json:"status_code"`
	Duration   float64 `json:"duration_ms"`
	Error      string  `json:"error,omitempty"`
	Count      int     `json:"count"`
}

func RunBenchmark(url string, headers map[string]string) Result {
	start := time.Now()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Result{Error: err.Error()}
	}

	// Adiciona headers se houver
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	duration := time.Since(start).Seconds() * 1000 // ms

	if err != nil {
		return Result{Duration: duration, Error: err.Error()}
	}
	defer resp.Body.Close()

	return Result{
		StatusCode: resp.StatusCode,
		Duration:   duration,
		Error:      "",
	}
}

func LoadResults() ([]StoredResult, error) {
	var results []StoredResult

	file, err := os.Open("data/results.json")
	if err != nil {
		// Caso o arquivo não exista, retorna uma lista vazia
		if os.IsNotExist(err) {
			return results, nil
		}
		return nil, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func ClearResults() error {
	return os.WriteFile("data/results.json", []byte("[]"), 0644)
}
