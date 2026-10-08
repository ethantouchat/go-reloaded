package entrée

import (
	"fmt"
	"strconv"
	"strings"
)

func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}

	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}

func convertCase(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {

		// Vérifie si c'est une instruction avec un nombre
		if strings.HasPrefix(words[i], "(up,") {
			number := strings.TrimSuffix(
				strings.TrimPrefix(words[i], "(up,"),
				")",
			)

			n, err := strconv.Atoi(number)
			if err == nil {
				for j := 1; j <= n && i-j >= 0; j++ {
					words[i-j] = strings.ToUpper(words[i-j])
				}
			}

			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}

		if strings.HasPrefix(words[i], "(low,") {
			number := strings.TrimSuffix(
				strings.TrimPrefix(words[i], "(low,"),
				")",
			)

			n, err := strconv.Atoi(number)
			if err == nil {
				for j := 1; j <= n && i-j >= 0; j++ {
					words[i-j] = strings.ToLower(words[i-j])
				}
			}

			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}

		if strings.HasPrefix(words[i], "(cap,") {
			number := strings.TrimSuffix(
				strings.TrimPrefix(words[i], "(cap,"),
				")",
			)

			n, err := strconv.Atoi(number)
			if err == nil {
				for j := 1; j <= n && i-j >= 0; j++ {
					words[i-j] = capitalize(words[i-j])
				}
			}

			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}

		// (up)
		if words[i] == "(up)" {
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}

		// (low)
		if words[i] == "(low)" {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}

		// (cap)
		if words[i] == "(cap)" {
			words[i-1] = capitalize(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}
	}

	return strings.Join(words, " ")
}

func main() {
	text := "This is so exciting (up, 2)"

	fmt.Println(convertCase(text))
}
