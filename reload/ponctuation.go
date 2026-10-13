package entrée

import (
	"regexp"
	"strings"
)

func Punctuation(text string) string {
	// Supprime les espaces avant la ponctuation
	re1 := regexp.MustCompile(`\s+([.,!?;:]+)`)
	text = re1.ReplaceAllString(text, "$1")

	// Ajoute un espace après la ponctuation si nécessaire
	// sauf si le caractère suivant est déjà un espace ou la fin du texte
	re2 := regexp.MustCompile(`([.,!?;:]+)([^\s.,!?;:])`)
	text = re2.ReplaceAllString(text, "$1 $2")

	// Nettoie les espaces multiples
	text = strings.TrimSpace(text)
	re3 := regexp.MustCompile(`\s+`)
	text = re3.ReplaceAllString(text, " ")

	return text
}
