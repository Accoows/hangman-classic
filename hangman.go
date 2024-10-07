package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

var maxtentative int = 10
var compteurMax int = 8
var compteurMin int = 0
var hangmanStages []string

// Fonction pour lire les mots depuis le fichier
func readWordsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, scanner.Err()
}

// Fonction pour lire les affichages du pendu depuis le fichier
func readHangmanStages(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var stages []string
	scanner := bufio.NewScanner(file)
	var stage strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		// Séparation d'étapes de hangman
		if line == "" && stage.Len() > 0 {
			stages = append(stages, stage.String())
			stage.Reset()
		} else {
			stage.WriteString(line + "\n")
		}
	}
	if stage.Len() > 0 {
		stages = append(stages, stage.String()) // ajouter la dernière étape
	}
	return stages, scanner.Err()
}

// Fonction pour ajuster les compteurs pour la sélection des mots
// Affichage de 8 par 8
func adjustCounters(wordsCount int, shouldRemove bool) {
	if !shouldRemove {
		compteurMin += 8
		compteurMax += 8
	}
	if compteurMax > wordsCount {
		compteurMax = wordsCount
	}
	if compteurMin > wordsCount {
		compteurMin = wordsCount - 8
		if compteurMin < 0 {
			compteurMin = 0
		}
	}
}

// Fonction pour obtenir les mots avec les compteurs ajustés
func trait(shouldRemove bool) []string {
	trait, _ := readWordsFromFile("hangman.txt")
	wordsCount := len(trait)

	adjustCounters(wordsCount, shouldRemove)

	if wordsCount == 0 || compteurMin >= wordsCount {
		return nil
	}

	return trait[compteurMin:compteurMax]
}

// Fonction pour choisir un mot et révéler quelques lettres
func findWordInFile(words []string) (string, []bool) {
	word := words[rand.Intn(len(words))]
	wordrevealed := make([]bool, len(word))
	lettersreveal := (len(word)/2 - 1)

	for i := 0; i < lettersreveal; i++ {
		indexletter := rand.Intn(len(word))
		wordrevealed[indexletter] = true
	}
	return word, wordrevealed
}

// Fonction pour afficher le mot avec les lettres révélées
func displayWordFind(word string, wordrevealed []bool) {
	for i, letter := range word {
		if wordrevealed[i] {
			fmt.Printf("%c ", letter)
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
}

// Fonction pour afficher l'état du pendu
func displayHangman(stage int) {
	if stage >= 0 && stage < len(hangmanStages) {
		fmt.Println(hangmanStages[stage])
	}
}

// Fonction pour vérifier si le caractère est bien une lettre
func isAlphabetic(char rune) bool {
	return unicode.IsLetter(char)
}

// Fonction pour vérifier si tous les lettres sont révélées
func allRevealed(wordrevealed []bool) bool {
	for _, ltr := range wordrevealed {
		if !ltr {
			return false
		}
	}
	return true
}

// Fonction principale
func main() {
	var err error
	// Initialisation de fichier
	hangmanStages, err = readHangmanStages("hangman.txt") // Fichier des étapes du pendu
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier des étapes du pendu :", err)
		return
	}

	words, err := readWordsFromFile(os.Args[1]) // Fichier des mots à partir de la ligne de commande
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}
	//

	word, wordrevealed := findWordInFile(words)
	tentative := maxtentative
	lettersreveal := map[rune]bool{}

	fmt.Println("Bonne chance, vous avez", maxtentative, "tentatives.")
	displayWordFind(word, wordrevealed)

	user := bufio.NewReader(os.Stdin)

	for tentative > 0 {
		fmt.Print("Choisissez une lettre : ")
		input, _ := user.ReadString('\n')
		letter := strings.TrimSpace(strings.ToUpper(input))

		if len(letter) != 1 {
			fmt.Println()
			displayWordFind(word, wordrevealed)
			fmt.Println()
			fmt.Println("Veuillez saisir une seule lettre")
			continue
		}

		revealedrune := rune(letter[0])

		// Vérification de caractère (alphabétique)
		if !isAlphabetic(revealedrune) {
			fmt.Println("Ce n'est pas une lettre.")
			continue
		}

		if lettersreveal[revealedrune] {
			fmt.Println()
			displayWordFind(word, wordrevealed)
			fmt.Println()
			fmt.Println("Vous avez déjà donné cette lettre.")
			continue
		}
		lettersreveal[revealedrune] = true

		wordfound := false
		for i, ltr := range word {
			if rune(strings.ToUpper(string(ltr))[0]) == revealedrune {
				wordrevealed[i] = true
				wordfound = true
			}
		}

		if wordfound {
			displayWordFind(word, wordrevealed)
			if allRevealed(wordrevealed) {
				fmt.Printf("Bravo ! Le mot est : %s\n", word)
				return
			}
			continue
		} else {
			tentative--
			fmt.Printf("Non présent dans le mot, %d tentatives restantes. \n", tentative)
			displayHangman(maxtentative - tentative - 1) //Départ de pendu à la position départ du fichier txt
		}

		displayWordFind(word, wordrevealed)
	}

	fmt.Println("Perdu ! Le mot était :", word)
}
