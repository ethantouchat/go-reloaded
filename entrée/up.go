package entrée

import "strings"

func convertUp(text string) string {
	words := strings.Fields(text)

	for i := 1; i < len(words); i++ {
		if words[i] == "(up)" {
			// Convertit le mot précédent en majuscules
			words[i-1] = strings.ToUpper(words[i-1])

			// Supprime "(up)"
			words = append(words[:i], words[i+1:]...)

			// Recule l'index
			i--
		}
	}

	return strings.Join(words, " ")
}
