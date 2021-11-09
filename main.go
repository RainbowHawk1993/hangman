package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

//chosing a random word from the list
func chooserand(list []string) string {
	rand.Seed(time.Now().Unix()) //needed so randomization works correctly
	randomindexes := rand.Intn(len(list))
	randword := list[randomindexes]
	return randword
}

func game(word string, text []string) {
	word = strings.ToLower(word) //changing this to lowercase so we don't have issues with capitalization
	length := len(word)

	var underscores []string
	for i := 0; i < length; i++ {
		underscores = append(underscores, "_")
	}
	underscoresJustString := strings.Join(underscores, "") //need underscores to not be a slice

	var letters, allletters []string
	lives := 0
	win := true
	for win {
		hangman(lives)
		if lives == 6 {
			fmt.Println("YOU LOSE!")
			fmt.Println("The word was:", word)

			var PlayAgain string
			fmt.Print("\nWould you like to play again? (Y/N) ")
			fmt.Scanf("%s \n", &PlayAgain)
			PlayAgain = strings.ToLower(PlayAgain)

			if PlayAgain == "y" {
				randword := chooserand(text)
				game(randword, text)
			} else {
				exit("YOU LOSE!")
			}

		}

		fmt.Println(underscoresJustString)

		if lives > 0 {
			fmt.Println("Letters you have already checked for that aren't in the word:", letters)
		}

		var letter string

		temptrue := true
		for temptrue {
			fmt.Print("\nGuess a letter: ")
			fmt.Scanf("%s \n", &letter)

			//checking if more than 1 letter was entered
			if len(letter) == 1 {
				temptrue = false
			} else {
				fmt.Println("You need to enter only 1 letter")
			}
			letter = strings.TrimSpace(letter) //not needed?
			letter = strings.ToLower(letter)   //Changing letter to lowercase because words are only in lowercase

			//checks if the letter was entered before
			for _, let := range allletters {
				if letter == let {
					fmt.Println("You have already entered this letter")
					temptrue = true
				}
			}
			//checks if letter is in fact a letter
			alphabetcheck := false
			for r := 'a'; r < 'z'; r++ {
				if letter == string(r) {
					alphabetcheck = true
				}
			}
			if alphabetcheck == false {
				fmt.Println("Enter only letters, words don't contain numbers or special symbols")
				temptrue = true
			}
		}

		allletters = append(allletters, letter) //keeping list of all entered letters

		if strings.Contains(word, letter) {
			m := regexp.MustCompile(letter)
			indexes := m.FindAllStringIndex(word, -1) //this gets us a 2D array of indexes of this letter in the word and I convert it to 1D to make stuff simpler

			var index []int
			var row = 0

			//this gets us 1D array of indexes
			for _, column := range indexes {
				index = append(index, column[row])
			}

			for i := range index {
				indexnumber := index[i]
				underscoresJustString = underscoresJustString[:indexnumber] + letter + underscoresJustString[indexnumber+1:]
			}
		} else {
			fmt.Println("Seems like this letter doesn't exist in this word!")
			letters = append(letters, letter) //making list of letters that user checked for already
			lives++
		}

		if strings.Contains(underscoresJustString, "_") {
			//do nothing
		} else {
			word = strings.ToUpper(word)
			fmt.Println("\nCONGRATULATIONS, YOU WIN! THE WORD WAS: ", word)

			var PlayAgain string
			fmt.Print("\nWould you like to play again? (Y/N) ")
			fmt.Scanf("%s \n", &PlayAgain)
			PlayAgain = strings.ToLower(PlayAgain)

			if PlayAgain == "y" {
				randword := chooserand(text)
				game(randword, text)
			} else {
				exit("YOU WIN!")
			}
		}
	}

}

func hangman(counter int) {
	switch counter {
	case 0:
		fmt.Println("+---+\n|   |\n|\n|\n|\n|\n=========") //0
	case 1:
		fmt.Println("+---+\n|   |\n|   O\n|\n|\n|\n=========") //1
	case 2:
		fmt.Println("+---+\n|   |\n|   O\n|   |\n|\n|\n=========") //2
	case 3:
		fmt.Println("+---+\n|   |\n|   O\n|  /|\n|\n|\n=========") //3
	case 4:
		fmt.Println("+---+\n|   |\n|   O\n|  /|\\\n|\n|\n=========") //4
	case 5:
		fmt.Println("+---+\n|   |\n|   O\n|  /|\\\n|  /\n|\n=========") //5
	case 6:
		fmt.Println("+---+\n|   |\n|   O\n|  /|\\\n|  / \\\n|\n=========") //6
	}

}

func main() {
	txtFilename := flag.String("txt", "words.txt", "a txt file with 1 word per line")
	flag.Parse()

	file, err := os.Open(*txtFilename)
	if err != nil {
		exit("Failed to open txt file.")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	randword := chooserand(text)
	game(randword, text)

}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
