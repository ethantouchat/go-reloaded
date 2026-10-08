package entrée

import "strings"

func fixPunctuation(text string) string {

	// Ponctuation classique
	punctuation := []string{".", ",", "!", "?", ":", ";"}

	for _, p := range punctuation {
		// "word ," -> "word,"
		text = strings.ReplaceAll(text, " "+p, p)
	}

	// Nettoyage des espaces avant la ponctuation
	// et ajout d'un espace après
	words := strings.Fields(text)
	result := ""

	for i := 0; i < len(words); i++ {
		word := words[i]

		if result != "" {
			result += " "
		}

		result += word
	}

	text = result

	// Gestion des groupes de ponctuation
	text = strings.ReplaceAll(text, ". . .", "...")
	text = strings.ReplaceAll(text, "! !", "!!")
	text = strings.ReplaceAll(text, "? ?", "??")
	text = strings.ReplaceAll(text, "! ?", "!?")
	text = strings.ReplaceAll(text, "? !", "!?")

	// Gestion des apostrophes
	text = strings.ReplaceAll(text, "' ", "'")
	text = strings.ReplaceAll(text, " '", "'")

	return text
}
