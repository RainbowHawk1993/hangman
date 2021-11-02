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

func game(word string) {

	word = "unsubstantiated"
	fmt.Println(word) //for testing, remove later

	length := len(word)

	var underscores []string
	for i := 0; i < length; i++ {
		underscores = append(underscores, "_")
	}

	lives := 0
	win := true
	for win {
		hangman(lives)
		fmt.Print(underscores)

		var letter string
		//checking if more than 1 letter was entered
		temptrue := true
		for temptrue {
			fmt.Print("\nGuess a letter: ")
			fmt.Scanf("%s \n", &letter)
			if len(letter) == 1 {
				temptrue = false
			} else {
				fmt.Println("You need to enter only 1 letter")
			}
		}

		letter = strings.TrimSpace(letter) //I'm 99% sure Scanf trims spaces automatically, but I want to be safe
		letter = strings.ToLower(letter)   //Changing letter to lowercase because words are only in lowercase
		//fmt.Print(letter)

		m := regexp.MustCompile(letter)
		indexes := m.FindAllStringIndex(word, -1) //this gets us a 2D array and I convert it to 1D to stuff simpler
		//fmt.Print("indexes: ", indexes, indexes[0][1], indexes[1][1], indexes[0][0], indexes[1][0])
		var index []int
		var row = 0

		//this gets us 1D array of indexes
		for _, column := range indexes {
			index = append(index, column[row])
		}
		fmt.Print(index)

		/*for i, r := range underscores{
			if indexes[i][i+1]
		}*/
		win = false
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
	txtFilename := flag.String("txt", "words.txt",
		"a txt file with 1 word per line, all words are expected to be lowercase")
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
	game(randword)

}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
