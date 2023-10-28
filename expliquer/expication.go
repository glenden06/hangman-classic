// package + import necessaire au fonctionnement du programme
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

func terminalreset() {
	fmt.Print("\033c")
}

func main() {
	//Lire le fichier words.txt

	content, err := ioutil.ReadFile("words.txt") //Lire le document words.txt avec ioutil et prendre le contenu(content) et l'erreur en cas de probleme(err)

	if err != nil { //Si il y a une erreur
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
	win := "bien joué good job !"
	lose := "t'as perdu dommages écoute l'album her loss pour te réconforter"

	var lettressai []string

	for nimput := 1; nimput <= 26; nimput++ {
		terminalreset()
		PrintHangman(erreur)
		fmt.Println(ligne)
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
			//Si le mot est complet, afficher le message victoire et arreter le programme
			if strings.Join(ligne, "") == hidenword {
				terminalreset()
				fmt.Println(ligne)
				fmt.Println(win)
				time.Sleep(2 * time.Second)
				hangman()
				break
			}
		} else {
			//Si la lettre n'est pas dans le mot, afficher l'erreur et augmenter erreur
			terminalreset()
			fmt.Println("Erreur, cette lettre n'est pas dans le mot")
			erreur++
			time.Sleep(1 * time.Second)
			//Si erreur est egal a 6, afficher le message defaite et arreter le programme
			if erreur == 6 {
				terminalreset() //Affichage d'une erreur
				fmt.Println(lose)
				fmt.Println("Le mot etait ", hidenword) //affiche de la réponse
				time.Sleep(2 * time.Second)
				hangman()
				break
			}
		}
	}
}

func PrintHangman(erreur int) {
	//le tabeau vide du pendu
	hangman := []string{
		"  +---+",
		"  |   |",
		"      |",
		"      |",
		"      |",
		"      |",
		"=========",
	}
	//on corrige les lignes du tableau dans le cas ou errcount est equivalent à 1,2,3,4,5,6
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
	//On affiche le tableau
	for _, line := range hangman {
		fmt.Println(line)
	}
}
func hangman() {
	terminalreset()
	fmt.Println("Voulez-vous recommencer le pendu [o/n] ?")
	var retry string
	fmt.Scanln(&retry)
	if retry == "o" { //si oui
		main()
	}
	if retry == "n" { //si non
		terminalreset() //fin du programme et reset du terminal
		return
	}
	if retry != "n" && retry != "o" { //Si le choix n'est pas valide
		terminalreset()
		fmt.Println("[o] pour oui et [n] pour non, c'est pourtant pas complique.")
	}
}
