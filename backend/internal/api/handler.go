package api

import (
	"benchmark-api/internal/benchmark"
	"benchmark-api/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func BenchmarkHandler(w http.ResponseWriter, r *http.Request) {
	// Verifica se o método é POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var req models.BenchmarkRequest

	// Transforma de JSON para Go e depois verifica se tem erro
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.URL == "" {
		http.Error(w, "Requisição inválida", http.StatusBadRequest)
		return
	}

	// Definir a quantidade de requisições (default é 1)
	count := 1
	if r.URL.Query().Has("count") {
		count, err = strconv.Atoi(r.URL.Query().Get("count"))
		if err != nil || count <= 0 {
			http.Error(w, "Parâmetro 'count' inválido", http.StatusBadRequest)
			return
		}
	}

	var totalDuration float64
	var result benchmark.Result

	// Realiza múltiplas requisições
	for i := 0; i < count; i++ {
		result = benchmark.RunBenchmark(req.URL)
		totalDuration += result.Duration
	}

	// Média do tempo
	avgDuration := totalDuration / float64(count)
	result.Duration = avgDuration
	result.Count = count

	// Salva o resultado no arquivo results.json
	_ = benchmark.SaveResult(benchmark.StoredResult{
		URL:        req.URL,
		StatusCode: result.StatusCode,
		Duration:   result.Duration,
		Error:      result.Error,
		Count:      result.Count,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func ResultsHandler(w http.ResponseWriter, r *http.Request) {
	results, err := benchmark.LoadResults()
	if err != nil {
		http.Error(w, "Erro ao carregar resultados", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
