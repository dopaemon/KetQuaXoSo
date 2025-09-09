package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"KetQuaXoSo/internal/configs"
	"KetQuaXoSo/internal/rss"
)

type CheckRequest struct {
	Province string `json:"province"`
}

type CheckResponse struct {
	Province string      `json:"province"`
	Results  interface{} `json:"results"`
	Error    string      `json:"error,omitempty"`
}

func RunAPI() {
	mux := http.NewServeMux()
	RegisterHandlers(mux)

	fmt.Println("API server chạy tại http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Server error:", err)
	}
}

func RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/api/province", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		writeJSON(w, configs.Provinces)
	})

	mux.HandleFunc("/api/check", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var req CheckRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		if req.Province == "" {
			http.Error(w, "province is required", http.StatusBadRequest)
			return
		}

		url := rss.Sources(req.Province)
		if url == "" {
			http.Error(w, fmt.Sprintf("Unknown province: %s", req.Province), http.StatusBadRequest)
			return
		}

		data, err := rss.Fetch(url)
		if err != nil {
			writeJSON(w, CheckResponse{
				Province: req.Province,
				Error:    "failed to fetch RSS: " + err.Error(),
			})
			return
		}

		results, err := rss.Parse(data)
		if err != nil {
			writeJSON(w, CheckResponse{
				Province: req.Province,
				Error:    "failed to parse RSS: " + err.Error(),
			})
			return
		}

		writeJSON(w, CheckResponse{
			Province: req.Province,
			Results:  results,
		})
	})
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(data)
}
