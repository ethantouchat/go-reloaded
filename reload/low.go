package entrée

import "strings"

func convertLow(text string) string {
	words := strings.Fields(text)

	for i := 1; i < len(words); i++ {
		if words[i] == "(low)" {
			// Convertit le mot précédent en minuscules
			words[i-1] = strings.ToLower(words[i-1])

			// Supprime "(low)"
			words = append(words[:i], words[i+1:]...)

			// Recule l'index
			i--
		}
	}

	return strings.Join(words, " ")
}
