package httpapi

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/ronitgavaskar/stocker/internal/orchestrator"
)

// Handler serves Stocker's JSON API. It is a thin translator: read the request,
// call the orchestrator, write the report as JSON. No business logic lives here.
type Handler struct {
	boss *orchestrator.Orchestrator
}

// NewHandler injects the orchestrator dependency.
func NewHandler(boss *orchestrator.Orchestrator) *Handler {
	return &Handler{boss: boss}
}

// Routes returns the API's HTTP routing, ready to hand to an http.Server. The
// "GET /report" pattern (Go 1.22+ ServeMux) also yields a 405 automatically for
// non-GET methods on that path.
func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /report", h.handleReport)
	return mux
}

// handleReport serves GET /report?ticker=AAPL with one assembled Report as JSON.
func (h *Handler) handleReport(w http.ResponseWriter, r *http.Request) {
	ticker := strings.ToUpper(strings.TrimSpace(r.URL.Query().Get("ticker")))
	if ticker == "" {
		// Bad INPUT -> 400. (A valid request whose sections all fail is still
		// 200: data quality lives in the body, not in the HTTP status code.)
		writeError(w, http.StatusBadRequest, "ticker query parameter is required")
		return
	}

	// r.Context() ties the work to the request: if the client disconnects, the
	// context cancels and the workers' tools see it and bail.
	rep := h.boss.Assemble(r.Context(), ticker)
	writeJSON(w, http.StatusOK, rep)
}

// errorBody is the consistent error envelope: { "error": "<message>" }.
type errorBody struct {
	Error string `json:"error"`
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, errorBody{Error: msg})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		// Status/headers are already sent, so we can't change the response now;
		// just record it. Usually this is a client disconnect mid-write.
		log.Printf("httpapi: encoding response: %v", err)
	}
}
