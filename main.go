package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	txtFilename := flag.String("txt", "words.txt",
		"a txt file with 1 word per line")
	flag.Parse()

	file, err := os.Open(*txtFilename)
	if err != nil {
		exit("Failed to open txt file.")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	fmt.Println("+---+\n|   |n|\n|\n|\n|\n=========")          //0
	fmt.Println("+---+\n|   |\n|   O\n|\n|\n|\n=========")     //
	fmt.Println("+---+\n|   |\n|   O\n|   |\n|\n|\n=========") //
	fmt.Println("+---+\n|   |\n|   O\n|  /|\n|\n|\n=========")
	fmt.Println("+---+\n|   |\n|   O\n|  /|\\\n|\n|\n=========")       //
	fmt.Println("+---+\n|   |\n|   O\n|  /|\\\n|  /\n|\n=========")    //
	fmt.Println("+---+\n|   |\n|   O\n|  /|\\\n|  / \\\n|\n=========") //6

}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
