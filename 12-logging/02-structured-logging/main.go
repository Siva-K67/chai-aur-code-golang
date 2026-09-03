package main

import "log/slog"

func main() {
	slog.Info("user signed up")

	slog.Info("user signed up", "username", "Itachi", "age", 19)

	slog.Error("Failed to save course", "courseID", 284, "error", "Connection timeout")

}
