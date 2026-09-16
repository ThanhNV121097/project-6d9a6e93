package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/ThanhNV121097/project-6d9a6e93/backend/migrations"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type greetingResponse struct {
	Text string `json:"text"`
}

type errorResponse struct {
	Error errorBody `json:"error"`
}

type errorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func main() {
	ctx := context.Background()
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		log.Fatal("DATABASE_URL is required")
	}
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		log.Fatalf("connect database: %v", err)
	}
	defer pool.Close()
	if err := migrate(ctx, pool); err != nil {
		log.Fatalf("migrate database: %v", err)
	}
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("probe database: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		probeCtx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		if err := pool.Ping(probeCtx); err != nil {
			writeError(w, http.StatusServiceUnavailable, "UNAVAILABLE", "Service unavailable.")
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	mux.HandleFunc("GET /v1/greeting", func(w http.ResponseWriter, r *http.Request) {
		text, err := getGreeting(r.Context(), pool)
		if err != nil {
			writeStoreError(w, err)
			return
		}
		writeJSON(w, http.StatusOK, greetingResponse{Text: text})
	})
	mux.HandleFunc("PUT /v1/greeting", func(w http.ResponseWriter, r *http.Request) {
		text, ok := decodeGreetingRequest(w, r)
		if !ok {
			return
		}
		if text == "" {
			writeError(w, http.StatusUnprocessableEntity, "VALIDATION_FAILED", "Greeting must not be empty.")
			return
		}
		saved, err := saveGreeting(r.Context(), pool, text)
		if err != nil {
			writeStoreError(w, err)
			return
		}
		writeJSON(w, http.StatusOK, greetingResponse{Text: saved})
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("APP_PORT")
	}
	if port == "" {
		port = "8080"
	}
	server := &http.Server{Addr: ":" + port, Handler: mux, ReadHeaderTimeout: 5 * time.Second}
	log.Printf("listening on :%s", port)
	log.Fatal(server.ListenAndServe())
}

func getGreeting(ctx context.Context, pool *pgxpool.Pool) (string, error) {
	var text string
	err := pool.QueryRow(ctx, `SELECT text FROM greetings WHERE id = true`).Scan(&text)
	if errors.Is(err, pgx.ErrNoRows) {
		err = pool.QueryRow(ctx, `INSERT INTO greetings (id, text, updated_at) VALUES (true, 'Hello, World!', now()) ON CONFLICT (id) DO UPDATE SET text = greetings.text RETURNING text`).Scan(&text)
	}
	return text, err
}

func saveGreeting(ctx context.Context, pool *pgxpool.Pool, text string) (string, error) {
	var saved string
	err := pool.QueryRow(ctx, `UPDATE greetings SET text = $1, updated_at = now() WHERE id = true RETURNING text`, text).Scan(&saved)
	if errors.Is(err, pgx.ErrNoRows) {
		err = pool.QueryRow(ctx, `INSERT INTO greetings (id, text, updated_at) VALUES (true, $1, now()) RETURNING text`, text).Scan(&saved)
	}
	return saved, err
}

func decodeGreetingRequest(w http.ResponseWriter, r *http.Request) (string, bool) {
	defer r.Body.Close()
	decoder := json.NewDecoder(io.LimitReader(r.Body, 1<<20))
	decoder.DisallowUnknownFields()
	var req struct {
		Text string `json:"text"`
	}
	if err := decoder.Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "MALFORMED_REQUEST", "Request body is malformed.")
		return "", false
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		writeError(w, http.StatusBadRequest, "MALFORMED_REQUEST", "Request body is malformed.")
		return "", false
	}
	return strings.TrimSpace(req.Text), true
}

func writeStoreError(w http.ResponseWriter, err error) {
	if isUnavailable(err) {
		writeError(w, http.StatusServiceUnavailable, "UNAVAILABLE", "Service unavailable.")
		return
	}
	writeError(w, http.StatusInternalServerError, "INTERNAL", "Internal server error.")
}

func isUnavailable(err error) bool {
	var netErr net.Error
	if errors.As(err, &netErr) {
		return true
	}
	return errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded)
}

func writeError(w http.ResponseWriter, status int, code string, message string) {
	writeJSON(w, status, errorResponse{Error: errorBody{Code: code, Message: message}})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func migrate(ctx context.Context, pool *pgxpool.Pool) error {
	if _, err := pool.Exec(ctx, `CREATE TABLE IF NOT EXISTS schema_migrations (version text PRIMARY KEY, applied_at timestamptz NOT NULL DEFAULT now())`); err != nil {
		return err
	}
	entries, err := fs.ReadDir(migrations.Files, ".")
	if err != nil {
		return err
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].Name() < entries[j].Name() })
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".up.sql") {
			continue
		}
		version := strings.TrimSuffix(name, ".up.sql")
		var applied bool
		if err := pool.QueryRow(ctx, `SELECT EXISTS(SELECT 1 FROM schema_migrations WHERE version = $1)`, version).Scan(&applied); err != nil || applied {
			if err != nil { return err }
			continue
		}
		sql, err := migrations.Files.ReadFile(name)
		if err != nil { return err }
		if strings.Contains(strings.ToUpper(string(sql)), "CREATE INDEX CONCURRENTLY") {
			if _, err = pool.Exec(ctx, string(sql)); err != nil { return err }
			if _, err = pool.Exec(ctx, `INSERT INTO schema_migrations (version) VALUES ($1)`, version); err != nil { return err }
			continue
		}
		tx, err := pool.Begin(ctx)
		if err != nil { return err }
		if _, err = tx.Exec(ctx, string(sql)); err == nil {
			_, err = tx.Exec(ctx, `INSERT INTO schema_migrations (version) VALUES ($1)`, version)
		}
		if err != nil { _ = tx.Rollback(ctx); return err }
		if err = tx.Commit(ctx); err != nil { return err }
	}
	return nil
}

