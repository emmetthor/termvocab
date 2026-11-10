package main

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/emmetthor/termvocab/internal/storage"
)

func main() {
	err := run()

	if err != nil {
		slog.Error(err.Error())
	}
}

// run runs the main CLI commands although it does not use CLI commands now.
func run() error {
	configDir, err := os.UserConfigDir()

	if err != nil {
		return err
	}

	dataDir := filepath.Join(configDir, "termvocab")

	storage, err := storage.Open_db(dataDir)

	_ = *storage

	if (err != nil) {
		return err
	}

	slog.Debug("Storage opened")

	return nil
}