package example

import "log/slog"

func foo() {
	slog.Info("hello") // want "log.Info method"
	slog.Warn("world") // want "log.Warn method"

	slog.Info("Server is started, it's password: amiripass")
	slog.Info("⚡️blinding lights⚡️")
}
