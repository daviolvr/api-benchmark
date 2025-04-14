package benchmark

import (
	"encoding/json"
	"os"
	"sync"
)

var mu sync.Mutex

type StoredResult struct {
	URL        string  `json:"url"`
	StatusCode int     `json:"status_code"`
	Duration   float64 `json:"duration_ms"`
	Error      string  `json:"error,omitempty"`
	Count      int     `json:"count"`
}

func SaveResult(result StoredResult) error {
	mu.Lock()
	defer mu.Unlock()

	filePath := "data/results.json"

	// Lê o conteudo atual
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var results []StoredResult
	json.Unmarshal(data, &results)

	// Adiciona novo resultado
	results = append(results, result)

	// Salva novamente
	updatedData, err := json.MarshalIndent(results, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, updatedData, 0644)
	return err
}
