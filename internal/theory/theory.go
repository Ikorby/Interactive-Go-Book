package theory

import (
	"fmt"
	"os"
	"path/filepath"
)

func Greeting(lang string) {
	switch lang {
	case "en":
		fmt.Println("\nYou selected English version!")
		break
	case "ru":
		fmt.Println("\nВы выбрали русскую версию!")
		break
	default:
		fmt.Println("\nThis language is not supported at the moment.")
		break
	}
	showChapters(lang)

}

func showChapters(lang string) {
	dir := filepath.Join("content", lang)

	files, err := os.ReadDir(dir)
	if err != nil {
		switch lang {
		case "ru":
			fmt.Println("\nОшибка при чтении глав:", err)
			break
		default:
			fmt.Println("\nError reading chapters:", err)
			break
		}
		return
	}

	switch lang {
	case "en":
		fmt.Printf("\nAvailable chapters:\n\n")
		break
	case "ru":
		fmt.Printf("\nДоступные главы:\n\n")
		break
	}
	for _, f := range files {
		fmt.Printf("%s\n", f.Name())
	}
	selectChapter()
}

func selectChapter() {
	
}
