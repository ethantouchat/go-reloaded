package entrée

import (
	"fmt"
	"strconv"
	"strings"
)

func convertHex(text string) string {
	words := strings.Fields(text)

	for i := 1; i < len(words); i++ {
		if words[i] == "(hex)" {
			// Convertit le mot précédent de l'hexadécimal vers le décimal
			n, err := strconv.ParseInt(words[i-1], 16, 64)

			if err == nil {
				words[i-1] = fmt.Sprintf("%d", n)
			}

			// Supprime "(hex)"
			words = append(words[:i], words[i+1:]...)

			// On recule l'index car la liste a changé
			i--
		}
	}

	return strings.Join(words, " ")
}

func entrée() {
	text := "1E (hex) files were added"

	result := convertHex(text)

	fmt.Println(result)
}