import random #programmation python pour le module random 

from words import word_list #indiquer que vous importez une variable nommée "word_list" à partir d'un module appelé "words"


def get_word(): #sélection des mots aléatoire 
    word = random.choice(word_list)
    return word.upper()


def play(word): #Pour utiliser la fonction de jeu pour le pendu le joueur doit deviner un mot. Il a 6 essais pour deviner les lettres ou le mot complet. La fonction affiche la progression du mot et informe le joueur s'il a gagné ou perdu
    word_completion = "_" * len(word)
    guessed = False
    guessed_letters = []
    guessed_words = []
    tries = 6
    print("Let's play Hangman!") #la fonction pour jouer a hangman
    printhard(tries)
    print(word_completion)
    print("\n")
    while not guessed and tries > 0:
        guess = input("Please guess a letter or word: ").upper() 
        if len(guess) == 1 and guess.isalpha():
            if guess in guessed_letters:
                print("You already guessed the letter", guess) # la fonction pour avoir réussi a trouver le mot 
            elif guess not in word:
                print(guess, "is not in the word.") #la fonction pour ne pas avoir trouver le mot 
                tries -= 1
                guessed_letters.append(guess)
            else:
                print("Good job,", guess, "is in the word!") 
                guessed_letters.append(guess)
                word_as_list = list(word_completion)
                indices = [i for i, letter in enumerate(word) if letter == guess]
                for index in indices:
                    word_as_list[index] = guess
                word_completion = "".join(word_as_list)
                if "_" not in word_completion:
                    guessed = True
        elif len(guess) == len(word) and guess.isalpha():
            if guess in guessed_words:
                print("You already guessed the word", guess)
            elif guess != word:
                print(guess, "is not the word.")
                tries -= 1
                guessed_words.append(guess)
            else:
                guessed = True
                word_completion = word
        else:
            print("Not a valid guess.") #
        printhard(tries) 
        print(word_completion)
        print("\n")
    if guessed:
        print("Congrats, you guessed the word! You win!")
    else:
        print("Sorry, you ran out of tries. The word was " + word + ". Maybe next time!")

def printhard(errcount): #une fonction pour afficher visuellement le pendu (`printhard`). La fonction principale (`main`) initialise un mot et lance le jeu, puis propose de jouer à nouveau. Enfin, une partie est automatiquement jouée à la fin.

    hangman = [
        "  +---+",
        "  |   |",
        "      |",
        "      |",
        "      |",
        "      |",
        "=========",
    ]

    if errcount == 5:
        hangman[2] = "  O   |"
    elif errcount == 4:
        hangman[2] = "  O   |"
        hangman[3] = "  |   |"
    elif errcount == 3:
        hangman[2] = "  O   |"
        hangman[3] = " /|   |"
    elif errcount == 2:
        hangman[2] = "  O   |"
        hangman[3] = " /|\\  |"
    elif errcount == 1:
        hangman[2] = "  O   |"
        hangman[3] = " /|\\  |"
        hangman[4] = " /    |"
    elif errcount == 0:
        hangman[2] = "  O   |"
        hangman[3] = " /|\\  |"
        hangman[4] = " / \\  |"

    for line in hangman:
        print(line)

def main(): # la fontion pour rejounée à nouveau ou pas 
    word = get_word()
    play(word)
    while input("joué a nouveau? (O/N) ").upper() == "o":
        word = get_word()
        play(word)


if __name__ == "__main__":
    main()

play(get_word())
