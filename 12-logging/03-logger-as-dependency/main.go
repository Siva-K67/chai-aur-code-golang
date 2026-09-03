package main

import (
	"log/slog"
	"net/http"
	"os"
)

// a handler that takes a logger as a dependency, same idea as your
// db *sql.DB parameter pattern from the courses API
func courseHandler(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("handling request", "path", r.URL.Path, "method", r.Method)
		w.Write([]byte("courses list here"))
	}
}

func main() {

	/*
	   slog.NewTextHandler(os.Stdout, nil) → "I want text-formatted logs, written to the terminal, no special config."
	   slog.New(...) → "now give me an actual logger that uses that handler."
	   logger := ... → store that logger in a variable so you can call .Info()/.Error() on it later.
	*/

	/*
		Why two steps instead of one? This separation is intentional in Go's design — the handler controls how/where logs
		are formatted and sent (text vs JSON, terminal vs file vs a remote logging service), while the logger is what
		your code actually calls. This means you could swap NewTextHandler for slog.NewJSONHandler(os.Stdout, nil) —
		same logger interface, but now every log line comes out as JSON instead of plain text — without changing any
		of your .Info()/.Error() calls anywhere else in your code.
	*/

	// create ONE logger, shared across the whole application
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	// pass it into whichever handler needs it
	http.HandleFunc("/courses", courseHandler(logger))

	logger.Info("server starting", "port", 8080)
	http.ListenAndServe(":8080", nil)
}
