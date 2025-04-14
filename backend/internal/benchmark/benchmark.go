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

func RunBenchmark(url string) Result {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		return Result{
			Error: err.Error(),
		}
	}
	defer resp.Body.Close()

	duration := time.Since(start)

	return Result{
		StatusCode: resp.StatusCode,
		Duration:   float64(duration.Milliseconds()),
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
