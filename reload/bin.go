package entrée

import (
	"fmt"
	"strconv"
	"strings"
)

func convertBin(text string) string {
	words := strings.Fields(text)

	for i := 1; i < len(words); i++ {
		if words[i] == "(bin)" {
			// Convertit le mot précédent du binaire vers le décimal
			n, err := strconv.ParseInt(words[i-1], 2, 64)

			if err == nil {
				words[i-1] = fmt.Sprintf("%d", n)
			}

			// Supprime "(bin)"
			words = append(words[:i], words[i+1:]...)

			// Recule l'index
			i--
		}
	}

	return strings.Join(words, " ")
}
