package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
)

func main() {

	content, err := ioutil.ReadFile("words.txt")

	if err != nil {
		log.Fatal(err)
	}

	var List []string
	List = strings.Split(string(content), "\n")

	randomlist := rand.Intn(len(List))
	hidenword := List[randomlist-1]

	var lettrecache []string
	lettrecache = strings.Split(hidenword, "")

	var ligne []string
	for i := 0; i < len(lettrecache); i++ {
		ligne = append(ligne, "_")
	}
	fmt.Println(ligne)

	erreur := 0
	win := "bien joué fils d'eup"
	lose := "t'as perdu dommages écoute l'album her loss pour te réconforter"

	var lettressai []string

	for nimput := 1; nimput <= 26; nimput++ {
		fmt.Println("Vous avez deja essayé les lettres suivantes: ", lettressai)
		fmt.Println("Vous avez encore le droit a ", 6-erreur, " erreur(s)")
		fmt.Println("Veuillez entrer une lettre: ")
		var input string
		fmt.Scanln(&input)
		lettressai = append(lettressai, input)
		//Si la lettre est dans le mot, afficher le mot avec la lettre
		if strings.Contains(hidenword, input) {
			for i := 0; i < len(lettrecache); i++ {
				if lettrecache[i] == input {
					ligne[i] = input
				}
			}
			fmt.Println(ligne)
			//Si le mot est complet, afficher le message victoire et arreter le programme
			if strings.Join(ligne, "") == hidenword {
				fmt.Println(win)
				break
			}
		} else {
			//Si la lettre n'est pas dans le mot, afficher l'erreur et augmenter erreur
			fmt.Println("Erreur, cette lettre n'est pas dans le mot")
			erreur++
			PrintHangman(erreur)
			//Si erreur est egal a 6, afficher le message defaite et arreter le programme
			if erreur == 6 {
				fmt.Println(lose)
				break
			}
		}
	}
}

func PrintHangman(erreur int) {
	hangman := []string{
		"  +---+",
		"  |   |",
		"      |",
		"      |",
		"      |",
		"      |",
		"=========",
	}

	switch erreur {
	case 1:
		hangman[2] = "  O   |"
	case 2:
		hangman[2] = "  O   |"
		hangman[3] = "  |   |"
	case 3:
		hangman[2] = "  O   |"
		hangman[3] = " /|   |"
	case 4:
		hangman[2] = "  O   |"
		hangman[3] = " /|\\  |"
	case 5:
		hangman[2] = "  O   |"
		hangman[3] = " /|\\  |"
		hangman[4] = " /    |"
	case 6:
		hangman[2] = "  O   |"
		hangman[3] = " /|\\  |"
		hangman[4] = " / \\  |"
	}

	for _, line := range hangman {
		fmt.Println(line)
	}
}
