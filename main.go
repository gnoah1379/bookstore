package main

import (
	"bookstore/cmd"
	"log/slog"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		slog.Error("failed to execute command", "error", err)
	}
}
