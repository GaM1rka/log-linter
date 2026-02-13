package a

import "log/slog"

func foo() {
	slog.Info("Server started") // want "message must start with lowercase letter"

	slog.Info("запуск сервера") // want "message must contain only English letters, digits and spaces" "message must not contain emoji or special symbols"

	slog.Info("token: abc") // want "message must contain only English letters, digits and spaces" "message must not contain emoji or special symbols" "message must not contain sensitive data"

	slog.Info("hello!") // want "message must contain only English letters, digits and spaces" "message must not contain emoji or special symbols"

	slog.Info("server started")
}
