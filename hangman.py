import random
from words import word_list

def get_word():
    word = random.choice(word_list)
    return word.upper()


def play(word):
    word_completion = "_" * len(word)
    guessed = False 
    guessed_letters = []
    guessed_words = []
    tries = 6
    print("jouons au pendu")
    printhard(tries)
    print(word_completion)
    print("\n")
    while not guessed and tries > 0:
        guess = input("s'il te plaît, devine une lettre ou un mot :").upper()
        if len(guess) == 1 and guess.isalpha():
            if guess in guessed_letters:
                print("vous avez devinez la lettre")
            elif guess not in word:
                print( guess,"n'est pas dans le mot")    
                tries -= 1
                guessed_letters.append(guess)
            else:
                print("bien joué,", guess, "est dans le mot!")
                guessed_letters.append(guess)
                word_as_list = list(word_completion)
                indices = [i for i,letter in enumerate(word)if letter == guess]
                for index in indices:
                    word_as_list[index] = guess
                word_completion = "".join(word_as_list)
                if "_" not in word_completion:
                    guessed = True
        elif len(guess) == len(word) and guess.isalpha():
            if guess in guessed_words:
                print("vous avez déjà deviné le mot", guess)
            elif guess != word:
                print(guess, "n'est pas dans le mot")
                tries -= 1 
                guessed_words.apprend(guess) 
            else:
                guessed = True 
                word_completion = word 

        else:
            print("pas une supposition valable") 
        printhard(tries)
        print(word_completion)
        print("\n") 
    if guessed:
        print("félicitations,vous avez deviné le mot! vous gagnez!") 
    else:
        print("désolé,vous avez manqué d'essais. le mot était "+ word +"peut-être la prochaine fois!")   
            
def printhard(errcount):

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

def main():
    word = get_word()
    play(word)
    while input("joué a nouveau? (O/N) ").upper() == "o":
        word = get_word()
        play(word)


if __name__ == "__main__":
    main()

play(get_word())