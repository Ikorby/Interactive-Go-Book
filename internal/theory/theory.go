package theory

import (
	"fmt"
	"os"
	"path/filepath"
)

func ShowChapters(lang string) {
	dir := filepath.Join("content", lang)

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading chapters:", err)
		return
	}

	fmt.Println("\nAvailable chapters:")
	for i, f := range files {
		fmt.Printf("%d. %s\n", i+1, f.Name())
	}

	// TODO: добавить выбор главы и вывод содержимого
}
