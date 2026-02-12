package example

import "log/slog"

func foo() {
	slog.Info("hello") // want "log.Info method"
	slog.Warn("world") // want "log.Warn method"
}
