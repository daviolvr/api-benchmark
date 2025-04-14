package benchmark

import (
	"net/http"
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
