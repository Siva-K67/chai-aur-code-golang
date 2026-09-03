package main

import (
	"fmt"
	"log/slog"
)

func main() {
	fmt.Println("server started")

	slog.Info("Server started")
	slog.Debug("debug message")
	slog.Info("info message")
	slog.Warn("warning message")
	slog.Error("error message")
}

/*

output

server started
2026/09/03 12:08:02 INFO Server started

*/
