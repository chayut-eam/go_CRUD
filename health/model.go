package health

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alexliesenfeld/health"
	"crud/chrono"
)

type JSONResultWriter struct{}

func (rw *JSONResultWriter) Write(result *health.CheckerResult, statusCode int, w http.ResponseWriter, r *http.Request) error {
	checkerResult := checkerResult(result)
	jsonResp, err := json.Marshal(checkerResult)
	if err != nil {
		return fmt.Errorf("cannot marshal response: %w", err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	_, err = w.Write(jsonResp)
	return err
}

type CheckerResult struct {
	Status  health.AvailabilityStatus `json:"status"`
	Details *map[string]CheckResult   `json:"details,omitempty"`
}

type CheckResult struct {
	Status    health.AvailabilityStatus `json:"status"`
	Timestamp *chrono.ISO8601           `json:"timestamp,omitempty"`
	Error     *string                   `json:"error,omitempty"`
}

func checkerResult(result *health.CheckerResult) *CheckerResult {
	return &CheckerResult{
		Status:  result.Status,
		Details: checkResult(result.Details),
	}
}

func checkResult(details *map[string]health.CheckResult) *map[string]CheckResult {
	results := make(map[string]CheckResult)
	for name, result := range *details {
		timestamp := chrono.ISO8601{
			Time: result.Timestamp.UTC(),
		}
		checkResult := CheckResult{
			Status:    result.Status,
			Timestamp: &timestamp,
			Error:     result.Error,
		}
		results[name] = checkResult
	}
	return &results
}
