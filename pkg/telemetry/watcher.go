package telemetry

import (
	"context"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/warpbuilds/warpbuild-agent/pkg/log"
)

var watcher *fsnotify.Watcher

func disableOtelOutputFileWatcher() {
	watcher.Close()
}

func enableOtelOutputFileWatcher(ctx context.Context) error {
	var err error
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("failed to create file watcher: %w", err)
	}

	go func() {
		defer handlePanic()
		watchOtelOutputFile(ctx)
	}()
	return nil
}

func watchOtelOutputFile(ctx context.Context) {
	defer watcher.Close()

	// Ensure the log file exists
	if _, err := os.Stat(otelCollectorOutputFilePath); os.IsNotExist(err) {
		file, err := os.Create(otelCollectorOutputFilePath)
		if err != nil {
			log.Logger().Errorf("failed to create log file: %v", err)
		}
		file.Close()
	}

	err := watcher.Add(otelCollectorOutputFilePath)
	if err != nil {
		log.Logger().Errorf("failed to watch file: %v", err)
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op == fsnotify.Write {
				log.Logger().Infof("Modified file:", event.Name)
				debouncedOtelUpload(ctx)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Logger().Errorf("Error watching file:", err)
		}
	}
}
