package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Ikorby/Interactive-Go-Book/internal/theory"
)

func MenuGreeting() {
	fmt.Println("\nWelcome to the Interactive Go Book!")
}
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
			theory.Greeting("en")
			return
		case "2", "russian", "русский":
			theory.Greeting("ru")
			return
		default:
			fmt.Printf("\nInvalid option. Please select 1 or 2.\n\n")
		}
	}
}