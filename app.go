package main

import (
	"context"
	"fmt"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) SplitFile(filePath string, fileCount int, delim string, clean bool) {
	openFile, err := os.Open(filePath)
	if err != nil {
		return
	}
	fmt.Printf("Splitting file: %s", openFile.Name())
	fmt.Printf("File count: %d, Delimiter: %s, Clean: %t", fileCount, delim, clean)
}

func (a *App) SystemDirectorySelect() (string, error) {
	path, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title:                "Select a Directory",
		ShowHiddenFiles:      true,
		CanCreateDirectories: true,
	})
	if err != nil {
		return "", err
	}
	return path, nil
}

func (a *App) SystemFileSelect() (string, error) {
	path, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Filters: []runtime.FileFilter{
			{
				DisplayName: "CSV",
				Pattern:     "*.csv",
			},
		},
	})
	if err != nil {
		return "", err
	}
	return path, nil

}
