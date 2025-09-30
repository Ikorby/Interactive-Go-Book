package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Ikorby/Interactive-Go-Book/internal/theory"
)

func ShowMainMenu() {
	for {
		fmt.Println("\nWelcome to the Interactive Go Book!")
		fmt.Println("\n1. English\n2. Russian")
		fmt.Print("\nSelect your language (1 or 2): ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		option := strings.ToLower(scanner.Text())

		switch option {
		case "1", "english", "английский":
			startEnglishVersion()
			return
		case "2", "russian", "русский":
			startRussianVersion()
			return
		default:
			fmt.Println("Invalid option. Please select 1 or 2.")
		}
	}
}

func startEnglishVersion() {
	fmt.Println("\nYou selected English version!")
	theory.ShowChapters("en") // передаем язык
}

func startRussianVersion() {
	fmt.Println("\nВы выбрали русскую версию!")
	theory.ShowChapters("ru") // передаем язык
}